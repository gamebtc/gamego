package main

import (
	"time"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)


const second = 1000 * time.Millisecond

// 百人游戏(龙虎/红黑/百家乐/色子)
type folksGame struct {
	room.DefaultRoomer
	table *Table
}

func NewGame() room.Roomer {
	g := &folksGame{}
	return g
}

func (this *folksGame) Update() {
	//log.Debugf("update:%v", this.count)
	//for _, v := range room.Tables {
	//	v.Update()
	//}
}

func (this *folksGame) Init(config *model.RoomInfo) {

	table := NewTable(config)
	this.table = table
	table.Init()

	this.DefaultRoomer.Init(config)
	this.EventHandler[room.EventConfigChanged] = configChange
	//this.EventHandler[room.EventRoomClose] = roomClose

	this.RegistHandler(msg.MsgId_BetReq, &msg.BetReq{}, betReq)

	//room.RegistMsg(msg.MsgId_BetAck, &msg.BetAck{})
	//room.RegistMsg(msg.MsgId_FolksGameInitAck, &msg.FolksGameInitAck{})
	//room.RegistMsg(msg.MsgId_UserBetAck, &msg.UserBetAck{})
	//room.RegistMsg(msg.MsgId_OpenBetAck, &msg.OpenBetAck{})
	//room.RegistMsg(msg.MsgId_CloseBetAck, &msg.CloseBetAck{})

	room.Call(table.Start)
}

// 创建游戏角色
func NewGameData(r *model.User) *Role {
	data := new(Role)
	data.Id, data.Name, data.Coin, data.Job = r.Id, r.Name, r.Coin, r.Job
	return data
}

// 用户上线
func (this *folksGame) UserOnline(sess *room.Session, user *model.User) {
	role := &Role{
		User: user,
	}
	role.table = this.table
	role.online = true
	sess.Role = role
	role.session = sess

	// 发送登录游戏信息
	sess.UnsafeSend(&msg.LoginRoomAck{
		Room: room.Config.Id,
		Kind: room.Config.Kind,
		Tab:  this.table.Id,
	})
	role.table.AddRole(role)

	// 发送游戏内容

}

// 用户下线
func (this *folksGame) UserOffline(sess *room.Session) {
	if data, ok := sess.Role.(*Role); ok && data != nil {
		data.online = false
	}
}

// 用户重新连接
func (this *folksGame) UserReline(oldSess *room.Session, newSess *room.Session) {
	if data, ok := oldSess.Role.(*Role); ok && data != nil {
		oldSess.Role = nil
		data.online = true
		newSess.Role = data
	}
}

// 房间配置更改
func configChange(event *room.GameEvent) {
	args := event.Arg.(*model.RoomInfo)

	args.Pause = 0
	//oldSess := args[0]
	//newSess := args[1]
	//oldSess.Role.Role = nil
	//newSess.Role.Role = nil
}

// 房间关闭通知
func roomClose(event *room.GameEvent) {
	args := event.Arg.(*model.RoomInfo)
	args.Pause = 0
}
