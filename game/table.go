package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/msg"
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
	Id     int32              //桌子ID
	CurId  int32              //当前局的ID
	LastId int32              //最后的局ID
	Log    []byte             //最后60局的输赢情况0:龙赢,1:虎赢,2:和
	State  int32              //0:暂停;1:洗牌;2:下注;3:结算
	Richer [RicherCount]*Role //富豪ID,位置0为赌神ID
	Roles  map[int32]*Role    //所有游戏玩家，含机器人
	round  *GameRound
	GameDriver

	timer     *time.Timer
	planIndex int
}

func NewTable() *Table {
	dealer := newDriver()
	t := &Table{
		Roles:      make(map[int32]*Role, 100),
		GameDriver: dealer,
	}
	return t
}

// 增加新的玩家
func (table *Table) AddRole(role *Role) {
	table.Roles[role.Id] = role
	// 真实玩家
	if role.Session != nil {
		ack := &msg.FolksGameInitAck{
			Id:    table.round.Id,
			State: table.State,
			Sum:   table.round.Group,
			Log:   table.Log,
		}
		if bill := role.bill; bill != nil {
			ack.Bet = bill.Group
		}
		for i := 0; i < RicherCount; i++ {
			if rich := table.Richer[i]; rich != nil {
				ack.Rich = append(ack.Rich, rich.GetMsgUser())
			}
		}
		role.Session.Send(ack)
	}
}

func (table *Table) NewGameRound() {
	table.State = 1
	table.CurId += 1
	id := room.NewGameRoundId()
	g := make([]int64, 2*betItem)
	round := &GameRound{
		Id:        id,
		Start:     room.Now(),
		Room:      room.RoomId,
		Tab:       table.Id,
		Group:     g[0:betItem],
		UserGroup: g[betItem:],
	}

	for _, role := range table.Roles {
		role.Reset()
	}
	//加入富豪ID
	for i := 0; i < RicherCount; i++ {
		if role := table.Richer[i]; role != nil {
			round.Bill = append(round.Bill, role.bill)
		}
	}
	table.round = round
}

func (table *Table) Init(){
}


func (table *Table) Start() {
	table.startRobot()
	table.timer = time.NewTimer(time.Microsecond)
}

func (table *Table) startRobot() {
	table.addRobot(5593465379, 4743924)
	table.addRobot(5593465381, 2743405)
	table.addRobot(5593465378, 2223178)
	table.addRobot(5593465380, 1775005)
}

func (table *Table) addRobot(agentId int64, userId int32) {
	if user := room.LoadRobot(agentId, userId); user != nil {
		role := &Role{
			User: user,
		}
		role.table = table
		role.online = true
		role.Reset()
		table.AddRole(role)
		log.Debugf("addRobot:%v", role)
	}
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
	prize, _, bet := Balance(table.round.UserGroup, odds)
	return bet - prize
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
	req, ok := m.Arg.(*msg.BetReq)
	if ok == false || req == nil {
		log.Debugf("betReq: is nil")
		return
	}
	log.Debugf("add bet: %v", req)
	role.AddBet(*req)
}

// 准备
func gameReady(table *Table) {
	table.State = 1
	table.NewGameRound()
	table.Ready(table)
}

// 开始
func gameOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = 2
	table.Open(table)
	log.Debugf("开始下注:%v", table.CurId)
}

func gamePlay(table *Table) {
	for _, role := range table.Roles {
		if role.Session == nil && (role.Id%5) == rand.Int31n(5) {
			bet := msg.BetReq{
				Item: rand.Int31n(31),
				Coin: 100 + rand.Int31n(100)*100,
			}
			if role.RobotCanBet(bet.Item) {
				role.AddBet(bet)
				//log.Debugf("机器人下注:%v,%v", bet, r)
			}
		}
	}
}

// 停止下注
func gameStop (table *Table) {
	table.State = 3
	log.Debugf("停止下注:%v", table.CurId)
	table.Stop(table)
}

// 发牌结算
func gameDeal(table *Table){
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌结算
	table.Deal(table)
}