// +build mgo

package mongodb

import (
	"github.com/globalsign/mgo/bson"

	"local.com/abc/game/model"
)

func (d *driver) GetRoom(roomId int32, ver int32) (obj *model.RoomInfo, err error) {
	query := bson.D{{"_id", roomId}, {"ver", bson.D{{"$ne", ver}}}}
	obj = new(model.RoomInfo)
	err = d.roomCache.Find(query).One(obj)
	return
}

func (d *driver) GetAllRoom(query interface{}) (obj []*model.RoomInfo, err error) {
	obj = make([]*model.RoomInfo,100)
	err = d.roomCache.Find(query).All(obj)
	return
}
