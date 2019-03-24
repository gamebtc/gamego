package mongodb

import (
	//log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/abc/game/model"
	//"local.com/abc/game/protocol"
)

var(
	robotSort = options.FindOneAndUpdate().SetReturnDocument(options.After).SetSort(bson.D{{"ver", 1}})
	robotZeroRoom = bson.D{{"room", 0}}
)

// 加载机器人
func (d *driver) LoadRobot(room int32, count int32)[]*model.User {
	up := bson.D{
		{"$set", bson.D{{"room", room}}},
		{"$inc", bson.D{{"ver", 1}}},
		lastTime,
	}
	users := make([]*model.User, 0, count)
	for i := int32(0); i < count; i++ {
		user := new(model.User)
		err := d.robot.FindOneAndUpdate(d.ctx, robotZeroRoom, up, robotSort).Decode(user)
		if err != nil || user.Id == 0 {
			return users
		}
		user.Job = model.JobRobot
		users = append(users, user)
	}
	return users
}

// 卸载机器人
func (d *driver) UnloadRobot(room int32, ids[]model.UserId) {
	query := bson.D{
		{"_id", bson.D{{"$in", ids}}},
		{"room", room},
	}
	up := bson.D{
		{"$set", robotZeroRoom},
		lastTime,
	}
	d.robot.UpdateMany(d.ctx, query, up)
}

// 清理机器人
func (d* driver) ClearRobot(room int32) {
	query := bson.D{
		{"room", room},
	}
	d.robot.UpdateMany(d.ctx, query, bson.D{ {"$set", robotZeroRoom}})
}