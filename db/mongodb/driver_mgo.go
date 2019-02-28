// +build mgo

package mongodb

import (
	"context"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	_"github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/util"
)

func init(){
	bson.SetJSONTagFallback(true)
}

type Collection = mgo.Collection
type Database = mgo.Database
type Session = mgo.Session

func (d *driver)GetColl(name string) Collection {
	return *d.Database.C(name)
}

type Retval struct{
	Data   bson.Raw  `bson:"_r"`
	Code   int32     `bson:"_err"`
	Msg    string    `bson:"msg"`
}

func NewGameDriver(conf *msg.DatabaseConfig)(d *driver, err error) {
	defer util.PrintPanicStack()
	var s *mgo.Session
	if s, err = mgo.DialWithTimeout(conf.Url, 10*time.Second); err != nil {
		return
	}
	s.SetSyncTimeout(5 * time.Minute)
	s.SetSocketTimeout(5 * time.Minute)
	s.SetPoolLimit(256)
	s.SetMode(mgo.Eventual, true)
	d = &driver{ctx: context.Background()}
	d.Init(s.DB(conf.Name))
	return
}

func Dial(host string, name string)(db *mgo.Database, err error ){
	defer util.PrintPanicStack()
	var s *mgo.Session
	if s, err = mgo.DialWithTimeout(host, 10*time.Second);err != nil {
		return
		s.SetSyncTimeout(5 * time.Minute)
		s.SetSocketTimeout(5 * time.Minute)
		s.SetPoolLimit(256)
		s.SetMode(mgo.Eventual,true)
		db = s.DB(name)
	}
	return
}

func (d *driver) Exec(args []interface{}, result interface{})(*model.ExecError) {
	//cmd := "function(...args){return exec(...args)}"
	r := ExecResult{}
	cmd := bson.D{{"eval", evalFun}, {"args", args}, {"nolock", true}}
	if e := d.Database.Run(cmd, &r); e == nil {
		if r.Ok == 1 && r.Result.Data.Kind == bson.ElementDocument {
			bson.Unmarshal(r.Result.Data.Data, result)
			return nil
		} else {
			return &model.ExecError{Code: r.Result.Code, Msg: r.Result.Msg}
		}
	} else {
		return &model.ExecError{Code: 9999, Msg: e.Error()}
	}
}

func (d *driver) Exec2(args []interface{}, result interface{})(err error) {
	cmd := bson.D{{"eval", evalFun}, {"args", args}, {"nolock", true}}
	return d.Database.Run(cmd, result)
}

// 获取IP配置
func (d *driver) GetIpInfo(ip model.IP)(*model.IpInfo) {
	info := new(model.IpInfo)
	for i := 0; i < 4; i++ {
		if d.ip.FindId(ip).One(info) == nil {
			return info
		}
		now := time.Now()
		info.Init = now
		info.Up = now
		d.ip.Insert(info)
	}
	return nil
}

func (d *driver) GetPackConf(id int32)(*model.PackInfo){
 	return d.packCache.GetById(id)
}

// 获取渠道配置
func (d *driver) GetChanConf(code string)(*model.ChanInfo){
	return d.chanCache.GetByCode(code)
}

// 获取新的ID
var notExists = bson.D{{"_id", bson.D{{"$type", 6}}}}  //6=undefined
func (d *driver) NewId()(id model.ObjectId, e error) {
	var r *mgo.ChangeInfo
	if r, e = d.id.Upsert(notExists, nil); e == nil {
		id, _ = r.UpsertedId.(model.ObjectId)
	}
	return
}


// 创建非唯一索引
func ensureIndex(c *mgo.Collection,key []string) error {
	return c.EnsureIndex(mgo.Index{
		Key:    key,
		Unique: false,
		Sparse: true,
	})
}

// 创建唯一索引
func ensureUniqueIndex(c *mgo.Collection, key []string) error {
	return c.EnsureIndex(mgo.Index{
		Key:    key,
		Unique: true,
		Sparse: true,
	})
}
