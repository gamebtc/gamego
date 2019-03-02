// +build !mgo

package mongodb

import (
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

var (
	upNow       = bson.E{Key: "$currentDate", Value: bson.D{{"up", true}}}
	initNow     = bson.E{Key: "$currentDate", Value: bson.D{{"init", true}, {"up", true}}}
	upNowChange = bson.D{{"$currentDate", bson.D{{"up", true}}}}

	retnew = options.FindOneAndUpdate().SetReturnDocument(options.After)
	upsert = options.FindOneAndUpdate().SetUpsert(true)
)

// 获取账号
var accountSelect = bson.D{{"env", zeroInt32}, {"dev", zeroInt32}}

func (d *driver) GetAccount(app int32, t int32, name string) (acc *model.Account, err error) {
	query := bson.D{
		{"app", app},
		{"type", t},
		{"name", name},
	}
	acc = new(model.Account)
	err = d.account.FindOne(d.ctx, query).Decode(acc)
	return
}

// 创建账号
func (d *driver) CreateAccount(acc *model.Account, req *msg.LoginReq) (err error) {
	if _, err = d.account.InsertOne(d.ctx, acc); err == nil {
		// 更新为数据库时间
		set := bson.D{initNow, {"$set", bson.D{{"env", req.Env}, {"dev", req.Dev}}}}
		d.account.UpdateOne(d.ctx, bson.D{{"_id", acc.Id}}, set)
	}
	log.Debugf("CreateAccount:%v,err%v", acc, err)
	return
}

func (d *driver) GetUser(id int32) (user *model.User, err error) {
	user = new(model.User)
	err = d.user.FindOne(d.ctx, bson.D{{"_id", id}}).Decode(user)
	return
}

var lastTime = bson.E{Key: "$currentDate", Value: bson.D{{"up", true}, {"last", true}}}

func (d *driver) LoadUser(user *model.User) error {
	up := bson.D{{"$set", bson.D{{"lastIp", user.LastIp}}}, lastTime}
	query := bson.D{{"_id", user.Id}}
	return d.user.FindOneAndUpdate(d.ctx, query, up, retnew).Decode(user)
}

var add100TChange = bson.D{{"$inc", bson.D{{"t", int32(100)}}}, upNow}
var zeroT = bson.D{{"t", zeroInt32}}

func (d *driver) newUserId() *userIdN {
	i := new(userIdN)
	d.userId.FindOneAndUpdate(d.ctx, zeroT, add100TChange, retnew).Decode(i)
	return i
}

// 创建账号
func (d *driver) CreateUser(user *model.User, req *msg.LoginReq) (err error) {
	i := d.newUserId()
	id := i.N
	if id == 0 {
		err = ErrorNoUserID
		return
	}
	// 加入玩家表
	user.Id = id
	user.Init = i.T
	user.Up = i.T
	user.Last = i.T
	if _, err = d.user.InsertOne(d.ctx, user); err == nil {
		// 更新账号的Users
		query := bson.D{{"_id", user.Act}}
		up := bson.D{{"$push", bson.D{{"users", id}}}, upNow}
		_, err = d.account.UpdateOne(d.ctx, query, up)
	}
	fmt.Printf("CreateUser:%#v", err)
	return
}

// 锁定玩家登录
var zeroRoom = bson.D{{"kind", zeroInt32}, {"room", zeroInt32}, {"table", zeroInt32}}
var maxRoom = bson.E{Key: "$max", Value: zeroRoom}

func (d *driver) LockUser(agent int64, user *model.User, req *msg.LoginReq) (*model.UserLocker, error) {
	t := user.Last
	newId := model.NewObjectId()
	newLock := bson.D{
		{"log1", newId},
		{"agent", agent},
		{"ip", user.LastIp},
		{"init", t},
		{"up", t},
	}

	up := bson.D{{"$set", newLock}, maxRoom}

	query := bson.D{{"_id", user.Id}}
	lock := new(model.UserLocker)
	replace := false
	if err := d.locker.FindOneAndUpdate(d.ctx, query, up, upsert).Decode(lock); err != nil {
		if err != model.ErrNoDocuments {
			return nil, err
		}
	} else {
		replace = !lock.Log1.IsZero()
		if replace {
			// 旧的连接设置为强制退出
			set := bson.D{{"$set", bson.D{
				{"up", t},
				{"state", 2},
				{"kind", lock.Kind},
				{"room", lock.Room},
				{"room", lock.Room},
			}}}
			query[0].Value = lock.Log1 //bson.D{{"_id", lock.Log1}
			d.loginLog.UpdateOne(d.ctx, query, set)
		}
	}

	//写登录日志
	newLock[0].Key = "_id"
	newLock = append(newLock,
		bson.E{Key: "state", Value: zeroInt32},
		bson.E{Key: "kind", Value: lock.Kind},
		bson.E{Key: "room", Value: lock.Room},
		bson.E{Key: "user", Value: user.Id},
		bson.E{Key: "bag", Value: user.Bag},
		bson.E{Key: "udid", Value: req.Udid},
		bson.E{Key: "env", Value: req.Env},
		bson.E{Key: "dev", Value: req.Dev})

	if replace{
		newLock = append(newLock, bson.E{Key: "f", Value:lock.Log1})
	}

	d.loginLog.InsertOne(d.ctx, newLock)

	lock.Ip = user.LastIp
	lock.Log1 = newId
	lock.Up = t
	lock.Init = t

	return lock, nil
}

// 解锁玩家登录
func (d *driver) UnlockUser(agent int64, userId int32) bool {
	query := bson.D{
		{"_id", userId},
		{"agent", agent},
	}
	// 先将当前连接ID更新为0
	up := bson.D{{"$set", bson.D{{"agent", zeroInt64}}}, upNow}
	newLock := new(model.UserLocker)
	if err := d.locker.FindOneAndUpdate(d.ctx, query, up, retnew).Decode(newLock); err == nil {
		// 更新对应的日志记录
		up = bson.D{{"$set", bson.D{{"state", int32(1)}}}, upNow}
		d.loginLog.UpdateOne(d.ctx, bson.D{{"_id", newLock.Log1}}, up)
		// 删除玩家在线记录
		d.locker.DeleteOne(d.ctx, bson.D{{"_id", userId}, {"agent", zeroInt64}, {"room", zeroInt32}})
		return true
	}
	return false
}

// 锁定用户到指定房间
func (d *driver) LockUserRoom(agent int64, userId int32, kind int32, roomId int32) (*model.User, error) {
	query := bson.D{
		{"_id", userId},
		{"agent", agent},
		{"room", bson.D{{"$in", []int32{zeroInt32, roomId}}}},
	}
	newId := model.NewObjectId()
	up := bson.D{{"$set", bson.D{{"log2", newId}, {"kind", kind}, {"room", roomId}}}, upNow}

	oldLock := new(model.UserLocker)
	if changed, err := d.locker.UpdateOne(d.ctx, query, up); err != nil {
		return nil, err
	} else {
		if changed.MatchedCount == 0 {
			// 登录已过期
			if err = d.locker.FindOne(d.ctx, query[:1]).Decode(oldLock); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	replace := !oldLock.Log2.IsZero()
	if replace{
		// 旧的连接设置为强制退出
		d.roomLog.UpdateOne(d.ctx, bson.D{{"_id", oldLock.Log2}}, bson.D{{"$set", bson.D{{"state", int32(2)}}}, upNow})
	}

	user := new(model.User)
	if err := d.user.FindOne(d.ctx, query[:1]).Decode(user); err != nil {
		return nil, err
	}
	// 写登录房间日志
	newLock := bson.D{
		{"_id", newId},
		{"win", zeroInt32},
		{"state", zeroInt32},
		{"kind", kind},
		{"room", roomId},
		{"user", user.Id},
		{"bag", user.Bag},
		{"ip", user.LastIp},
	}
	if replace {
		newLock = append(newLock, bson.E{Key: "f", Value:oldLock.Log2})
	}
	d.roomLog.InsertOne(d.ctx, newLock)

	return user, nil
}

// 解锁用户从指定房间
func (d *driver) UnlockUserRoom(agent int64, userId int32, roomId int32) bool {
	query := bson.D{
		{"_id", userId},
		{"agent", agent},
		{"room", roomId},
	}
	up := bson.D{{"$set", bson.D{{"kind", zeroInt32}, {"room", zeroInt32}}}, upNow}
	lock := new(model.UserLocker)
	if changed, err := d.locker.UpdateOne(d.ctx, query, up); err == nil {
		if changed.MatchedCount == 1 {
			// 更新对应的日志记录
			up1 := bson.D{{"$set", bson.D{{"state", int32(1)}}}, upNow}
			d.roomLog.UpdateOne(d.ctx, bson.D{{"_id", lock.Log2}}, up1)
			// 删除玩家在线记录
			query[1].Value = 0
			query[2].Value = 0
			d.locker.DeleteOne(d.ctx, query)
			return true
		}
	}
	return false
}

// 根据帐号查找ID
func (d *driver) FindUserIdByAccount(account string) int32 {
	var res struct {
		Id int32 `bson:"_id"`
	}
	d.user.FindOne(d.ctx, bson.D{{"user", account}}).Decode(&res)
	return res.Id
}

// 设置内容
func (d *driver) SetValue(id, v interface{}) error {
	_, err := d.user.UpdateOne(d.ctx, bson.D{{"_id", id}}, bson.D{{"$set", v}}, upsert2)
	return err
}

func (d *driver) UserLogin(req *msg.LoginReq) (*msg.LoginSuccessAck, error) {
	return nil, nil
}

// 获取agentId
func (d *driver) CheckUserAgent(userId int32, agent int64) bool {
	query := bson.D{
		{"_id", userId},
		{"agent", agent},
	}
	if changed, err := d.locker.UpdateOne(d.ctx, query, bson.D{upNow}); err == nil {
		return changed.MatchedCount > 0
	}
	return false
}
