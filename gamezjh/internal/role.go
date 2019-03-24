package internal

import (
	//"errors"
	//"fmt"
	//log "github.com/sirupsen/logrus"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

// 每个角色的游戏数据
type Role struct {
	model.User           // 玩家信息
	*room.Session        // 发送消息
	table  *Table        // 桌子ID
	bill   *zjh.GameBill // 输赢情况
	player *zjh.Player   // 玩家信息
}

func (role *Role) GetPlayer() *zjh.Player {
	role.player.Coin = role.Coin
	return role.player
}

func (role *Role) IsRobot() bool {
	return role.User.Job == model.JobRobot
}

func (role *Role) IsPlayer() bool {
	return role.User.Job == model.JobPlayer
}

func (role *Role) Reset() {
	role.bill = &zjh.GameBill{
		Uid:  role.User.Id,
		Job:  role.User.Job,
		Coin: role.Coin,
	}

	role.player.State = zjh.Player_None
	role.player.Look = false
	role.player.Bet = 0
}

// 存在指定的下注金额
func ExistsBetItem(bet int32) bool {
	return (bet >= betItems[0]) && (bet <= betItems[len(betItems)-1]) && (bet%100 == 0)
}
