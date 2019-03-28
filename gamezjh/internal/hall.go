package internal

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

var (
	gameName     = "炸金花" // 游戏名称
	betItemCount int     // 可投注的项
	betItems     []int32 // 可投注的项
	taxRate      []int64 // 税率千分比
	gameRand     *rand.Rand
	roles        []*Role  // 所有玩家
	robots       []*Role  // 所有机器人
	waitRoles    []*Role  // 等待配桌的玩家
	waitRobots   []*Role  // 等待配桌的机器人
	idleTables   []*Table // 空闲的桌子
	runTables    []*Table // 游戏中的桌子
)

type GameRound = zjh.GameRound

// 桌面对战游戏大厅(斗地主/炸金花/抢庄牛/五张/德州)
type gameHall struct {
}

func NewGame() room.Haller {
	g := &gameHall{}
	return g
}

func PopWaitRole() (role *Role) {
	if len(waitRoles) > 0 {
		role = waitRoles[0]
		waitRoles = waitRoles[1:]
	}
	return
}

func PopWaitRobot() (role *Role) {
	if len(waitRobots) > 0 {
		role = waitRobots[0]
		waitRobots = waitRobots[1:]
	}
	return
}

func (hall *gameHall) Update() {
	for i := 0; i < len(runTables); {
		t := runTables[i]
		t.Update()
		// 回收无人的桌子
		if t.State == GameStateReady && t.RoleCount == 0 {
			runTables = append(runTables[:i], runTables[i+1:]...)
			idleTables = append(idleTables, t)
		} else {
			i++
		}
	}

	if len(waitRoles) > 0 {
		matchRoles()
	}
}

func matchRoles() {
	log.Debugf("开始配桌:%v")
	// 先把等待中的桌子座满
	for _, t := range runTables {
		if t.State != GameStatePlaying && t.RoleCount < chairCount {
			// 随机少配1个玩家
			for remain := chairCount - t.RoleCount - gameRand.Int31n(1); remain > 0; remain-- {
				role := PopWaitRole()
				if role == nil {
					return
				}
				t.addRole(role)
			}
		}
	}

	// 再把游戏中的桌子座满
	if len(waitRoles) > 0 {
		for _, t := range runTables {
			if t.RoleCount < chairCount {
				for remain := chairCount - t.RoleCount; remain > 0; remain-- {
					role := PopWaitRole()
					if role == nil {
						return
					}
					t.addRole(role)
				}
			}
		}
	}

	for len(waitRoles) > 0 {
		// 获取空闲桌子
		idleLen := len(idleTables)
		if idleLen == 0 {
			return
		}
		t := idleTables[idleLen-1]
		idleTables = idleTables[:idleLen-1]
		runTables = append(runTables, t)
		for i := 0; i < chairCount; i++ {
			role := PopWaitRole()
			if role == nil {
				break
			}
			t.addRole(role)
		}
		if t.RoleCount <= 3 {
			// 配0-2个机器人
			for robotCount := gameRand.Int31n(3); robotCount > 0; robotCount-- {
				role := PopWaitRobot()
				if role == nil {
					break
				}
				t.addRole(role)
			}
		}
		t.Start()
	}
}

func loadRobot(count int32) {
	if count < 0 {
		// 退出机器人
		dec := int(-count)
		ids := make([]int32, 0, dec)
		for i := 0; i < len(waitRobots) && i < dec; i++ {
			role := waitRobots[i]
			ids = append(ids, role.Id)
		}
		dec = len(ids)
		if dec > 0 {
			waitRobots = waitRobots[dec:]
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
			coin := rand.Int63n(10*room.Config.PlayMin) + (2 * room.Config.PlayMin)
			robot := &Role{
				User:    *user,
				Session: sess,
			}
			robot.Online = true
			waitRobots = append(waitRobots, robot)
			log.Debugf("add robot: %v, coin:%v", user, coin)
		}
	}
}

func (hall *gameHall) Start() {
	gameRand = room.NewRand()

	// 获取参数
	switch room.KindId {
	case model.GameKind_ZJH:
		betItemCount = 6
		switch room.Config.Level {
		case 3:
		default:
			betItems = []int32{100, 200, 400, 600, 800, 1000}
		}
	}

	// 清理并加载机器人
	db.Driver.ClearRobot(room.RoomId)
	robotCount := room.PlanRobotCount(0)
	loadRobot(robotCount)

	// 注册消息和事件
	room.RegistEvent(room.EventConfigChanged, configChange)
	room.RegistMsg(int32(zjh.Code_CodeActionAck), &zjh.ActionAck{})
	room.RegistMsg(int32(zjh.Code_CodeGameInitAck), &zjh.GameInitAck{})
	room.RegistMsg(int32(zjh.Code_CodeGameStartAck), &zjh.GameStartAck{})
	room.RegistMsg(int32(zjh.Code_CodeGameEndAck), &zjh.GameEndAck{})

	room.RegistHandler(int32(zjh.Code_CodeActionReq), &zjh.ActionReq{}, action)

	// 创建所有的桌子
	tableCount := room.Config.Tab
	runTables = make([]*Table, 0, tableCount)
	idleTables = make([]*Table, 0, tableCount)
	for i := tableCount; i > 0; i-- {
		table := NewTable()
		table.Id = i
		idleTables = append(idleTables, table)
		table.Init()
	}

	// room.Call(table.Start)
}

// 用户上线
func (hall *gameHall) UserOnline(sess *room.Session, user *model.User) {
	role := &Role{
		User:    *user,
		Session: sess,
	}
	role.Online = true
	sess.Role = role
	// 发送登录游戏信息
	sess.UnsafeSend(&protocol.LoginRoomAck{
		Room: room.RoomId,
		Kind: room.KindId,
	})

	waitRoles = append(waitRoles, role)
}

// 用户重新连接
func (hall *gameHall) UserReline(oldSess *room.Session, newSess *room.Session) {
	if role, ok := oldSess.Role.(*Role); ok && role != nil {
		oldSess.Role = nil
		role.Online = true
		newSess.Role = role
		//if table := role.table; table != nil {
		//	table.sendGameInit(role)
		//}
	}
}

// 用户下线
func (hall *gameHall) UserOffline(sess *room.Session) {
	if role, ok := sess.Role.(*Role); ok && role != nil {
		role.Online = false
		if role.bill == nil {
			room.RemoveUser(sess)
			sess.UnlockRoom()
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

// 玩家动作
func action(m *room.NetMessage) {
	log.Debugf("action:%v", m)
	role, ok := m.Session.Role.(*Role)
	if ok == false || role == nil {
		log.Debugf("action: role is nil")
		return
	}
	//获取参数
	req, ok := m.Arg.(*zjh.ActionReq)
	if ok == false || req == nil {
		log.Debugf("action: is nil")
		return
	}

	if role.table == nil {
		return
	}

	if role.table.round == nil {
		return
	}

	switch req.Type {
	case zjh.ActionType_ActionReady: // 准备
	    role.Ready()
	case zjh.ActionType_ActionLook: // 看牌
		role.Look()
	case zjh.ActionType_ActionDiscard: // 主动弃牌
		role.Discard(false)
	case zjh.ActionType_ActionOvertime: // 超时弃牌
		role.Discard(true)
	case zjh.ActionType_ActionCompare: // 比牌
		role.Compare(req.Opponent)
	case zjh.ActionType_ActionAddBet: // 下注(跟注+加注)
		role.AddBet(req.Bet)
	case zjh.ActionType_ActionAllin: // 全压
		role.Allin()
	case zjh.ActionType_ActionRenew: // 换桌玩
		role.RenewDesk()
	}
}

