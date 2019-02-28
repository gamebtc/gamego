// +build mgo

package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	_ "github.com/sirupsen/logrus"
	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

var (
	upNow       = bson.DocElem{Name: "$currentDate", Value: bson.D{{"up", true}}}
	initNow     = bson.DocElem{Name: "$currentDate", Value: bson.D{{"init", true}, {"up", true}}}
	upNowChange = mgo.Change{Update: bson.D{{"$currentDate", bson.D{{"up", true}}}}, ReturnNew: true}
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
	err = d.account.Find(query).Select(accountSelect).One(acc)
	return
}

// 创建账号
func (d *driver) CreateAccount(acc *model.Account, req *msg.LoginReq) (err error) {
	err = d.account.Insert(acc)
	if err == nil {
		set := bson.DocElem{Name: "$set", Value: bson.D{{"env", req.Env}, {"dev", req.Dev}}}
		// 更新为数据库时间
		d.account.UpdateId(acc.Id, bson.D{initNow, set})
	}
	return
}

func (d *driver) GetUser(id int32) (user *model.User, err error) {
	user = new(model.User)
	err = d.user.FindId(id).One(user)
	return
}

func (d *driver) LoadUser(user *model.User) error {
	change := mgo.Change{
		Update: bson.D{
			{"$set", bson.D{{"lastIp", user.LastIp}}},
			{"$currentDate", bson.D{{"up", true}, {"lastTime", true}}},
		},
		ReturnNew: true,
	}
	if changed, e := d.user.FindId(user.Id).Apply(change, user); e != nil {
		return e
	} else {
		if changed.Matched == 1 {
			return nil
		}
	}
	return ErrorUserNotExists
}

// 分配新的玩家ID和数据库时间
var add100TChange = mgo.Change{
	Update:    bson.D{{"$inc", bson.D{{"t", int32(100)}}}, upNow},
	ReturnNew: true,
}
var zeroT = bson.D{{Name: "t", Value: zeroInt32}}

func (d *driver) newUserId() *userIdN {
	i := new(userIdN)
	d.C(CollUserId).Find(zeroT).Apply(add100TChange, i)
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
	if err = d.user.Insert(user); err == nil {
		// 更新账号的Users
		d.account.UpsertId(user.Act, bson.D{{"$push", bson.D{{"users", id}}}, upNow})
	}
	return
}

// 锁定玩家登录
var zeroRoom = bson.D{{"kind", zeroInt32}, {"room", zeroInt32}, {"table", zeroInt32}}
var maxRoom = bson.DocElem{Name: "$max", Value: zeroRoom}

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
	change := mgo.Change{
		Update: bson.D{{"$set", newLock}, maxRoom},
		Upsert: true,
	}
	lock := new(model.UserLocker)
	if _, err := d.locker.FindId(user.Id).Apply(change, lock); err != nil {
		return nil, err
	}
	if lock.Log1.Valid() {
		//旧的连接设置为强制退出
		set := bson.D{
			{"up", t},
			{"state", 2},
			{"kind", lock.Kind},
			{"room", lock.Room},
		}
		d.loginLog.UpdateId(lock.Log1, bson.D{{"$set", set}})
	}
	// 写登录日志
	newLock = append(newLock,
		bson.DocElem{"state", zeroInt32},
		bson.DocElem{"kind", lock.Kind},
		bson.DocElem{"room", lock.Room},
		bson.DocElem{"user", user.Id},
		bson.DocElem{"bag", user.Bag},
		bson.DocElem{"udid", req.Udid},
		bson.DocElem{"env", req.Env},
		bson.DocElem{"dev", req.Dev},
	)
	newLock[0].Name = "_id"
	d.loginLog.Insert(newLock)

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
	change := mgo.Change{
		Update:    bson.D{{"$set", bson.D{{"agent", zeroInt64}}}, upNow},
		ReturnNew: true,
	}
	newLock := new(model.UserLocker)
	if changed, err := d.locker.Find(query).Apply(change, newLock); err == nil {
		if changed.Updated == 1 {
			// 更新对应的日志记录
			d.loginLog.UpdateId(newLock.Log1, bson.D{{"$set", bson.D{{"state", int32(1)}}}, upNow})
			// 删除玩家在线记录
			d.locker.Remove(bson.D{{"_id", userId}, {"agent", zeroInt64}, {"room", zeroInt32}})
			return true
		}
	}
	return false
}

// 锁定用户到指定房间
func (d *driver) LockUserRoom(agent int64, userId int32, kind int32, roomId int32) (*model.User, error) {
	query := bson.D{
		{"_id", userId},
		{"agent", agent},
		{"room", bson.D{{"$in", []int32{0, roomId}}}},
	}
	newId := model.NewObjectId()
	change := mgo.Change{
		Update:    bson.D{{"$set", bson.D{{"log2", newId}, {"kind", kind}, {"room", roomId}}}, upNow},
		ReturnNew: false,
	}
	oldLock := new(model.UserLocker)
	if changed, err := d.locker.Find(query).Apply(change, oldLock); err != nil {
		return nil, err
	} else {
		if changed.Matched == 0 {
			// 登录已过期
			if err = d.locker.Find(query[0:1]).One(oldLock); err != nil {
				return nil, err
			}
			return nil, err
		}
	}
	if oldLock.Log2.Valid() {
		//旧的连接设置为强制退出
		d.roomLog.UpdateId(oldLock.Log2, bson.D{{"$set", bson.D{{"state", 2}}}, upNow})
	}

	user := new(model.User)
	if err := d.user.Find(query[0:1]).One(user); err != nil {
		return nil, err
	}
	//写登录房间日志
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
	d.roomLog.Insert(newLock)

	return user, nil
}

// 解锁用户从指定房间
func (d *driver) UnlockUserRoom(agent int64, userId int32, roomId int32) bool {
	query := bson.D{
		{"_id", userId},
		{"agent", agent},
		{"room", roomId},
	}
	change := mgo.Change{
		Update:    bson.D{{"$set", bson.D{{"kind", zeroInt32}, {"room", zeroInt32}}}, upNow},
		ReturnNew: false,
	}
	lock := new(model.UserLocker)
	if changed, err := d.locker.FindId(query).Apply(change, lock); err == nil {
		if changed.Matched == 1 {
			// 更新对应的日志记录
			d.roomLog.UpdateId(lock.Log2, bson.D{{"$set", bson.D{{"state", int32(1)}}}, upNow})
			// 删除玩家在线记录
			d.locker.Remove(bson.D{{"_id", userId}, {"agent", zeroInt64}, {"room", zeroInt32}})
			return true
		}
	}
	return false
}

//根据帐号查找ID
func (d *driver) FindUserIdByAccount(account string) int32 {
	var res struct {
		Id int32 `bson:"_id"`
	}
	d.user.Find(bson.D{{"user", account}}).One(&res)
	return res.Id
}

// 设置内容
func (d *driver) SetValue(id, v interface{}) error {
	_, err := d.user.Upsert(bson.D{{"_id", id}}, bson.M{"$set": v})
	return err
}

func (d *driver) UserLogin(req *msg.LoginReq) (*msg.LoginSuccessAck, error) {
	return nil, nil
}
