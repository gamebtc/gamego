// +build !mgo

package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

func (d *driver) LockRoomServer(room *msg.RoomConfig) (obj *model.RoomInfo, err error) {
	query := bson.D{
		{"_id", room.Id},
		{"kind", room.Kind},
		{"key", bson.D{{"$in", []string{"", room.Key}}}},
	}
	up := bson.D{
		{"$set", bson.D{{"key", room.Key}, {"addr", room.Addr}}},
		{"$inc", bson.D{{"ver", int32(1)}}},
		upNow,
	}
	obj = new(model.RoomInfo)
	err = d.roomCache.FindOneAndUpdate(d.ctx, query, up, returnNew).Decode(obj)
	return
}

func (d *driver) GetRoom(roomId int32, ver int32) (obj *model.RoomInfo, err error) {
	query := bson.D{{"_id", roomId}, {"ver", bson.D{{"$ne", ver}}}}
	obj = new(model.RoomInfo)
	err = d.roomCache.FindOne(d.ctx, query).Decode(obj)
	return
}

func (d *driver) GetAllRoom(query interface{}) (all []*model.RoomInfo, err error) {
	all = make([]*model.RoomInfo, 100)
	if cur, err := d.roomCache.Find(d.ctx, query); err == nil {
		defer cur.Close(d.ctx)
		for cur.Next(d.ctx) {
			a := new(model.RoomInfo)
			if err := cur.Decode(a); err == nil {
				all = append(all, a)
			}
		}
	}
	return
}

func (d *driver) SaveLog(collName string, value interface{}) (err error) {
	_, err = d.GetColl(collName).InsertOne(d.ctx, value)
	return
}
