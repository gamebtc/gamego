package internal

import (
	"math"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/protocol/fish"
	"local.com/abc/game/room"
)

const chairCount = 4

type GameState int32

type RefreshTroop struct {
	SendDes   bool
	SendTroop bool
	BeginTime float64
}

// 游戏桌子
type Table struct {
	fish.GameUpdateAck
	Id        int32             // 桌子ID
	State     GameState         // 0:准备游戏;1:游戏中;2:结算
	RoleCount int32             // 玩家数量
	Roles     [chairCount]*Role // 所有游戏玩家
	epoch     time.Time         // 游戏开始时间
	lastTick  int64             //
	fishId    int32             // 鱼ID
	bulletId  int32             // 子弹ID
	fishes    map[int32]*Fish   // 所有的鱼
	bullets   map[int32]*Bullet // 所有的子弹
	initAck   fish.GameInitAck

	nearFishPos [chairCount]Point
	startTime   [chairCount]time.Time

	sceneTime          float64      // 当前场景运行时间(秒)
	pauseTime          float64      // 当前场景暂停时间(秒)
	stopFire           bool         // 停止开火
	sceneConfig        *SceneConfig // 当前场景配置
	distributeFishTime []float64    // 生成鱼的时间
	distributeTroop    RefreshTroop // 鱼阵
	canLockList        []int32      // 可以锁定的鱼
	fishCount          int32        // 屏幕内鱼的数量
}

func NewTable() *Table {
	t := &Table{}
	t.Init()
	return t
}

func (table *Table) tickMs() int64 {
	return time.Since(table.epoch).Nanoseconds() / 1000000
}

func (table *Table) newBullId() int32 {
	id := int32(table.bulletId&0X0FFFFFFF) + 1
	table.bulletId++
	return id
}

func (table *Table) newFishId() int32 {
	id := int32(table.fishId&0X0FFFFFFF) + 1
	table.fishId++
	return id
}

// 查找一个空位
func (table *Table) findEmptyChair() int32 {
	for i, v := range table.Roles {
		if v == nil {
			return int32(i)
		}
	}
	return -1
}

// 检查是否有真实玩家
func (table *Table) hasRealPlayer() bool {
	for _, role := range table.Roles {
		if role.isRobot() == false {
			return true
		}
	}
	return false
}

func (table *Table) Init() {
	for i := 0; i < chairCount; i++ {
		table.nearFishPos[i].X = SystemConf.ScreenWidth / 2
		table.nearFishPos[i].Y = SystemConf.ScreenHeight / 2
	}
}

func (table *Table) clear() {
	for i := 0; i < chairCount; i++ {
		if role := table.Roles[i]; role != nil {
			table.removeRole(role)
		}
	}
	if table.fishes != nil {
		for _, f := range table.fishes {
			table.removeFish(f)
		}
	}
	if table.bullets != nil {
		for _, b := range table.bullets {
			table.removeBullet(b)
		}
	}
	table.fishes = nil
	table.bullets = nil
	table.canLockList = nil
	table.initAck = fish.GameInitAck{}
	table.GameUpdateAck = fish.GameUpdateAck{}
	table.sceneConfig = nil
	table.distributeFishTime = nil
}

func (table *Table) Start() {
	const stopFire = true
	table.StopFire = stopFire
	table.fishes = make(map[int32]*Fish, 256)
	table.bullets = make(map[int32]*Bullet, 128)
	table.canLockList = make([]int32, 0, 256)

	ia := &table.initAck
	ia.Table = table.Id
	ia.Players = make([]*fish.Player, 0, 4)
	ia.Fishes = make([]*fish.Fish, 0, 128)
	ia.Bullets = make([]*fish.Bullet, 0, 64)
	ia.StopFire = stopFire
	ia.MaxBullet = maxBulletBet
	ia.FireInterval = fireInterval

	up := &table.GameUpdateAck
	up.StopFire = stopFire
	up.Players = make([]*fish.Player, 0, 4)
	up.Fishes = make([]*fish.Fish, 0, 64)
	up.Bullets = make([]*fish.Bullet, 0, 16)
	up.DieBullets = make([]int32, 0, 16)
	up.DieFishes = make([]int32, 0, 16)
	up.Kills = make([]*fish.KillFish, 0, 16)
	up.Seed = make([]*fish.FishSeed, 0, 16)
	up.Offline = make([]int32, 0, 4)
	up.Describe = make([]string, 0, 2)
	table.epoch = time.Now()
	table.lastTick = 0
	table.sceneTime = 0
	table.pauseTime = 0
	table.fishCount = 0
	table.stopFire = true

	table.sceneConfig = &SceneConfigs[0]
	table.resetDistribute()
}

func (table *Table) resetDistribute() {
	table.distributeFishTime = make([]float64, len(table.sceneConfig.DistributeFish))
	table.distributeTroop = RefreshTroop{}
}

// 增加玩家
func (table *Table) addRole(role *Role) bool {
	i := table.findEmptyChair()
	if i < 0 {
		return false
	}
	role.Player = &fish.Player{
		Id:    role.Id,
		Icon:  role.Icon,
		Vip:   role.Vip,
		Name:  role.Name,
		Coin:  role.Coin,
		Chair: i,
	}
	role.round = &fish.GameRound{
		Uid:     role.Id,
		Start:   room.Now(),
		Room:    room.RoomId,
		Tab:     table.Id,
		OldCoin: role.Coin,
		Log:     make([]int32, 0, 300),
	}
	role.table = table
	table.Roles[i] = role
	table.RoleCount++
	if !role.isRobot() {
		//role.Player.Robot = &RobotAi{}
		table.sendGameInit(role)
	} else {
		//role.Player.Robot = &RobotAi{}
	}
	upAck := &table.GameUpdateAck
	upAck.Players = append(upAck.Players, role.Player)
	// 移除删除
	for i := 0; i < len(upAck.Offline); i++ {
		if upAck.Offline[i] == role.Id {
			upAck.Offline = append(upAck.Offline[:i], upAck.Offline[i+1:]...)
		}
	}
	log.Debugf("[%v]addRole:%v", table.Id, role.Id)
	return true
}

// 删除子弹
func (table *Table) removeBullet(bullet *Bullet) {
	delete(table.bullets, bullet.Id)
	table.GameUpdateAck.DieBullets = append(table.GameUpdateAck.DieBullets, bullet.Id)
	pushBullet(bullet)
}

// 删除鱼
func (table *Table) removeFish(fish *Fish) {
	delete(table.fishes, fish.Id)
	table.GameUpdateAck.DieFishes = append(table.GameUpdateAck.DieFishes, fish.Id)
	pushFish(fish)
}

func (table *Table) removeRole(role *Role) bool {
	chair := role.Player.Chair
	if role != nil && role == table.Roles[chair] {
		// 回收玩家所有子弹
		coin := int64(0)
		for _, v := range table.bullets {
			if v.Uid == role.Id {
				table.removeBullet(v)
				coin += int64(v.Bet)
			}
		}
		role.addCoin(coin)
		if role.round.Bet > 0 && role.round.Win != 0 {
			role.writeCoin()
		}
		table.RoleCount--
		table.Roles[chair] = nil

		role.table = nil
		role.Player = nil
		role.round = nil
		role.bulletCount = 0
		log.Debugf("[%v]removeRole:%v", table.Id, role.Id)
		return true
	}
	return false
}

// 发送初始化场景消息
func (table *Table) sendGameInit(role *Role) {
	if !role.isRobot() {
		ack := &table.initAck
		ack.Tick = table.tickMs()
		ack.Scene = table.sceneConfig.Id
		ack.StopFire = table.StopFire
		for _, role := range table.Roles {
			if role != nil {
				ack.Players = append(ack.Players, role.Player)
			}
		}
		if table.fishes != nil {
			for _, f := range table.fishes {
				ack.Fishes = append(ack.Fishes, &f.Fish)
			}
		}
		if table.bullets != nil {
			for _, b := range table.bullets {
				ack.Bullets = append(ack.Bullets, &b.Bullet)
			}
		}
		role.Send(ack)
		//clear ack
		for i := len(ack.Players) - 1; i >= 0; i-- {
			ack.Players[i] = nil
		}
		for i := len(ack.Fishes) - 1; i >= 0; i-- {
			ack.Fishes[i] = nil
		}
		for i := len(ack.Bullets) - 1; i >= 0; i-- {
			ack.Bullets[i] = nil
		}
		ack.Players = ack.Players[0:0]
		ack.Fishes = ack.Fishes[0:0]
		ack.Bullets = ack.Bullets[0:0]
	}
}

func (table *Table) CatchFish(bullet *Bullet, fish *Fish) {
	pdf := fish.Temp.Probability
	r := gameRand.Float64() * 1000

	if r >= (pdf * SystemConf.RobotProbMul / 3) {
		//这里先简单判断一下
		return
	}
	//table.removeFish(fish)
}

// 更新鱼群
func (table *Table) updateFishes(second float64) {
	table.canLockList = table.canLockList[0:0]
	for _, fish := range table.fishes {
		fish.Update(second)
		if fish.EndPath {
			if fish.InSide {
				// TODO:创建新的移动，一直延原来的方向移动出屏幕
				fish.Path = nil
				fish.initAngle()
			} else {
				table.removeFish(fish)
			}
		}
		if fish.InSide {
			if fish.Temp.LockLevel > 0 {
				table.canLockList = append(table.canLockList, fish.Id)
			}
			table.fishCount++
		}
	}
}

// 更新子弹
func (table *Table) updateBullets(second float64) {
	for _, bullet := range table.bullets {
		bullet.Move(second)
		if bullet.isRobot() == false {
			continue
		}
		// 机器人由服务器检查碰撞
		if bullet.Fish != 0 {
			if fish, ok := table.fishes[bullet.Fish]; ok {
				if fish.State < ObjectState_Dead {
					if bullet.HitTest(fish) {
						table.CatchFish(bullet, fish)
						table.removeBullet(bullet)
					}
				}
				continue
			} else {
				bullet.Fish = 0
			}
		}

		// 循环鱼群，进行碰撞检查
		for _, fish := range table.fishes {
			if fish.State < ObjectState_Dead {
				if bullet.HitTest(fish) {
					table.CatchFish(bullet, fish)
					table.removeBullet(bullet)
					break
				}
			}
		}
	}
}

func (table *Table) DistributeFish(second float64) {

	if table.pauseTime > 0 {
		table.pauseTime -= second
		return
	}

	table.sceneTime += second
	if table.stopFire &&
		table.fishCount > 0 &&
		table.sceneTime > SystemConf.SwitchSceneTime {
		table.stopFire = false
		table.GameUpdateAck.StopFire = false
	}

	conf := table.sceneConfig
	if table.sceneTime < conf.Time {
		if ts := conf.findTroop(table.sceneTime); ts != nil {
			table.troopFish(ts, conf.Time)
			return
		}

		if table.sceneTime > SystemConf.SwitchSceneTime {
			for i := 0; i < len(conf.DistributeFish); i++ {
				if i >= len(table.distributeFishTime) {
					break
				}
				dis := &conf.DistributeFish[i]
				table.distributeFishTime[i] += second
				if table.distributeFishTime[i] > dis.Time {
					table.distributeFishTime[i] -= dis.Time
					table.refreshFish(dis)
				}
				////
			}
		}
	} else {
		// 切换到下一个场景
		table.sceneConfig = findSceneConfig(table.sceneConfig.Next)
		table.resetDistribute()

		table.stopFire = true
		table.GameUpdateAck.SwitchScene = table.sceneConfig.Id
		table.GameUpdateAck.StopFire = true

		// Clear Fish
		for _, fish := range table.fishes {
			table.removeFish(fish)
		}
		table.canLockList = table.canLockList[0:0]
		table.sceneTime = 0
	}
}

func (table *Table) refreshFish(dis *DistributeFishConfig) {
	total := gameRand.Int31n(dis.MaxCount-dis.MinCount+1) + dis.MinCount
	if dis.RefreshType == RefershType_Snak {
		total += 2
	}

	fishType := int32(-1)
	snakeType := int32(0)
	refreshId := table.newFishId()
	pathId := GetRandNormalPathID()
	for i := total; i > 0; i-- {
		if fishType == -1 || dis.RefreshType == RefershType_Normal {
			pathId = GetRandNormalPathID()
			fishType = dis.RandFish()
			snakeType = fishType
		}

		if dis.RefreshType == RefershType_Snak {
			if i == total {
				fishType = SystemConf.SnakeHead
			} else if i == 1 {
				fishType = SystemConf.SnakeTail
			}
		}

		if tmp, ok := fishTemplateMap[fishType]; ok {
			var xOffset, yOffset, delay float64
			if dis.RefreshType == RefershType_Line || dis.RefreshType == RefershType_Snak {
				xOffset = dis.OffsetX
				yOffset = dis.OffsetY
				delay = dis.OffsetTime * float64(total-i)
			} else { // 去掉了特殊鱼
				x := int32(dis.OffsetX) * 100
				y := int32(dis.OffsetY) * 100
				dist := int32(dis.OffsetTime) * 100
				xOffset = float64(gameRand.Int31n(x*2)-x) / 100
				yOffset = float64(gameRand.Int31n(y*2)-y) / 100
				delay = float64(gameRand.Int31n(dist)) / 100
			}

			table.CreateFish(tmp, xOffset, yOffset, 0, delay, tmp.Speed, pathId, false, FishType_Normal, refreshId)
		}
		if fishType == SystemConf.SnakeHead {
			fishType = snakeType
		}
	}
}

func (table *Table) troopFish(ts *TroopConfig, time float64) {
	ref := &table.distributeTroop
	if ref.SendDes == false {
		ref.SendDes = true
		// AddBuffer(EBT_CHANGESPEED, 5, 60);
		if troop := GetTroop(ts.TroopId); troop != nil {
			table.GameUpdateAck.Describe = append(table.GameUpdateAck.Describe, troop.Describes...)
		}
	} else if ref.SendTroop == false && table.sceneTime > (ref.BeginTime+ts.BeginTime) {
		ref.SendTroop = true
		if troop := GetTroop(ts.TroopId); troop != nil {
			table.createTroopFish(troop)
		} else {
			table.sceneTime += time
		}
	}
}

func (table *Table) OnFirstFire(role *Role) {
	for i := 0; i < chairCount; i++ {
		if table.Roles[i] != role {
			continue
		}
		pt := Cannons[i].MovePoint
		createCount := SystemConf.FirstFire.CreatCount
		for count := int32(0); count < createCount; count++ {
			fishId := SystemConf.FirstFire.RandFish()
			dir := pt.Direction - (math.Pi / 2) + math.Pi/float64(createCount)*float64(count)
			tmp, ok := fishTemplateMap[fishId]
			if ok {
				delay := gameRand.Float64()
				table.CreateFish(tmp, pt.X, pt.Y, dir, delay, tmp.Speed, -2, false, FishType_Normal, 0)
			}
		}
	}
}

func (table *Table) OnProduceFish(source *Fish, args []int32) {
	a0 := args[0] // 参数0:表示要生成的鱼的类型
	a1 := args[1] // 参数1:表示要生成的鱼的批次
	a2 := args[2] // 参数2:表示每个批次要生成的鱼的数量
	a3 := args[3] // 参数3:表示每个批次之间的时间间隔
	refFish, ok := fishTemplateMap[a0]
	if ok == false {
		return
	}
	direction := math.Pi * 2 / float64(a2)

	king := int32(-1)
	for i := int32(0); i < a1; i++ {
		if i == a1-1 && a1 > 2 && a2 > 10 {
			// 最后一批鱼，随机出鱼王
			king = gameRand.Int31n(a2)
		}
		delay := 1.0 + float64(a3*i)
		for j := int32(0); j < a2; j++ {
			var FishType FishType
			if j == king {
				FishType = FishType_King
			} else {
				FishType = FishType_Normal
			}
			table.CreateFish(refFish, source.X, source.Y, direction*float64(j), delay, refFish.Speed, -2, false, FishType, 0)
		}
	}
}

func (table *Table) createTroopFish(troop *Troop) {
	if len(troop.Step) == 0 {
		return
	}
	n := 0
	refreshId := table.newFishId()
	for _, step := range troop.Step {
		fishId := int32(-1)
		for j := int32(0); j < step; j++ {
			if n >= len(troop.Shape) {
				break
			}
			tp := &troop.Shape[n]
			n++
			for i := int32(0); i < tp.Count; i++ {
				if fishId == -1 || (!tp.Same) {
					fishId = tp.RandFish()
				}
				if tmp, ok := fishTemplateMap[fishId]; ok {
					delay := float64(i) * tp.Interval
					table.CreateFish(tmp, tp.X, tp.Y, 0, delay, tp.Speed, tp.Path, true, FishType_Normal, refreshId)
				}
			}
		}
	}
}

func (table *Table) CreateFish(tmp *FishTemplate, x, y, direction, delay, speed float64, pathId int32, troop bool, fishType FishType, refreshId int32) {
	id := table.newFishId()
	f := popFish()
	f.Fish = fish.Fish{
		Id:        id,
		X:         x,
		Y:         y,
		Direction: direction,
		Troop:     troop,
		PathId:    pathId,
	}
	f.Temp = tmp
	f.Delay = delay
	f.Speed = speed
	f.RefreshId = refreshId
	f.Name =  tmp.Name
	f.FishType = fishType

	boxId := tmp.BoundBox
	if fishType != FishType_Normal {
		f.BroadCast = true
		var specialMap []SpecialFishTemplate
		switch fishType {
		case FishType_King:
			specialMap = kingFishes
			f.Name = tmp.Name + "鱼王"
		case FishType_KingQuan:
			specialMap = kingFishes
			f.Name = tmp.Name + "鱼王"
		case FishType_Sanyuan:
			specialMap = sanYuanFishes
			f.Name = "大三元"
		case FishType_Sixi:
			specialMap = siXiFishes
			f.Name = "大四喜"
		}
		for i := len(specialMap) - 1; i >= 0; i-- {
			special := &specialMap[i]
			if special.Id == tmp.Id {
				switch fishType {
				case FishType_King:
					f.Probability = special.CatchProbability
				case FishType_KingQuan:
					f.Probability = tmp.Probability / 5
					boxId = special.BoundBox
				case FishType_Sanyuan:
					f.Probability = tmp.Probability / 3
					boxId = special.BoundBox
				case FishType_Sixi:
					f.Probability = tmp.Probability / 4
					boxId = special.BoundBox
				}
				f.LockLevel = special.LockLevel
				break
			}
		}
	}

	bbx, ok := bbxMap[boxId]
	if ok == false {
		pushFish(f)
		return
	}
	f.BoxId = boxId
	f.InitBoxes = bbx.Boxes
	f.Boxes = make([]BoundingBox, len(bbx.Boxes))

	if pathId > 0 {
		f.Path = GetPathData(pathId, troop)
	} else {
		f.initAngle()
	}
	f.updateBBX()

	table.fishes[id] = f
	table.GameUpdateAck.Fishes = append(table.GameUpdateAck.Fishes, &f.Fish)
}

func (table *Table) Update() {
	tick := table.tickMs()
	elapsed := float64(tick - table.lastTick)
	if elapsed < 0 {
		elapsed = 1000.0 / float64(GAME_FPS)
	}
	second := elapsed / 1000.0
	// 更新鱼群
	table.updateFishes(second)
	// 更新子弹
	table.updateBullets(second)
	// 鱼群生成
	table.DistributeFish(second)

	// 机器人子弹生成
	upAck := &table.GameUpdateAck
	upAck.Tick = tick
	table.SendToAll(upAck)

	table.lastTick = tick
	upAck.SwitchScene = 0
	upAck.Players = upAck.Players[0:0]
	upAck.Fishes = upAck.Fishes[0:0]
	upAck.Bullets = upAck.Bullets[0:0]
	upAck.DieBullets = upAck.DieBullets[0:0]
	upAck.DieFishes = upAck.DieFishes[0:0]
	upAck.Kills = upAck.Kills[0:0]
	upAck.Seed = upAck.Seed[0:0]
	upAck.Offline = upAck.Offline[0:0]
}

// 发送消息给所有在线玩家
func (table *Table) SendToAll(val interface{}) {
	if val, err := room.Encode(val); err != nil {
		for _, role := range table.Roles {
			if role != nil {
				role.UnsafeSend(val)
			}
		}
	}
}

// 发送消息给其它玩家
func (table *Table) SendToOther(val interface{}, my *Role) {
	if val, err := room.Encode(val); err != nil {
		for _, role := range table.Roles {
			if role != nil && role != my {
				role.UnsafeSend(val)
			}
		}
	}
}
