package internal

import (
	"math/rand"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol/folks"
	"local.com/abc/game/room"
)

var (
	gameRand     *rand.Rand
	newDealer    func() Dealer
	betItemCount int     // 可投注的项
	betItems     []int32 // 可投注的项
	robotBetRate []int32 // 机器人投注的概率
	robotSumRate int32   // 机器人投注项的概率总和
	badBet       []byte  // 不相容的投注项
	taxRate      []int64 // 税率千分比
	gameName     string  // 游戏名称

	// 最多输的倍数，百人牛为10，其它为1
	multipleLost int64 = 1
	// 大厅
	hall = &gameHall{tables: make([]*Table, 0, 1)}
)

type GameRound = folks.GameRound

func robotRandBetItem() int32 {
	r := rand.Int31n(robotSumRate)
	for i, v := range robotBetRate {
		if v > r {
			return int32(i)
		}
	}
	return 0
}

// 百人游戏(龙虎/红黑/百家乐/色子)
type gameHall struct {
	tables []*Table
}

func NewGame() room.Haller {
	return hall
}

func (hall *gameHall) Update() {
	for _, v := range hall.tables {
		v.update()
	}

	if room.Config.Close > 0 {
		canClose := true
		for _, v := range hall.tables {
			if v.State != GameStateReady {
				canClose = false
				break
			}
		}
		if canClose {
			room.Close()
		}
	}
}

func (hall *gameHall) Start() {
	gameRand = room.NewRand()
	switch room.KindId {
	case model.GameKind_BJL:
		newDealer = NewBjlDealer
		gameName = "百家乐"
		betItemCount = 5
		taxRate = []int64{0, 50, 50, 50, 50}
		robotBetRate = []int32{89, 90, 7, 2, 2}
		badBet = []byte{1, 0}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
		//betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100, 1000 * 100, 5000 * 100, 10000 * 100}
	case model.GameKind_BRNN:
		newDealer = NewBjlDealer
		gameName = "百人牛牛"
		betItemCount = 4
		taxRate = []int64{50, 50, 50, 50}
		robotBetRate = []int32{100, 100, 100, 100}
		badBet = []byte{1, 0, 3, 2}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
		multipleLost = 10
	case model.GameKind_HHDZ:
		newDealer = NewRbdzDealer
		gameName = "红黑大战"
		betItemCount = 3
		taxRate = []int64{50, 50, 50}
		robotBetRate = []int32{92, 92, 8}
		badBet = []byte{1, 0}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
	case model.GameKind_LHDZ:
		newDealer = NewLhdzDealer
		gameName = "龙虎大战"
		betItemCount = 3
		taxRate = []int64{50, 50, 50}
		robotBetRate = []int32{92, 92, 8}
		badBet = []byte{1, 0}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
	case model.GameKind_SBAO:
		newDealer = NewSbaoDealer
		gameName = "骰宝"
		betItemCount = 31
		taxRate = []int64{
			50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
			50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
			50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
		}
		robotBetRate = []int32{
			200, 200, 200, 200, 20, 20, 20, 20, 20, 20,
			1, 1, 1, 1, 1, 1, 1, 3, 5, 6,
			7, 8, 9, 9, 9, 9, 8, 7, 6, 5, 3,
		}
		badBet = []byte{1, 0, 3, 2}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
	}

	// 预算概率
	robotSumRate = robotBetRate[0]
	for i := 1; i < len(robotBetRate); i++ {
		robotSumRate += robotBetRate[i]
		robotBetRate[i] = robotSumRate
	}

	// 清理机器人
	db.Driver.ClearRobot(room.RoomId)
	// 注册消息和事件
	room.RegistMsg(int32(folks.Folks_UserBetAck), (*folks.UserBetAck)(nil))
	room.RegistMsg(int32(folks.Folks_OpenBetAck), (*folks.OpenBetAck)(nil))
	room.RegistMsg(int32(folks.Folks_StopBetAck), (*folks.StopBetAck)(nil))
	room.RegistMsg(int32(folks.Folks_GameInitAck), (*folks.GameInitAck)(nil))
	room.RegistMsg(int32(folks.Folks_BetAck), (*folks.BetAck)(nil))

	room.RegistEvent(room.EventConfigChanged, configChange)
	room.RegistHandler(int32(folks.Folks_BetReq), (*folks.BetReq)(nil), betReq)

	// 创建桌子
	table := NewTable()
	table.Id = 1
	hall.tables = append(hall.tables, table)
	table.Init()
	room.Call(table.Start)
}

// 用户上线
func (hall *gameHall) UserOnline(sess *room.Session, user *model.User) {
	// 分配桌子
	table := hall.tables[0]
	role := &Role{
		User:    *user,
		Session: sess,
	}
	sess.Role = role
	role.Online = true
	table.addRole(role)
	table.sendGameInit(role)
}

// 用户重新连接
func (hall *gameHall) UserReline(oldSess *room.Session, newSess *room.Session) {
	if role, ok := newSess.Role.(*Role); ok && role != nil {
		role.Online = true
		role.Session = newSess
		if table := role.table; table != nil {
			log.Debugf("reline:old%v,new:%v,uid:%v", oldSess.AgentId, newSess.AgentId, role.UserId)
			table.sendGameInit(role)
		}
	}
}

// 用户下线
func (hall *gameHall) UserOffline(sess *room.Session) {
	if role, ok := sess.Role.(*Role); ok && role != nil {
		role.Online = false
		if role.bill == nil {
			room.RemoveUser(sess)
		}
	}
}

// 房间配置更改
func configChange(event *room.GameEvent) {
	arg := event.Arg.(*model.RoomInfo)

	// 房间关闭
	if arg.Close > 0 {
		arg.Pause = arg.Close
		arg.Lock = arg.Close
	}
	room.Config = *arg
	//oldSess := arg[0]
	//newSess := arg[1]
	//oldSess.Role.Role = nil
	//newSess.Role.Role = nil
}

// 房间关闭通知
func roomClose(event *room.GameEvent) {
	args := event.Arg.(*model.RoomInfo)
	args.Pause = 0
}

// 预算输赢(prize:扣税前总返奖，tax:总税收，bet:总下注)
func Balance(group []int64, odds []int32) (prize, tax, bet int64) {
	for i := 0; i < betItemCount; i++ {
		// 下注金币大于0
		if b := group[i]; b > 0 {
			bet += b
			//有钱回收,包含输1半
			if odd := int64(odds[i]); odd != lostRadix {
				w := b * odd / radix
				if w > b {
					// 赢钱了收税，税率按千分比配置，需除以1000
					tax += (w - b) * taxRate[i] / 1000
				}
				prize += w
			}
		}
	}
	return
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

	ack := &folks.BetAck{
		Sn:   req.Sn,
		Item: req.Item,
	}
	if err := role.addBet(*req); err != nil {
		role.SendError(int32(folks.Folks_BetReq), 1000, err.Error(), "")
		log.Debugf("add bet error: %v", err)
	} else {
		ack.Bet = req.Bet
		log.Debugf("add bet success: %v", ack)
	}
	ack.Coin = role.Coin
	role.Send(ack)
}
