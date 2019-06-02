package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
)

func (d *driver) LockRoomServer(room *protocol.RoomConfig) (obj *model.RoomInfo, err error) {
	query := bson.D{
		{"_id", room.Id},
		{"game", room.Game},
		{"state", int32(1)},
		{"key", bson.D{{"$in", []string{"", room.Key}}}},
	}
	up := bson.D{
		{"$set", bson.D{{"key", room.Key}, {"addr", room.Addr}, {"pause", zero32}, {"close", zero32}, {"lock", zero32}}},
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

func (d *driver) FindRoomInfos(query interface{}) (all []*model.RoomInfo, err error) {
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

func(d *driver) FindRoomConfigs(query interface{}) (all []*protocol.RoomConfig, err error){
	op:=options.Find().SetProjection(bson.D{{"game", true},{"addr", true}})
	cur, err := d.roomCache.Find(d.ctx, query, op)
	if err != nil {
		return nil, err
	}
	defer cur.Close(d.ctx)
	all = make([]*protocol.RoomConfig, 100)
	for cur.Next(d.ctx) {
		a := new(protocol.RoomConfig)
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
