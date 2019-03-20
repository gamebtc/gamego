package internal

import (
	//"errors"
	//"fmt"
	//log "github.com/sirupsen/logrus"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)


// 每个玩家的游戏数据
type Role struct {
	*model.User   // 玩家信息
	*room.Session // 发送消息
	*zjh.Player
	Online bool          // 是否在线
	Ready  bool          // 是否准备好
	Coin   int64         // 当前房间使用的币
	table  *Table        // 桌子ID
	chair  int32         // 椅子ID
	bill   *zjh.GameBill // 输赢情况
	flowSn int64         // 最后的写分序号，返回时用于验证
}

func (role *Role)GetPlayer() *zjh.Player {
	return role.Player
	//return &zjh.Player{
	//	Id:    role.Id,
	//	Icon:  role.Icon,
	//	Vip:   role.Vip,
	//	Name:  role.Name,
	//	Coin:  role.Coin,
	//}
}

func(role *Role)IsRobot() bool {
	return role.User.Job == model.JobRobot
}

func(role *Role)IsPlayer() bool {
	return role.User.Job == model.JobPlayer
}

func (role *Role) Reset() {
	role.bill = &zjh.GameBill{
		Uid:  role.User.Id,
		Coin: role.Coin,
		Job:  role.Job,
	}
}

// 存在指定的下注金额
func ExistsBetItem(bet int32)bool {
	return (bet >= betItems[0]) && (bet <= betItems[len(betItems)-1]) && (bet%100 == 0)
}
