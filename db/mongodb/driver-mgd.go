package mongodb

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

type Collection = *mongo.Collection
type Database = mongo.Database
type Session = mongo.Session

func (d *driver)GetColl(name string) Collection {
	return d.Database.Collection(name)
}

type Retval struct {
	Data bson.Raw `bson:"_r"`
	Code int32    `bson:"_err"`
	Msg  string   `bson:"msg"`
}

func NewGameDriver(conf *protocol.DatabaseConfig)(d *driver, err error) {
	defer util.PrintPanicStack()
	var client *mongo.Client
	if client, err = mongo.NewClient(options.Client().ApplyURI(conf.Url)); err != nil {
		log.Fatal(err)
		return
	}

	ctx := context.Background()
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
		return
	}

	d = &driver{ctx: ctx}
	d.Init(client.Database(conf.Name))
	return
}

func (d *driver) Exec(args []interface{}, result interface{})*model.ExecError {
	r := ExecResult{}
	if e := d.Exec2(args, r); e != nil {
		return &model.ExecError{Code: 9999, Msg: e.Error()}
	}
	if r.Ok != 1 {
		return &model.ExecError{Code: r.Result.Code, Msg: r.Result.Msg}
	}
	bson.Unmarshal(([]byte)(r.Result.Data), result)
	return nil
}

func (d *driver) Exec2(args []interface{}, result interface{})error {
	cmd := bson.D{
		{"eval", evalFun},
		{"args", args},
		{"nolock", true},
	}
	reader := d.Database.RunCommand(d.ctx, cmd)
	if e := reader.Err(); e != nil {
		return e
	}
	raw, e := reader.DecodeBytes()
	if e != nil {
		return e
	}
	return bson.Unmarshal(([]byte)(raw), result)
}

//获取IP配置
func (d *driver) GetIpInfo(ip model.IP)*model.IpInfo {
	info := new(model.IpInfo)
	filter := bson.D{{"_id", ip}}
	for i := 0; i < 4; i++ {
		err := d.ip.FindOne(d.ctx, filter).Decode(info)
		if err == nil {
			return info
		}
		if err == mongo.ErrNoDocuments {
			now := time.Now()
			info.Id = ip
			info.Born = now
			info.Up = now
			_, err = d.ip.InsertOne(d.ctx, info)
			if err == nil {
				return info
			}
		} else {
			return nil
		}
	}
	return nil
}

func (d *driver) GetPackConf(id int32)*model.PackInfo{
	return d.packCache.GetById(id)
}

// 获取渠道配置
func (d *driver) GetChanConf(code string)*model.ChanInfo{
	return d.chanCache.GetByCode(code)
}

// 获取新的ID
var emptyMap = bson.D{}
func (d *driver) NewId()(id model.ObjectId, err error) {
	var r *mongo.InsertOneResult
	if r, err = d.id.InsertOne(d.ctx, emptyMap); err == nil {
		id, _ = r.InsertedID.(model.ObjectId)
	}
	return
}

// 创建非唯一索引
func ensureIndex(c *mongo.Collection,key []string) error {
	//return c.EnsureIndex(mongo.Index{
	//	Key:    key,
	//	Unique: false,
	//	Sparse: true,
	//})
	return nil
}

// 创建唯一索引
func ensureUniqueIndex(c *mongo.Collection, key []string) error {
	//return c.EnsureIndex(mongo.Index{
	//	Key:    key,
	//	Unique: true,
	//	Sparse: true,
	//})
	return nil
}

