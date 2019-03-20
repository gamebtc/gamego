package internal

import (
	"math/rand"
	"time"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

var (
	betItemCount int     // 可投注的项
	betItems     []int32 // 可投注的项
	badBet       []byte  // 不相容的投注项
	taxRate      []int64 // 税率千分比
	mustWinRate  int32   // 必赢概率百分比
	mustWinRand  *rand.Rand

	gameName = "炸金花" // 游戏名称
)

const second = 1000 * time.Millisecond


type GameRound struct {
	zjh.GameRound
	Log []interface{} // 日志
}

// 桌面对战游戏大厅(斗地主/炸金花/抢庄牛/五张/德州)
type gameHall struct {
	config *model.RoomInfo
	room.DefaultRoomer
	roles  []*Role
	robots []*Role
	waits  []*Role    // 等待配桌
	tables []*Table
	ticker *time.Ticker
}

func NewGame() room.Roomer {
	g := &gameHall{
		tables: make([]*Table, 0, 1),
		ticker :time.NewTicker(time.Second),
	}
	return g
}

func (this *gameHall) Update() {

	select {
	case <-this.ticker.C:
          // TODO：配桌
	default:
	}

	for _, v := range this.tables {
		v.Update()
	}
}

func (this *gameHall) Init(config *model.RoomInfo) {
	this.config = config
	mustWinRate = config.WinRate
	mustWinRand = room.NewRand()

	switch config.Kind {
	case model.GameKind_ZJH:
		//newDriver = ZjhNealer
	}

	db.Driver.ClearRobot(config.Id)

	table := NewTable()
	this.tables = append(this.tables, table)
	table.Init()

	this.DefaultRoomer.Init(config)
	this.EventHandler[room.EventConfigChanged] = configChange

	//this.RegistHandler(protocol.MsgId_BetReq, &protocol.BetReq{}, betReq)
	//this.EventHandler[room.EventRoomClose] = roomClose
	//room.RegistMsg(protocol.MsgId_BetAck, &protocol.BetAck{})
	//room.RegistMsg(protocol.MsgId_FolksGameInitAck, &protocol.FolksGameInitAck{})
	//room.RegistMsg(protocol.MsgId_UserBetAck, &protocol.UserBetAck{})
	//room.RegistMsg(protocol.MsgId_OpenBetAck, &protocol.OpenBetAck{})
	room.RegistMsg(int32(zjh.Zjh_ActionAllinAck), &zjh.ActionAllinAck{})

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
		Session: sess,
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


