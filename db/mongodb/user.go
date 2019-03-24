package mongodb

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"local.com/abc/game/model"
)

var (
	ErrorUserNotExists = errors.New("玩家不存在")
	ErrorNoUserID      = errors.New("分配ID失败")
	ErrorLockUserID    = errors.New("锁定ID失败")
	ErrorLockRoomUser  = errors.New("锁定到房间失败")
)

// 分配新的玩家ID和数据库时间
type userIdN struct {
	Id model.UserId `bson:"n"`
	Up time.Time    `bson:"up"`
}

func (d *driver) NewToken() int64 {
	return rand.Int63n(math.MaxInt64-int64(math.MaxInt32)) + int64(math.MaxInt32)
}

func (d *driver) GetRandName() string {
	return "G" + GetRandomString(7)
}

// 获取机器人
func (d *driver) GetRobot(room int32, count int32) []*model.User {
	return nil
}

// 机器人退出
func (d *driver) ExitRobot(room int32, users []model.UserId) {

}
