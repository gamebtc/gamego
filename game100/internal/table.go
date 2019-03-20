package internal

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol/folks"
	"local.com/abc/game/room"
)

const RicherCount = 6

type GameDriver interface {
	// 准备游戏, 状态1
	Ready(table *Table)
	// 开始下注, 状态2
	Open(table *Table)
    // 游戏中, 状态2
    Play(table *Table)
	// 停止下注, 状态3
	Stop(table *Table)
	// 发牌结算, 状态4
	Deal(table *Table)
}

// 开始下注：下注时间12秒
// 停止下注：发牌开：2秒
// 结算：    2秒
// 游戏桌子
type Table struct {
	GameDriver
	Id        int32           // 桌子ID
	CurId     int64           // 当前局的ID
	LastId    int64           // 最后的局ID
	Log       []byte          // 最后60局的发牌情况
	State     int32           // 0:暂停;1:准备;2:开始下注;3:停止下注;4:结算
	Roles     []*Role         // 所有真实游戏玩家
	Robot     []*Role         // 所有机器人
	Richer    []*Role 		  // 富豪
	round     *GameRound      // 1局
	roundFlow int             // 下注流索引
	continued int             // 持续秒数
}

func NewTable() *Table {
	dealer := newDriver()
	t := &Table{
		Roles:      make([]*Role, 0, 100),
		Robot:      make([]*Role, 0, 100),
		GameDriver: dealer,
	}
	return t
}

func(table *Table)MustWin()bool {
	// round.BetGroup != nil; 有真人下注
	return (table.round.UserBet != nil) && (mustWinRate > mustWinRand.Int31n(100))
}

func (table *Table) GetRichPlayer()[]*folks.Player{
	richer := make([]*folks.Player, len(table.Richer))
	for i, role := range table.Richer {
		richer[i] = role.GetRicher()
	}
	return richer
}

// 增加真实的玩家
func (table *Table) AddRole(role *Role) {
	table.Roles = append(table.Roles, role)
	// 真实玩家
	ack := &folks.GameInitAck{
		Id:    table.round.Id,
		State: table.State,
		Sum:   table.round.Group,
		Log:   table.Log,
		Rich:  table.GetRichPlayer(),
	}
	if bill := role.bill; bill != nil {
		ack.Bet = bill.Group
	}
	role.Send(ack)
}

// 查找1位赌神和5位富豪
func  (table *Table) FindRicher()[]int32 {
	roleCount := len(table.Roles) + len(table.Robot)
	if roleCount == 0 {
		table.Richer = []*Role{}
		return nil
	}

	roles := make([]*Role, roleCount)

	i := 0
	for _, role := range table.Robot {
		roles[i] = role
		i++
	}
	for _, role := range table.Roles {
		roles[i] = role
		i++
	}

	// 查找1位赌神
	richIndex := 0
	rich := roles[0]
	for i := 1; i < roleCount; i++ {
		b := roles[i]
		if rich.LastWinCount < b.LastWinCount || (rich.LastWinCount == b.LastWinCount && rich.LastBetSum < b.LastBetSum) {
			roles[i] = rich
			rich = b
			richIndex = i
		}
	}
	richer := []*Role{rich}

	roles = append(roles[:richIndex], roles[richIndex+1:]...)
	roleCount--

	// 查找5位富豪(以最近20局的下注金额排序,下注金额一样就以身上的钱排序)
	for c := 0; c < (RicherCount-1) && c < roleCount; c++ {
		rich := roles[c]
		for i := c + 1; i < roleCount; i++ {
			b := roles[i]
			if rich.LastBetSum < b.LastBetSum || (rich.LastBetSum == b.LastBetSum && rich.Coin < b.Coin) {
				// 交换
				roles[i] = rich
				rich = b
			}
		}
		richer = append(richer, rich)
	}
	table.Richer = richer

	richerId := make([]int32, len(richer))
	for i, r := range richer {
		richerId[i] = r.Id
	}
	return richerId
}

func (table *Table) NewGameRound() {
	id := room.NewGameRoundId()
	round := &GameRound{
		Id:        id,
		Start:     room.Now(),
		Room:      room.RoomId,
		Tab:       table.Id,
		Group:     make([]int64, betItemCount),
	}

    round.Rich = table.FindRicher()
	table.round = round
	table.roundFlow = 0
}

func (table *Table) Init(){
}

func (table *Table) Start() {
	table.continued = 1
	gameReady(table)
}

func (table *Table) Update() {
	table.continued--
	switch table.State {
	case 1:
		if table.continued == 0 {
			table.continued = 10
			gameOpen(table)
		}
	case 2:
		if table.continued == 0 {
			table.continued = 1
			gameStop(table)
		} else {
			gamePlay(table)
		}
	case 3:
		if table.continued == 0 {
			table.continued = 2
			gameDeal(table)
		}
	case 4:
		if room.Config.Pause == 0 {
			if table.continued == 0 {
				table.continued = 1
				gameReady(table)
			}
		} else {
			// 暂停
			table.continued++
			log.Debugf("Pause: %v", time.Now())
		}
	}
}

// 返回系统输赢
func (table *Table)CheckWin(odds []int32) int64 {
	if round := table.round; round != nil && round.UserBet != nil {
		prize, _, bet := Balance(round.UserBet, odds)
		return bet - prize
	}
	return 0
}

// 发送消息给所有在线玩家
func(table *Table)SendToAll(val interface{}) {
	if len(table.Roles) > 0 {
		if val, err := room.Coder.Encode(val); err != nil {
			for _, role := range table.Roles {
				role.UnsafeSend(val)
			}
		}
	}
}

// 投注
func betReq(m *room.NetMessage) {
	log.Debugf("betReq:%v", m)

	role, ok := m.Session.Role.(*Role)
	if ok == false || role == nil {
		log.Debugf("betReq: role is nil")
		return
	}
	//获取参数
	req, ok := m.Arg.(*folks.BetReq)
	if ok == false || req == nil {
		log.Debugf("betReq: is nil")
		return
	}
	log.Debugf("add bet: %#v", req)

	ack := &folks.BetAck{
		Sn:   req.Sn,
		Item: req.Item,
	}
	if err := role.AddBet(*req); err != nil {
		role.SendError(int32(folks.Folks_BetReq), 1000, err.Error(), "")
	} else {
		ack.Bet = req.Bet
	}
	ack.Coin = role.Coin
	role.Send(ack)
}

func(table *Table) loadRobot(count int32) {
	if count < 0 {
		// 退出部分机器人
	} else if count > 0 {
		// 增加机器人
		robots := db.Driver.LoadRobot(room.RoomId, count)
		for _, user := range robots {
			user.Job = model.JobRobot
			robot := &Role{
				Session: &room.Session{
					AgentId:  0,
					Ip:       user.Ip,
					Created:  time.Now(),
				},
				User:   user,
			}
			coin := rand.Int63n(200*room.Config.PlayMin) + (2 * room.Config.PlayMin)
			robot.table = table
			robot.Online = true
			robot.Coin = coin
			robot.Reset()
			table.Robot = append(table.Robot, robot)
			//log.Debugf("add robot: %#v, coin:%v", user, coin)
		}
	}
}

// 读取机器人配置
func(table *Table) robotConfig()int32{
	robotConf := room.Config.Robot
	end := len(robotConf) / 6
	now := time.Now()
	minute := int32(now.Hour()*60 + now.Minute())
	for count := 0; count < end; count++ {
		i := 6 * count
		if minute >= robotConf[i] && minute < robotConf[i+1] {
			min := robotConf[i+2]  //最小人数
			max := robotConf[i+3]  //最大人数
			base := robotConf[i+4] //基础人数
			rate := robotConf[i+5] //真实玩家的百分比人数
			//真人数量
			roleCount := int32(len(table.Roles))
			count := base + roleCount*rate/100
			if count < min {
				count = min
			} else if count > max {
				count = max
			}
			return count
		}
	}
	return -1
}

// 准备
func gameReady(table *Table) {
	// 删除已断线的玩家
	for i := 0; i < len(table.Roles); {
		role := table.Roles[i]
		role.Reset()
		if role.Online == false {
			table.Roles = append(table.Roles[:i], table.Roles[i+1:]...)
			role.UnlockRoom()
		} else {
			i++
		}
	}

	// 删除钱不足或者赢钱多的机器人
	var ids []int32
	for i := 0; i < len(table.Robot); {
		role := table.Robot[i]
		role.Reset()
		if role.Coin < room.Config.PlayMin ||
			role.TotalRound > rand.Int31n(100)+10 ||
			role.TotalWin > 10000*100 {
			ids = append(ids, role.Id)
			table.Robot = append(table.Robot[:i], table.Robot[i+1:]...)
		} else {
			i++
		}
	}
	if len(ids) > 0 {
		db.Driver.UnloadRobot(room.RoomId, ids)
	}

	count := table.robotConfig()
	if count >= 0 {
		table.loadRobot(count - int32(len(table.Robot)))
	}

	table.State = 1
	table.CurId += 1
	log.Debugf("%v准备:%v", gameName, table.CurId)
	table.NewGameRound()
	table.Ready(table)
}

// 开始
func gameOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = 2
	log.Debugf("开始下注:%v", table.CurId)
	table.Open(table)

	table.SendToAll(&folks.OpenBetAck{
		Id: table.CurId,
		Rich: table.GetRichPlayer(),
	})
}

func gamePlay(table *Table) {
	table.Play(table)
	// TODO: 需要优化机器人的投注项选择
	for _, role := range table.Robot {
		if rand.Int31n(4) == 1 {
			betIndex := rand.Intn(len(betItems))
			bet := folks.BetReq{
				Item: robetRandBetItem(),
				Bet:  betItems[betIndex],
			}
			for addCount := rand.Intn(3); addCount >= 0; addCount-- {
				if role.RobotCanBet(bet.Item, bet.Bet) {
					role.AddBet(bet)
					//log.Debugf("R%v下注:%v_%v,%v", role.Id, bet.Item, bet.Bet/100, role.Coin/100)
				}
			}
		}
	}

	l := len(table.round.Flow)
	if l > table.roundFlow {
		// 发送这段时间其他玩家的下注数据
		table.SendToAll(&folks.UserBetAck{
			Bet: table.round.Flow[table.roundFlow:l],
		})
		table.roundFlow = l
	}
}

// 停止下注
func gameStop (table *Table) {
	table.State = 3
	log.Debugf("停止下注:%v", table.CurId)
	table.Stop(table)
	table.SendToAll(&folks.StopBetAck{
		Id: table.CurId,
	})
}

// 发牌结算
func gameDeal(table *Table) {
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌结算
	table.Deal(table)

	round := table.round
	table.LastId = table.CurId
	round.End = room.Now()
	room.SaveLog(round)

	// 最后60局的对战日志
	table.Log = append(table.Log, round.Poker...)
	over := len(table.Log) - 60*betItemCount
	if over > 0 {
		table.Log = table.Log[over:]
	}

	// 结算结果发给玩家
	if len(table.Roles) > 0 {
		// 富豪玩家的输赢
		rich := make([]int64, len(table.Richer))
		for i, role := range table.Richer {
			if role.bill != nil {
				rich[i] = role.bill.Win
			}
		}
		r := &folks.GameResult{
			Id:    table.CurId,
			Poker: round.Poker,
			Odd:   round.Odds,
			Sum:   round.Group,
			Rich:  rich,
		}

		for _, role := range table.Roles {
			win := int64(0)
			if role.bill != nil {
				win = role.bill.Win
			}
			ack := &folks.GameDealAck{
				R:    r,
				Win:  win,
				Coin: role.Coin,
			}
			role.UnsafeSend(ack)
		}
	}
	log.Debugf("总下注:%v", round.Group)
}