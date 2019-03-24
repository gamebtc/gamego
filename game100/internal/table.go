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

type GameState int32

const (
	GameStateReady   GameState = 0
	GameStatePlaying GameState = 1
	GameStateDeal    GameState = 2
)

type Dealer interface {
	// 发牌
	Deal(table *Table)
}

// 开始下注：下注时间12秒
// 停止下注：发牌开：2秒
// 结算：    2秒
// 游戏桌子
type Table struct {
	Dealer
	Id       int32                  // 桌子ID
	CurId    int64                  // 当前局的ID
	LastId   int64                  // 最后的局ID
	Log      []byte                 // 最后60局的发牌情况
	State    GameState              // 0:准备;1:下注中;2:结算
	Roles    map[model.UserId]*Role // 所有真实游戏玩家
	Robots   []*Role                // 所有机器人
	Richer   []*Role                // 富豪
	round    *GameRound             // 1局
	lastFlow int                    // 下注流索引
	delay    int32                  // 持续秒数
}

func NewTable() *Table {
	dealer := newDealer()
	t := &Table{
		Roles:  make(map[model.UserId]*Role, 256),
		Robots: make([]*Role, 0, 256),
		Dealer: dealer,
	}
	return t
}

func (table *Table) MustWin() bool {
	// round.BetGroup != nil; 有真人下注
	return (table.round.UserBet != nil) && (room.Config.WinRate > gameRand.Int31n(1000))
}

func (table *Table) GetRichPlayer() []*folks.Player {
	richer := make([]*folks.Player, len(table.Richer))
	for i, role := range table.Richer {
		richer[i] = role.GetPlayer()
	}
	return richer
}

// 增加真实的玩家
func (table *Table) addRole(role *Role) {
	role.player = &folks.Player{
		Id:   role.Id,
		Icon: role.Icon,
		Vip:  role.Vip,
		Name: role.Name,
		Coin: role.Coin,
	}
	role.table = table
	table.Roles[role.Id] = role
}

// 增加机器人
func (table *Table) addRobot(robot *Role) {
	robot.player = &folks.Player{
		Id:   robot.Id,
		Icon: robot.Icon,
		Vip:  robot.Vip,
		Name: robot.Name,
		Coin: robot.Coin,
	}
	robot.table = table
	table.Robots = append(table.Robots, robot)
}

// 初始化场景
func (table *Table) sendGameInit(role *Role) {
	// 真实玩家
	ack := &folks.GameInitAck{
		State: int32(table.State),
		Time:  table.delay,
		Log:   table.Log,
		Rich:  table.GetRichPlayer(),
	}
	if round := table.round; round != nil {
		ack.Id = round.Id
		ack.Sum = round.Group
	}
	if bill := role.bill; bill != nil {
		ack.Bet = bill.Group
	}
	role.Send(ack)
}

// 查找1位赌神和5位富豪
func (table *Table) findRicher() []model.UserId {
	roleCount := len(table.Roles) + len(table.Robots)
	if roleCount == 0 {
		table.Richer = []*Role{}
		return nil
	}

	roles := make([]*Role, 0, roleCount)
	roles = append(roles, table.Robots...)
	for _, role := range table.Roles {
		roles = append(roles, role)
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

func (table *Table) newGameRound() {
	count := room.PlanRobotCount(int32(len(table.Roles)))
	if count >= 0 {
		table.loadRobot(count - int32(len(table.Robots)))
	}
	id := room.NewGameRoundId()
	round := &GameRound{
		Id:    id,
		Start: room.Now(),
		Room:  room.RoomId,
		Tab:   table.Id,
		Group: make([]int64, betItemCount),
	}

	round.Rich = table.findRicher()
	table.round = round
	table.lastFlow = 0
}

func (table *Table) Init() {
}

func (table *Table) Start() {
	table.gameReady()
}

func (table *Table) Update() {
	table.delay--
	switch table.State {
	case GameStateReady:
		if room.Config.Pause != 0 {
			table.delay++
			return
		}
		if table.delay <= 0 {
			table.gameOpen()
		}
	case GameStatePlaying:
		if table.delay > 0 {
			table.gamePlay()
		} else {
			table.gameDeal()
		}
	case GameStateDeal:
		if table.delay <= 0 {
			table.gameReady()
		}
	}
}

// 返回系统输赢
func (table *Table) CheckWin(odds []int32) int64 {
	if round := table.round; round != nil && round.UserBet != nil {
		prize, _, bet := Balance(round.UserBet, odds)
		return bet - prize
	}
	return 0
}

// 发送消息给所有在线玩家
func (table *Table) SendToAll(val interface{}) {
	if len(table.Roles) > 0 {
		if val, err := room.Encode(val); err != nil {
			for _, role := range table.Roles {
				role.UnsafeSend(val)
			}
		}
	}
}

func (table *Table) loadRobot(count int32) {
	if count < 0 {
		// 退出机器人
		dec := int(-count)
		ids := make([]int32, 0, dec)
		for i := 0; i < len(table.Robots) && i < dec; i++ {
			role := table.Robots[i]
			ids = append(ids, role.Id)
		}
		dec = len(ids)
		if dec > 0 {
			table.Robots = table.Robots[dec:]
			db.Driver.UnloadRobot(room.RoomId, ids)
		}
	} else if count > 0 {
		// 增加机器人
		robots := db.Driver.LoadRobot(room.RoomId, count)
		for _, user := range robots {
			sess := &room.Session{
				Ip:      user.Ip,
				Created: time.Now(),
				AgentId: 0,
			}
			robot := &Role{
				User:    *user,
				Session: sess,
				table:   table,
			}
			robot.Online = true
			robot.Coin = rand.Int63n(200*room.Config.PlayMin) + (2 * room.Config.PlayMin)
			table.addRobot(robot)
		}
	}
}

func (table *Table) clearOffline() {
	// 删除已断线的玩家
	var ids []int32
	for k, role := range table.Roles {
		role.Reset()
		if role.Online == false {
			ids = append(ids, k)
			room.RemoveUser(role.Session)
			role.UnlockRoom()
			role.table = nil
			role.player = nil
		}
	}
	if len(ids) > 0 {
		for _, id := range ids {
			delete(table.Roles, id)
		}
		ids = ids[0:0]
	}

	// 删除钱不足或者钱多的机器人
	for i := 0; i < len(table.Robots); {
		role := table.Robots[i]
		role.Reset()
		if role.TotalRound > rand.Int31n(64)+10 ||
			role.Coin < room.Config.PlayMin ||
			role.Coin > room.Config.PlayMax ||
			role.TotalWin > 10000*100 {
			ids = append(ids, role.Id)
			table.Robots = append(table.Robots[:i], table.Robots[i+1:]...)
			role.Online = false
			role.table = nil
			role.player = nil
		} else {
			i++
		}
	}
	if len(ids) > 0 {
		db.Driver.UnloadRobot(room.RoomId, ids)
	}
}

// 结算结果发给玩家
func (table *Table) sendDealResult() {
	if len(table.Roles) > 0 {
		round := table.round
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
}

// 结算
func (table *Table) balance() {
	// 结算真人
	for _, role := range table.Roles {
		role.Balance()
	}
	// 结算机器人
	for _, role := range table.Robots {
		role.Balance()
	}
}

// 准备
func (table *Table) gameReady() {
	table.delay = 5
	table.State = GameStateReady
	table.CurId += 1
	log.Debugf("%v准备:%v", gameName, table.CurId)
}

// 开始下注
func (table *Table) gameOpen() {
	// 发送开始下注消息给所有玩家
	table.delay = 15
	table.State = GameStatePlaying
	log.Debugf("开始下注:%v", table.CurId)

	table.newGameRound()
	table.SendToAll(&folks.OpenBetAck{
		Id:   table.CurId,
		Time: table.delay,
		Rich: table.GetRichPlayer(),
	})
}

// 下注中，每秒调用1次
func (table *Table) gamePlay() {
	// TODO: 需要优化机器人的投注项选择
	for _, role := range table.Robots {
		if rand.Int31n(4) != 1 {
			continue
		}
		betIndex := rand.Intn(len(betItems))
		bet := folks.BetReq{
			Item: robotRandBetItem(),
			Bet:  betItems[betIndex],
		}
		for addCount := rand.Intn(3); addCount >= 0; addCount-- {
			if role.RobotCanBet(bet.Item, bet.Bet) {
				role.AddBet(bet)
				//log.Debugf("R%v下注:%v_%v,%v", role.Id, bet.Item, bet.Bet/100, role.Coin/100)
			}
		}
	}

	l := len(table.round.Flow)
	if l > table.lastFlow {
		// 发送这段时间其他玩家的下注数据
		table.SendToAll(&folks.UserBetAck{
			Time: table.delay,
			Bet:  table.round.Flow[table.lastFlow:l],
		})
		table.lastFlow = l
	}
}

// 发牌结算
func (table *Table) gameDeal() {
	table.delay = 5
	table.State = GameStateDeal
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌
	table.Dealer.Deal(table)
	// 结算
	table.balance()
	//
	round := table.round
	table.LastId = table.CurId
	round.End = room.Now()
	room.SaveLog(round)

	// 最后60局的对战日志
	table.Log = append(table.Log, round.Poker...)
	if over := len(table.Log) - 60*betItemCount; over > 0 {
		table.Log = table.Log[over:]
	}
	// 结算结果发给玩家
	table.sendDealResult()
	// 清理离线玩家
	table.clearOffline()

	log.Debugf("总下注:%v", round.Group)
}
