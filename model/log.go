package model

import (
	"time"

	"local.com/abc/game/protocol"
)

// 游戏币流水
type CoinFlow struct {
	Sn    int64       `json:"i" bson:"i" msg:"i"`                               // 交易流水号,全局交易中唯一，建有索引，不要更改i
	Uid   int32       `json:"u" bson:"u" msg:"u"`                               // 用户ID,建有索引，不要更改u
	Add   int64       `json:"a" bson:"a" msg:"a"`                               // 游戏币变化量
	Old   int64       `json:"o" bson:"o" msg:"o"`                               // 变化前的金币
	New   int64       `json:"w" bson:"w" msg:"w"`                               // 变化后的金币
	Tax   int64       `json:"x,omitempty" bson:"x,omitempty" msg:"x,omitempty"` // 税收
	Bet   int64       `json:"b,omitempty" bson:"b,omitempty" msg:"b,omitempty"` // 投注
	Kind  int32       `json:"k" bson:"-" msg:"k"`                               // 所在游戏
	Room  int32       `json:"r,omitempty" bson:"r,omitempty" msg:"r,omitempty"` // 所在房间ID
	Type  int32       `json:"t,omitempty" bson:"t,omitempty" msg:"t,omitempty"` // 原因
	State int32       `json:"s,omitempty" bson:"s,omitempty" msg:"s,omitempty"` // 状态
	Note  string      `json:"n,omitempty" bson:"n,omitempty" msg:"n,omitempty"` // 备注
	LogId int64       `json:"l,omitempty" bson:"l,omitempty" msg:"l,omitempty"` // 关联的游戏日志ID
	Poker []byte      `json:"p,omitempty" bson:"p,omitempty" msg:"p,omitempty"` // 开牌情况
	More  interface{} `json:"m,omitempty" bson:"m,omitempty" msg:"m,omitempty"` // 更多信息
}

type LoginLog struct {
	Id    ObjectId  `bson:"_id"` //唯一ID
	Agent int64     `bson:"agent"`
	Ip    IP        `bson:"ip"`
	Init  time.Time `bson:"init"`
	Up    time.Time `bson:"up"`
	State int32     `bson:"state"`
	Kind  int32     `bson:"kind"`
	Room  int32     `bson:"room"`
	Uid   int32     `bson:"uid"`
	Bag   CoinBag   `bson:"bag"`
	Udid  string    `bson:"udid"`

	Dev *protocol.DeviceInfo `bson:"dev"`
	Env *protocol.Envirnment `bson:"env"`
}