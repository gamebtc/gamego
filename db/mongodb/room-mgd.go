package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
)

func (d *driver) LockRoomServer(room *protocol.RoomConfig) (obj *model.RoomInfo, err error) {
	query := bson.D{
		{"_id", room.Id},
		{"kind", room.Kind},
		{"key", bson.D{{"$in", []string{"", room.Key}}}},
	}
	up := bson.D{
		{"$set", bson.D{{"key", room.Key}, {"addr", room.Addr}, {"pause", int32(0)}, {"close", int32(0)}, {"lock", int32(0)}}},
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
	cur, err := d.roomCache.Find(d.ctx, query)
	if err != nil {
		return nil, err
	}
	defer cur.Close(d.ctx)
	all = make([]*model.RoomInfo, 100)
	for cur.Next(d.ctx) {
		a := new(model.RoomInfo)
		if e := cur.Decode(a); e == nil {
			all = append(all, a)
		}
	}
	return
}

func (d *driver) SaveLog(collName string, value interface{}) error {
	_, err := d.GetColl(collName).InsertOne(d.ctx, value)
	return err
}
