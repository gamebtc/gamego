// +build mgo

package mongodb

import (
	"time"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"local.com/abc/game/util"
)

func(d *driver) GetIncrementKey(key string, n int) int64 {
	change := mgo.Change{
		Update:    bson.D{{"$inc", bson.D{{"n", n}}}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := incrementKeyDoc{}
	d.gameConf.FindId(key).Apply(change, &doc)
	return doc.N
}

func (d *driver) GetConfigValue(id string) int64 {
	doc := incrementKeyDoc{}
	d.gameConf.FindId(id).One(&doc)
	return doc.N
}

func (d *driver) SetConfigValue(id string, v interface{}) error {
	_, err := d.gameConf.Upsert(bson.D{{"_id", id}}, bson.D{{"$set", bson.D{{"n", v}}}})
	return err
}

//初始化纪元时间
func (d *driver)initEpoch() {
	change := mgo.Change{
		Update:    bson.D{{"$setOnInsert", bson.D{{"n", util.Now()}}}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := incrementKeyDoc{}
	d.conf.FindId("epoch").Apply(change, &doc)
	d.epoch = int32(doc.N)
}

func (d *driver)initUserId() {
	change := mgo.Change{
		Update:    bson.D{{"$setOnInsert", bson.D{{"n", 1000000}}}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := incrementKeyDoc{}
	d.conf.FindId(CollUser).Apply(change, &doc)
}

// 新的序列号
func (d *driver) NewSN(key interface{}, count int64) int64 {
	change := mgo.Change{
		Update:    bson.D{{"$inc", bson.D{{"n", count}}}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := incrementKeyDoc{}
	d.conf.Find(bson.D{{"_id", game}}).Apply(change,&doc)
	start := (doc.N - count) + 1
	if start > 0 {
		return start
	}
	return 0
}

var idNowUp = mgo.Change{Update: bson.D{{"$currentDate", bson.D{{"up", true}}}}, ReturnNew: true, Upsert: true}
var idNow = bson.D{{"_id","now"}}
func (d *driver) Now()time.Time {
	var t UpTime
	if _, e := d.conf.Find(idNow).Apply(idNowUp, &t); e == nil {
		return t.Up.Local()
	}
	return time.Now()
}