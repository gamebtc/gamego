package internal

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)

var (
	newDriver   func()GameDriver
	betItem     int     // 可投注的项
	taxRate     []int64 // 税率千分比
	gameName    string  // 游戏名称
	schedule    []Plan
	mustWinRate int32   // 必赢概率百分比
	mustWinRand *rand.Rand
)

const second = 1000 * time.Millisecond
type GameRound = msg.FolksGameRound

func newRand()*rand.Rand{
	bin := make([]byte, 8)
	crand.Read(bin)
	seed := binary.LittleEndian.Uint64(bin)
	return rand.New(rand.NewSource(int64(seed)))
}

func MustWin()bool {
	r := mustWinRand.Int31n(100)
	return mustWinRate > r
}

// 百人游戏(龙虎/红黑/百家乐/色子)
type folksGame struct {
	config *model.RoomInfo
	room.DefaultRoomer
	Tables []*Table
}

func NewGame() room.Roomer {
	g := &folksGame{
		Tables: make([]*Table, 0, 1),
	}
	return g
}

func (this *folksGame) Update() {
	for _, v := range this.Tables {
		v.Update()
	}
}

func (this *folksGame) Init(config *model.RoomInfo) {
	this.config = config
	mustWinRate = config.WinRate
	mustWinRand = newRand()

	switch config.Kind {
	case model.GameKind_BJL:
		schedule = bjlSchedule
		newDriver = NewBjlDealer
		gameName = "百家乐"
		betItem = 5
		taxRate = []int64{0, 50, 50, 50, 50}
	case model.GameKind_HHDZ:
		schedule = rbdzSchedule
		newDriver = NewRbdzDealer
		gameName = "红黑大战"
		betItem = 3
		taxRate = []int64{50, 50, 50}
	case model.GameKind_LHDZ:
		schedule = lhdzSchedule
		newDriver = NewLhdzDealer
		gameName = "龙虎大战"
		betItem = 3
		taxRate = []int64{50, 50, 50}
	case model.GameKind_SBAO:
		schedule = sbaoSchedule
		newDriver = NewSbaoDealer
		gameName = "骰宝"
		betItem = 31
		taxRate = []int64{
			50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
			50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
			50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
		}
	}

	table := NewTable()
	this.Tables = append(this.Tables, table)
	table.Init()

	this.DefaultRoomer.Init(config)
	this.EventHandler[room.EventConfigChanged] = configChange
	this.RegistHandler(msg.MsgId_BetReq, &msg.BetReq{}, betReq)

	//this.EventHandler[room.EventRoomClose] = roomClose
	//room.RegistMsg(msg.MsgId_BetAck, &msg.BetAck{})
	//room.RegistMsg(msg.MsgId_FolksGameInitAck, &msg.FolksGameInitAck{})
	//room.RegistMsg(msg.MsgId_UserBetAck, &msg.UserBetAck{})
	//room.RegistMsg(msg.MsgId_OpenBetAck, &msg.OpenBetAck{})
	//room.RegistMsg(msg.MsgId_CloseBetAck, &msg.CloseBetAck{})

	room.Call(table.Start)
}

//// 创建游戏角色
//func NewGameData(r *model.User) *Role {
//	data := new(Role)
//	data.Id, data.Name, data.Coin, data.Job = r.Id, r.Name, r.Coin, r.Job
//	return data
//}

// 用户上线
func (this *folksGame) UserOnline(sess *room.Session, user *model.User, coin int64) {
	table := this.Tables[0]

	role := &Role{
		User:   user,
		Coin:   coin,
		table:  table,
		Online: true,
		Sender: sess,
	}
	sess.Role = role

	// 发送登录游戏信息
	sess.UnsafeSend(&msg.LoginRoomAck{
		Room: room.RoomId,
		Kind: room.KindId,
		Tab:  table.Id,
	})
	table.AddRole(role)
	// 发送游戏内容
}

// 用户下线
func (this *folksGame) UserOffline(sess *room.Session) {
	if data, ok := sess.Role.(*Role); ok && data != nil {
		data.Online = false

	}
}

// 用户重新连接
func (this *folksGame) UserReline(oldSess *room.Session, newSess *room.Session) {
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
	for i := 0; i < betItem; i++ {
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

// 去掉数组结尾的0
func TrimEndZero(a []int64) []int64 {
	for i := len(a); i > 0; i-- {
		if a[i-1] > 0 {
			return a[:i]
		}
	}
	return a
}