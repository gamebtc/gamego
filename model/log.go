package model

import (
	"time"

	"local.com/abc/game/msg"
)

type CoinFlow = msg.CoinFlow

type LoginLog struct {
	Id    ObjectId        `bson:"_id"` //唯一ID
	Agent int64           `bson:"agent"`
	Ip    IP              `bson:"ip"`
	Init  time.Time       `bson:"init"`
	Up    time.Time       `bson:"up"`
	State int32           `bson:"state"`
	Kind  int32           `bson:"kind"`
	Room  int32           `bson:"room"`
	User  int32           `bson:"user"`
	Coin  CoinBag         `bson:"coin"`
	Udid  string          `bson:"udid"`
	Dev   *msg.DeviceInfo `bson:"dev"`
	Env   *msg.Envirnment `bson:"env"`
}