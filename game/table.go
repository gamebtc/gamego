package main

import (
	"math/rand"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)

const RicherCount = 6

var betItem int                // 可投注的项

type Dealer interface {
	// 随机发牌
	Deal() ([]byte, []byte)
	// 比较大小并返回赔率
	Pk(a []byte, b []byte) (odds []int32)
	// 税率千分比
	Tax(i int) int64
	// 可投注的项
	BetItem() int
	// 控制发牌
	Control(interface{}) ([]byte, []byte)
}

func newDealer(config *model.RoomInfo) Dealer {
	switch config.Kind {
	case model.GameKind_HHDZ:
		return NewRbdzDealer(config)
	case model.GameKind_LHDZ:
		return NewLhdzDealer(config)
	}
	return nil
}

// 开始下注：下注时间12秒
// 停止下注：发牌开：2秒
// 结算：    2秒
// 游戏桌子
type Table struct {
	Id      int32              //桌子ID
	CurId   int32              //当前局的ID
	LastId  int32              //最后的局ID
	Log     []byte             //最后60局的输赢情况0:龙赢,1:虎赢,2:和
	State   int32              //0:暂停;1:洗牌;2:下注;3:结算
	Richer  [RicherCount]*Role //富豪ID,位置0为赌神ID
	Roles   map[int32]*Role    //所有游戏玩家，含机器人
	round   *msg.FolksGameRound
	dealer  Dealer
}

func NewTable(config *model.RoomInfo) *Table {
	dealer := newDealer(config)
	betItem = dealer.BetItem()
	t := &Table{
		Roles:  make(map[int32]*Role, 100),
		dealer: dealer,
	}
	return t
}

// 增加新的玩家
func (table *Table) AddRole(role *Role) {
	table.Roles[role.Id] = role
	// 真实玩家
	if role.session != nil {
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
		role.session.Send(ack)
	}
}

func (table *Table) Init() {
}

func (table *Table) NewGameRound() {
	id := room.NewGameRoundId()
	round := &msg.FolksGameRound{
		Id:    id,
		Start: room.Now(),
		Room:  room.RoomId,
		Tab:   table.Id,
		Group: make([]int64, betItem),
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

func (table *Table) Start() {
	table.startRobot()
	table.gameReady()
}

func (table *Table) startRobot() {
	table.addRobot(5593465379, 4743924)
	table.addRobot(5593465381, 2743405)
	table.addRobot(5593465378, 2223178)
	table.addRobot(5593465380, 1775005)
}

func (table *Table) addRobot(agentid int64, userid int32) {
	if user := room.LoadRobot(agentid, userid); user != nil {
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

// 准备
func (table *Table) gameReady() {
	table.State = 1
	room.AfterCall(second, table.openBet)
	table.CurId += 1

	// 重置玩家下注信息
	for _, role := range table.Roles {
		role.Reset()
	}

	table.NewGameRound()
}

// 开始
func (table *Table) openBet() {
	// 发送开始下注消息给所有玩家
	table.State = 2
	log.Debugf("开始下注:%v", table.CurId)
	table.robotBet(0)
}

func (table *Table) robotBet(i int32) {
	if i >= 10 {
		table.closeBet()
	} else {
		i += 1
		room.AfterCall(second, func() { table.robotBet(i) })
		for _, role := range table.Roles {
			if role.session == nil && (role.Id%5) == rand.Int31n(5) {
				bet := msg.BetReq{
					Item: rand.Int31n(3),
					Coin: 100 + rand.Int31n(100)*100,
				}
				if role.RobotCanBet(bet.Item) {
					role.AddBet(bet)
					//log.Debugf("机器人下注:%v,%v", bet, r)
				}
			}
		}
	}
}

// 结算
func (table *Table) closeBet() {
	table.State = 3
	room.AfterCall(second, table.gameReady)
	//log.Debugf("停止下注:%v", table.CurId)
	// 统计下注信息

	// 发牌PK
	a, b := table.dealer.Deal()
	note := model.PokerArrayString(a, "") + "|" + model.PokerArrayString(b, "")
	round := table.round
	round.Odds = table.dealer.Pk(a, b)
	round.Poker = []byte{a[0], b[0]}
	round.Note = note
	//log.Debugf("发牌:%v,%v", note, round.Odds)

	for _, role := range table.Roles {
		if flow := role.Balance(); flow != nil {
			room.WriteCoin(flow)
			if role.session != nil {
				log.Debugf("结算:%v", flow)
			}
		}
	}
	// 结算结果发给玩家
	table.LastId = table.CurId
	round.End = room.Now()

	room.SaveLog(round)

	// 添加发牌日志

	//log.Debugf("round:%v", round)
}

func (table *Table) Update() {

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
