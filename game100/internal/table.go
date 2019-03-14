package internal

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
	"local.com/abc/game/room"
)

const RicherCount = 6

type Plan struct {
	// 要执行的函数
	f func(*Table)
	// 下一步执行的时间
	d time.Duration
}

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
	Id     int32       //桌子ID
	CurId  int32       //当前局的ID
	LastId int32       //最后的局ID
	Log    []byte      //最后60局的发牌情况
	State  int32       //0:暂停;1:洗牌;2:下注;3:结算
	Roles  []*Role     //所有真实游戏玩家
	Robot  []*Role     //所有机器人
	Richer []*protocol.User //富豪
	round  *GameRound  //1局
	roundFlow int      //

	timer     *time.Timer
	planIndex int
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
	return (table.round.BetGroup != nil) && (mustWinRate > mustWinRand.Int31n(100))
}

// 增加真实的玩家
func (table *Table) AddRole(role *Role) {
	table.Roles = append(table.Roles, role)
	// 真实玩家
	ack := &protocol.FolksGameInitAck{
		Id:    table.round.Id,
		State: table.State,
		Sum:   table.round.Group,
		Log:   table.Log,
		Rich:  table.Richer,
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
		table.Richer = []*protocol.User{}
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
		if rich.LastWinCount < b.LastWinCount || (rich.LastWinCount == b.LastWinCount && rich.LastBetSum < b.LastBetSum){
			roles[i] = rich
			rich = b
			richIndex = i
		}
	}
	richer := []*protocol.User{rich.GetMsgUser()}
	log.Debugf("richer: %v, win:%v, bet:%v, coin:%v", rich.Id,rich.LastWinCount, rich.LastBetSum, rich.Coin)

	roles = append(roles[:richIndex], roles[richIndex+1:]...)
	roleCount--

	// 查找5位富豪(以最近20局的下注金额排序,下注金额一样就以身上的钱排序)
	for c := 0; c < 5 && c < roleCount; c++ {
		rich := roles[c]
		for i := c + 1; i < roleCount; i++ {
			b := roles[i]
			if rich.LastBetSum < b.LastBetSum || (rich.LastBetSum == b.LastBetSum && rich.Coin < b.Coin) {
				// 交换位置
				roles[i] = rich
				rich = b
			}
		}
		log.Debugf("richer: %v, win:%v, bet:%v, coin:%v", rich.Id,rich.LastWinCount, rich.LastBetSum, rich.Coin)
		richer = append(richer, rich.GetMsgUser())
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

	for _, role := range table.Roles {
		role.Reset()
	}
	for _, role := range table.Robot {
		role.Reset()
	}

    round.Rich = table.FindRicher()
	table.round = round
	table.roundFlow = 0
}

func (table *Table) Init(){
}

func (table *Table) Start() {
	table.timer = time.NewTimer(time.Microsecond)
}

func (table *Table) Update() {
	select {
	case <-table.timer.C:
		cur := table.planIndex
		table.timer.Reset(schedule[cur].d)
		if cur >= len(schedule)-1 {
			table.planIndex = 0
		} else {
			table.planIndex = cur + 1
		}
		schedule[cur].f(table)
	default:
	}
}

// 返回系统输赢
func (table *Table)CheckWin(odds []int32) int64 {
	if round := table.round; round != nil && round.BetGroup != nil {
		prize, _, bet := Balance(round.BetGroup, odds)
		return bet - prize
	}
	return 0
}

func(table *Table)SendToAll(val interface{}) {
	if len(table.Roles) > 0 {
		if val, err := room.Coder.Encode(val); err != nil {
			for _, role := range table.Roles {
				role.SendRaw(val)
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
	req, ok := m.Arg.(*protocol.BetReq)
	if ok == false || req == nil {
		log.Debugf("betReq: is nil")
		return
	}
	log.Debugf("add bet: %#v", req)

	ack := &protocol.BetAck{
		Sn:   req.Sn,
		Item: req.Item,
	}
	if err := role.AddBet(*req); err != nil {
		role.SendError(int32(protocol.MsgId_BetReq), 1000, err.Error(), "")
	} else {
		ack.Bet = req.Bet
	}
	ack.Coin = role.Coin
	role.Send(ack)
}

func(table *Table) unloadRobot() {
	// 删除钱不足或者赢钱多的机器人
	var ids []int32
	for i := 0; i < len(table.Robot); {
		role := table.Robot[i]
		if role.Coin < room.Config.PlayMin ||
			role.TotalRound > rand.Int31n(100)+10 ||
			role.TotalWin-role.TotalLost > 10000*100 {
			ids = append(ids, role.Id)
			table.Robot = append(table.Robot[:i], table.Robot[i+1:]...)
		} else {
			i++
		}
	}
	if len(ids) > 0 {
		db.Driver.UnloadRobot(room.RoomId, ids)
	}
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
				Sender: room.RobotSender,
				User:   user,
			}
			coin := rand.Int63n(200*room.Config.PlayMin) + (2 * room.Config.PlayMin)
			robot.table = table
			robot.Online = true
			robot.Coin = coin
			robot.Reset()
			table.Robot = append(table.Robot, robot)
			log.Debugf("add robot: %#v, coin:%v", user, coin)
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
	table.unloadRobot()
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
}

func gamePlay(table *Table) {
	table.Play(table)
	// TODO: 需要优化机器人的投注项选择
	for _, role := range table.Robot {
		if rand.Int31n(4) == 1 {
			betIndex := rand.Intn(len(betItems))
			bet := protocol.BetReq{
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
		table.SendToAll(&protocol.UserBetAck{
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
}

// 发牌结算
func gameDeal(table *Table) {
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌结算
	table.Deal(table)

	// 结算结果发给玩家
	table.LastId = table.CurId
	round := table.round
	round.End = room.Now()
	room.SaveLog(round)

	log.Debugf("总下注:%v", round.Group)
}