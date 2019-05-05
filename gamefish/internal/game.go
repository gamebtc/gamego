package internal

import (
	"local.com/abc/game/protocol/fish"
)

const (
	ObjectType_Player = 0 // 玩家
	ObjectType_Bullet = 1 // 子弹
	ObjectType_Fish   = 2 // 鱼
)

type ObjectState int32

const (
	ObjectState_Live     ObjectState = 0
	ObjectState_Hit      ObjectState = 1
	ObjectState_Dead     ObjectState = 2
	ObjectState_Destory  ObjectState = 3
	ObjectState_Lighting ObjectState = 4
)

const (
	EventType_ChangedState = 0 // 状态变化
	EventType_SpeedMulti   = 1 // 查询速度倍率
	EventType_AddMulti     = 2 // 查询额外增加的倍率
)

const (
	EffectType_AddMoney   = 0 // 增加金币
	EffectType_Kill       = 1 // 杀死其它鱼
	EffectType_AddBuffer  = 2 // 增加BUFFER
	EffectType_Produce    = 3 // 生成其它鱼
	EffectType_Blackwater = 4 // 乌贼喷墨汁效果
	EffectType_Award      = 5 // 抽奖
)

const (
	RefershType_Normal = 0 //
	RefershType_Group  = 1 // 鱼群
	RefershType_Line   = 2 // 鱼队
	RefershType_Snak   = 3 // 大蛇
)

type FishType int32
const(
	FishType_Normal FishType= 0 //
	FishType_King FishType= 1 //
	FishType_KingQuan FishType= 2 //
	FishType_Sanyuan FishType= 3 //
	FishType_Sixi FishType= 4 //
	FishType_Max FishType= 5 //
)

type Fish struct {
	fish.Fish
	FishType  FishType
	Name      string
	BroadCast bool
	Pause     bool
	EndPath   bool
	Offset    Point
	Delay     float64
	BeginMove bool
	Path      MovePoints
	Elapse    float64
	TargetId  int32
	LockLevel int32
	angle     float64
	dx        float64
	dy        float64

	Probability float64

	Temp      *FishTemplate
	RefreshId int32
	Score     int64
	InSide    bool
	State     ObjectState

	InitBoxes []BoundingBox
	Boxes     []BoundingBox
	table     *Table
}

func (fish *Fish) Update(second float64) {
	if fish.Pause || fish.EndPath {
		return
	}
	if fish.Delay > 0 {
		fish.Delay -= second
		return
	}
	if fish.Path != nil {
		fish.moveByPath(second)
	} else {
		fish.moveByDirection(second)
	}
	fish.updateBBX()
}

func (fish *Fish) updateBBX() {
	for i, ib := range fish.InitBoxes{
		bps := GetRotationPosByOffset(fish.X, fish.Y, ib.X, ib.Y, fish.Direction, 1, 1)
		fish.Boxes[i].X = bps.X
		fish.Boxes[i].Y = bps.Y
	}
	const diff = bulletRadius
	if fish.X > diff && fish.X < SystemConf.ScreenWidth-diff && fish.Y > diff && fish.Y < SystemConf.ScreenHeight-diff {
		fish.InSide = true
	} else {
		fish.InSide = false
	}
}
