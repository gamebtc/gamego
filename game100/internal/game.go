package internal

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
	"local.com/abc/game/room"
)

var (
	newDriver    func()GameDriver
	betItemCount int     // 可投注的项
	betItems     []int32 // 可投注的项
	robotBetRate []int32 // 机器人投注的概率
	robotSumRate int32   // 机器人投注项的概率总和
	badBet       []byte  // 不相容的投注项
	taxRate      []int64 // 税率千分比
	gameName     string  // 游戏名称
	schedule     []Plan
	mustWinRate  int32   // 必赢概率百分比
	mustWinRand  *rand.Rand
)

const second = 1000 * time.Millisecond
type GameRound = protocol.FolksGameRound

func newRand()*rand.Rand{
	bin := make([]byte, 8)
	crand.Read(bin)
	seed := binary.LittleEndian.Uint64(bin)
	return rand.New(rand.NewSource(int64(seed)))
}

func robetRandBetItem() int32 {
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
	config *model.RoomInfo
	room.DefaultRoomer
	tables []*Table
}

func NewGame() room.Roomer {
	g := &gameHall{
		tables: make([]*Table, 0, 1),
	}
	return g
}

func (this *gameHall) Update() {
	for _, v := range this.tables {
		v.Update()
	}
}

func (this *gameHall) Init(config *model.RoomInfo) {
	this.config = config
	mustWinRate = config.WinRate
	mustWinRand = newRand()

	switch config.Kind {
	case model.GameKind_BJL:
		schedule = bjlSchedule
		newDriver = NewBjlDealer
		gameName = "百家乐"
		betItemCount = 5
		taxRate = []int64{0, 50, 50, 50, 50}
		robotBetRate = []int32{89, 90, 7, 2, 2}
		badBet = []byte{1, 0}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
		//betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100, 1000 * 100, 5000 * 100, 10000 * 100}
	case model.GameKind_HHDZ:
		schedule = rbdzSchedule
		newDriver = NewRbdzDealer
		gameName = "红黑大战"
		betItemCount = 3
		taxRate = []int64{50, 50, 50}
		robotBetRate = []int32{92, 92, 8}
		badBet = []byte{1, 0}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
	case model.GameKind_LHDZ:
		schedule = lhdzSchedule
		newDriver = NewLhdzDealer
		gameName = "龙虎大战"
		betItemCount = 3
		taxRate = []int64{50, 50, 50}
		robotBetRate = []int32{92, 92, 8}
		badBet = []byte{1, 0}
		betItems = []int32{1 * 100, 10 * 100, 50 * 100, 100 * 100, 500 * 100}
	case model.GameKind_SBAO:
		schedule = sbaoSchedule
		newDriver = NewSbaoDealer
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

	robotSumRate = robotBetRate[0]
	for i := 1; i < len(robotBetRate); i++ {
		robotSumRate += robotBetRate[i]
		robotBetRate[i] = robotSumRate
	}

	db.Driver.ClearRobot(config.Id)

	table := NewTable()
	this.tables = append(this.tables, table)
	table.Init()

	this.DefaultRoomer.Init(config)
	this.EventHandler[room.EventConfigChanged] = configChange
	this.RegistHandler(protocol.MsgId_BetReq, &protocol.BetReq{}, betReq)

	//this.EventHandler[room.EventRoomClose] = roomClose
	//room.RegistMsg(protocol.MsgId_BetAck, &protocol.BetAck{})
	//room.RegistMsg(protocol.MsgId_FolksGameInitAck, &protocol.FolksGameInitAck{})
	//room.RegistMsg(protocol.MsgId_UserBetAck, &protocol.UserBetAck{})
	//room.RegistMsg(protocol.MsgId_OpenBetAck, &protocol.OpenBetAck{})
	//room.RegistMsg(protocol.MsgId_CloseBetAck, &protocol.CloseBetAck{})

	room.Call(table.Start)
}

//// 创建游戏角色
//func NewGameData(r *model.User) *Role {
//	data := new(Role)
//	data.Id, data.Name, data.Coin, data.Job = r.Id, r.Name, r.Coin, r.Job
//	return data
//}

// 用户上线
func (this *gameHall) UserOnline(sess *room.Session, user *model.User, coin int64) {
	table := this.tables[0]

	role := &Role{
		User:   user,
		Coin:   coin,
		table:  table,
		Online: true,
		Sender: sess,
	}
	sess.Role = role

	// 发送登录游戏信息
	sess.UnsafeSend(&protocol.LoginRoomAck{
		Room: room.RoomId,
		Kind: room.KindId,
		Tab:  table.Id,
	})
	table.AddRole(role)
	// 发送游戏内容
}

// 用户下线
func (this *gameHall) UserOffline(sess *room.Session) {
	if data, ok := sess.Role.(*Role); ok && data != nil {
		data.Online = false

	}
}

// 用户重新连接
func (this *gameHall) UserReline(oldSess *room.Session, newSess *room.Session) {
	if data, ok := oldSess.Role.(*Role); ok && data != nil {
		oldSess.Role = nil
		data.Online = true
		newSess.Role = data
	}
}

// 房间配置更改
func configChange(event *room.GameEvent) {
	args := event.Arg.(*model.RoomInfo)

	args.Pause = 0
	mustWinRate = args.WinRate
	//oldSess := args[0]
	//newSess := args[1]
	//oldSess.Role.Role = nil
	//newSess.Role.Role = nil

	//
}

// 房间关闭通知
func roomClose(event *room.GameEvent) {
	args := event.Arg.(*model.RoomInfo)
	args.Pause = 0
}

// 预算输赢(prize:扣税前总返奖，tax:总税收，bet:总下注)
func Balance(group []int64, odds []int32)(prize, tax, bet int64) {
	for i := 0; i < betItemCount; i++ {
		// 下注金额大于0
		if b := group[i]; b > 0 {
			bet += b
			//有钱回收,包含输1半
			if odd := int64(odds[i]); odd > lostRadix {
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