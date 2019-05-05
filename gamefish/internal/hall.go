package internal

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol/fish"
	"local.com/abc/game/room"
)

var (
	gameName     = "大闹天宫"  // 游戏名称
	betItemCount = 10      // 可投注的项
	betItems     []int32    // 可投注的项
	taxRate      []int64    // 税率千分比
	gameRand     *rand.Rand // 游戏随机生成器
	waitRoles    []*Role    // 等待配桌的玩家
	waitRobots   []*Role    // 等待配桌的机器人
	robotCount   int        // 机器人数量
	idleTables   []*Table   // 空闲的桌子
	runTables    []*Table   // 游戏中的桌子
)

// 桌面对战游戏大厅
type gameHall struct {
	t *time.Ticker
}

func NewGame() room.Haller {
	g := &gameHall{
		t: time.NewTicker(10 * time.Second),
	}
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

func PutWaitRobot(role *Role) {
	waitRobots = append(waitRobots, role)
}

func (hall *gameHall) Update() {
	for i := 0; i < len(runTables); {
		t := runTables[i]
		t.Update()
		// 回收无人的桌子
		if t.RoleCount == 0 {
			t.clear()
			runTables = append(runTables[:i], runTables[i+1:]...)
			idleTables = append(idleTables, t)
		} else {
			i++
		}
	}

	select {
	case <-hall.t.C:
		roleCount := room.UserCount()
		planCount := room.PlanRobotCount(roleCount)
		loadRobots(planCount - robotCount)
	default:
	}

	//if len(waitRoles) > 0 {
	matchRoles()
	//}
}

func matchRoles() {
	//log.Debugf("开始配桌")
	// 先把等待中的桌子座满
	for _, t := range runTables {
		if t.RoleCount < chairCount {
			// 随机少配1个玩家
			remain := chairCount - t.RoleCount - gameRand.Int31n(1)
			for ; remain > 0; remain-- {
				role := PopWaitRole()
				if role == nil {
					role = PopWaitRobot()
					if role == nil {
						log.Debugf("[%v]配桌失败:%v", t.Id, remain)
						return
					}
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

	// 测试，配1桌机器人
	if len(waitRobots) > 0 && len(runTables) == 0 {
		// 获取空闲桌子
		idleLen := len(idleTables)
		if idleLen == 0 {
			return
		}
		t := idleTables[idleLen-1]
		idleTables = idleTables[:idleLen-1]
		runTables = append(runTables, t)
		for i := 0; i < chairCount; i++ {
			role := PopWaitRobot()
			if role == nil {
				break
			}
			t.addRole(role)
		}
		t.Start()
	}
}

func unloadRobots(ids []model.UserId) {
	if len(ids) > 0 {
		robotCount -= len(ids)
		db.Driver.UnloadRobots(room.RoomId, ids)
	}
}

func loadRobots(count int) {
	if count < 0 {
		// 退出机器人
		dec := int(-count)
		if dec > len(waitRobots) {
			dec = len(waitRobots)
		}
		if dec == 0 {
			return
		}
		temp := waitRobots[:dec]
		waitRobots = waitRobots[dec:]
		ids := make([]model.UserId, 0, len(temp))
		for i, r := range temp {
			ids[i] = r.Id
		}
		unloadRobots(ids)
	} else if count > 0 {
		// 增加机器人
		robots := db.Driver.LoadRobots(room.RoomId, count)
		for _, user := range robots {
			sess := &room.Session{
				Ip:      user.Ip,
				Created: time.Now(),
				AgentId: 0,
			}
			coin := rand.Int63n(10*room.Config.PlayMin) + (2 * room.Config.PlayMin)
			user.Coin = coin
			robot := &Role{
				User:    *user,
				Session: sess,
			}
			robot.Online = true
			robotCount++
			waitRobots = append(waitRobots, robot)
			log.Debugf("add robot: %v, coin:%v", user, coin)
		}
	}
}

func (hall *gameHall) Start() {
	gameRand = room.NewRand()

	//// 获取参数
	//switch room.KindId {
	//case model.GameKind_ZJH:
	//	betItemCount = 6
	//	switch room.Config.Level {
	//	case 3:
	//	default:
	//		betItems = []int32{100, 200, 400, 600, 800, 1000}
	//		maxBetItem = 1000 * 2
	//	}
	//}

	// 清理并加载机器人
	db.Driver.ClearRobot(room.RoomId)
	//robotCount := room.PlanRobotCount(0)
	loadRobots(8)

	// 注册消息和事件
	room.RegistEvent(room.EventConfigChanged, configChange)
	room.RegistMsg(int32(fish.Code_CodeTimeSyncAck), (*fish.TimeSyncAck)(nil))
	room.RegistMsg(int32(fish.Code_CodeGameInitAck), (*fish.GameInitAck)(nil))
	room.RegistMsg(int32(fish.Code_CodeGameUpdateAck), (*fish.GameUpdateAck)(nil))

	room.RegistHandler(int32(fish.Code_CodeTimeSyncReq), (*fish.TimeSyncReq)(nil), timeSync)
	room.RegistHandler(int32(fish.Code_CodeShootReq), (*fish.ShootReq)(nil), shoot)
	room.RegistHandler(int32(fish.Code_CodeHitReq), (*fish.HitReq)(nil), hit)

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
	waitRoles = append(waitRoles, role)
}

// 用户重新连接
func (hall *gameHall) UserReline(oldSess *room.Session, newSess *room.Session) {
	if role, ok := newSess.Role.(*Role); ok && role != nil {
		role.Online = true
		role.Session = newSess
		if table := role.table; table != nil {
			log.Debugf("reline:%#v", role)
			table.sendGameInit(role)
		}
	}
}

// 用户下线
func (hall *gameHall) UserOffline(sess *room.Session) {
	if role, ok := sess.Role.(*Role); ok && role != nil {
		role.Online = false
		if table := role.table; table != nil {
			if table.removeRole(role) {
				room.RemoveUser(sess)
			}
		} else {
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

// 同步时间
func timeSync(m *room.NetMessage) {
	log.Debugf("timeSync:%v", m)
	role, ok := m.Session.Role.(*Role)
	if ok == false || role == nil {
		log.Debugf("timeSync: role is nil")
		return
	}
	//获取参数
	req, ok := m.Arg.(*fish.TimeSyncReq)
	if ok == false || req == nil {
		log.Debugf("timeSync: is nil")
		return
	}
	if role.table == nil {
		return
	}
	role.UnsafeSend(&fish.TimeSyncAck{
		Client: req.Tick,
		Server: role.table.tickMs(),
	})
}

// 发子弹
func shoot(m *room.NetMessage) {
	log.Debugf("shoot:%v", m)
	role, ok := m.Session.Role.(*Role)
	if ok == false || role == nil {
		log.Debugf("shoot: role is nil")
		return
	}
	//获取参数
	req, ok := m.Arg.(*fish.ShootReq)
	if ok == false || req == nil {
		log.Debugf("shoot: is nil")
		return
	}
	if err := role.shoot(req); err != nil {
		role.SendError(int32(fish.Code_CodeShootReq), 0, err.Error(), "")
	}
}

// 命中
func hit(m *room.NetMessage) {
	log.Debugf("hit:%v", m)
	role, ok := m.Session.Role.(*Role)
	if ok == false || role == nil {
		log.Debugf("hit: role is nil")
		return
	}
	//获取参数
	req, ok := m.Arg.(*fish.HitReq)
	if ok == false || req == nil {
		log.Debugf("hit: is nil")
		return
	}
	role.hit(req.BulletId, req.FishId)
}
