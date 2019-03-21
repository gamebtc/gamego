package internal

import (
	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
	"math/rand"
)

var (
	gameName = "炸金花" // 游戏名称
	betItemCount int     // 可投注的项
	betItems     []int32 // 可投注的项
	taxRate      []int64 // 税率千分比
	mustWinRate  int32   // 必赢概率百分比
	gameRand     *rand.Rand
	roles  []*Role    // 所有玩家
	robots []*Role    // 所有机器人
	waits  []*Role    // 等待配桌
	idleTables []*Table // 空闲的桌子
	runTables  []*Table // 游戏中的桌子
)

type GameRound struct {
	zjh.GameRound
	Log []interface{} // 日志
}

// 桌面对战游戏大厅(斗地主/炸金花/抢庄牛/五张/德州)
type gameHall struct {
}

func NewGame() room.Haller {
	g := &gameHall{
	}
	return g
}

func (hall *gameHall) Update() {
	//select {
	//case <-hall.ticker.C:
	//      // TODO：配桌
	//default:
	//}

	for i := 0; i < len(runTables); {
		v := runTables[i]
		v.Update()
		if v.TryFree() {
			runTables = append(runTables[:i], runTables[i+1:]...)
			idleTables = append(idleTables, v)
		} else {
			i++
		}
	}
}

func (hall *gameHall) Start() {
	config := &room.Config
	mustWinRate = config.WinRate
	gameRand = room.NewRand()

	switch room.KindId {
	case model.GameKind_ZJH:
		betItemCount = 6
		switch config.Level {
		case 3:
		default:
			betItems = []int32{100, 200, 400, 600, 800, 1000}
		}
	}
	db.Driver.ClearRobot(room.RoomId)

	room.RegistEvent(room.EventConfigChanged, configChange)
	room.RegistMsg(int32(zjh.Zjh_ActionAllinAck), &zjh.ActionAllinAck{})

	for i := int32(0); i < config.Tab; i++ {
		table := NewTable()
		table.Id = i+1
		idleTables = append(idleTables, table)
		table.Init()
	}

	// room.Call(table.Start)
}

// 用户上线
func (hall *gameHall) UserOnline(sess *room.Session, user *model.User, coin int64) {
	role := &Role{
		User:   user,
		Coin:   coin,
		Online: true,
		Session: sess,
	}
	sess.Role = role
	// 发送登录游戏信息
	sess.UnsafeSend(&protocol.LoginRoomAck{
		Room: room.RoomId,
		Kind: room.KindId,
		Tab:  0,
	})
}

// 用户下线
func (hall *gameHall) UserOffline(sess *room.Session) {
	if role, ok := sess.Role.(*Role); ok && role != nil {
		role.Online = false

	}
}

// 用户重新连接
func (hall *gameHall) UserReline(oldSess *room.Session, newSess *room.Session) {
	if role, ok := oldSess.Role.(*Role); ok && role != nil {
		oldSess.Role = nil
		role.Online = true
		newSess.Role = role
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
	mustWinRate = arg.WinRate
	room.Config = *arg
	//oldSess := arg[0]
	//newSess := arg[1]
	//oldSess.Role.Role = nil
	//newSess.Role.Role = nil

	//
}

// 房间关闭通知
func roomClose(event *room.GameEvent) {
	args := event.Arg.(*model.RoomInfo)
	args.Pause = 0
}


