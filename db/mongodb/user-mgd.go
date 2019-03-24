package mongodb

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
)

var (
	upNow       = bson.E{Key: "$currentDate", Value: bson.D{{"up", true}}}
	initNow     = bson.E{Key: "$currentDate", Value: bson.D{{"init", true}, {"up", true}}}
	upNowChange = bson.D{{"$currentDate", bson.D{{"up", true}}}}

	retnew = options.FindOneAndUpdate().SetReturnDocument(options.After)
	upsert = options.FindOneAndUpdate().SetUpsert(true)
)

// 创建账号
func (d *driver) CreateAccount(acc *model.Account, req *protocol.LoginReq) (err error) {
	if _, err = d.account.InsertOne(d.ctx, acc); err == nil {
		// 更新为数据库时间
		set := bson.D{initNow, {"$set", bson.D{{"env", req.Env}, {"dev", req.Dev}}}}
		d.account.UpdateOne(d.ctx, bson.D{{"_id", acc.Id}}, set)
	}
	log.Debugf("CreateAccount:%v,err%v", acc, err)
	return
}

// 获取账号
var accountSelect = options.FindOne().SetProjection(bson.D{{"env", false}, {"dev", false}})

func (d *driver) GetAccount(app int32, t int32, name string) (acc *model.Account, err error) {
	query := bson.D{
		{"app", app},
		{"type", t},
		{"name", name},
	}
	acc = new(model.Account)
	err = d.account.FindOne(d.ctx, query, accountSelect).Decode(acc)
	return
}

var zeroT = bson.D{{"t", zero32}}
var add100TChange = bson.D{{"$inc", bson.D{{"t", int32(100)}}}, upNow}

// 创建账号
func (d *driver) CreateUser(user *model.User, req *protocol.LoginReq) (e error) {
	// 获取账号ID和时间
	i := userIdN{}
	d.userId.FindOneAndUpdate(d.ctx, zeroT, add100TChange, retnew).Decode(&i)

	id := i.Id
	if id == 0 {
		e = ErrorNoUserID
		return
	}
	// 加入玩家表
	user.Id = id
	user.Init = i.Up
	user.Up = i.Up
	user.Last = i.Up
	f := func(sc mongo.SessionContext) (err error) {
		if _, err = d.user.InsertOne(d.ctx, user); err == nil {
			// 创建背包
			query := bson.D{{"_id", id}}
			up := bson.D{{"$set", user.Bag}, upNow}
			if _, err = d.bag.UpdateOne(d.ctx, query, up, upsert2); err == nil {
				// 更新账号的Users
				query = bson.D{{"_id", user.Act}}
				up = bson.D{{"$push", bson.D{{"users", id}}}, upNow}
				if _, err = d.account.UpdateOne(d.ctx, query, up); err == nil {
					return sc.CommitTransaction(sc)
				}
			}
		}
		sc.AbortTransaction(sc)
		return
	}
	e = d.Client().UseSessionWithOptions(d.ctx, sessionOp, f)
	log.Debugf("CreateUser:%v,err%v", user, e)
	return
}

var lastTime = bson.E{Key: "$currentDate", Value: bson.D{{"up", true}, {"last", true}}}
var bagOp = options.FindOne().SetProjection(bson.D{{"_id", false}, {"up", false}})

func (d *driver) LoadUser(uid model.UserId, ip model.IP) (user *model.User, err error) {
	up := bson.D{{"$set", bson.D{{"ip", ip}}}, lastTime}
	query := bson.D{{"_id", uid}}
	user = &model.User{}
	if err = d.user.FindOneAndUpdate(d.ctx, query, up, retnew).Decode(user); err == nil {
		bag := model.CoinBag{}
		d.bag.FindOne(d.ctx, query, bagOp).Decode(&bag)
		user.Bag = bag
	}
	return
}

// 锁定玩家登录
var zeroRoom = bson.D{{"kind", zero32}, {"room", zero32}, {"table", zero32}}
var maxRoom = bson.E{Key: "$max", Value: zeroRoom}

func (d *driver) LockUser(agent int64, uid model.UserId, ip model.IP, t time.Time, req *protocol.LoginReq) (*model.UserLocker, error) {
	newId := model.NewObjectId()
	newLock := bson.D{
		{"log1", newId},
		{"agent", agent},
		{"ip", ip},
		{"init", t},
		{"up", t},
	}

	up := bson.D{{"$set", newLock}, maxRoom}

	query := bson.D{{"_id", uid}}
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
			}}}
			query[0].Value = lock.Log1 //bson.D{{"_id", lock.Log1}
			d.loginLog.UpdateOne(d.ctx, query, set)
		}
	}

	//写登录日志
	newLock[0].Key = "_id"
	newLock = append(newLock,
		bson.E{Key: "state", Value: zero32},
		bson.E{Key: "kind", Value: lock.Kind},
		bson.E{Key: "room", Value: lock.Room},
		bson.E{Key: "uid", Value: uid},
		bson.E{Key: "udid", Value: req.Udid},
		bson.E{Key: "env", Value: req.Env},
		bson.E{Key: "dev", Value: req.Dev})

	if replace {
		newLock = append(newLock, bson.E{Key: "f", Value: lock.Log1})
	}

	d.loginLog.InsertOne(d.ctx, newLock)

	lock.Ip = ip
	lock.Log1 = newId
	lock.Up = t
	lock.Init = t

	return lock, nil
}

// 解锁玩家登录
func (d *driver) UnlockUser(agent int64, uid model.UserId) bool {
	query := bson.D{
		{"_id", uid},
		{"agent", agent},
	}
	// 先将当前连接ID更新为0
	up := bson.D{{"$set", bson.D{{"agent", zero64}}}, upNow}
	lock := new(model.UserLocker)
	if err := d.locker.FindOneAndUpdate(d.ctx, query, up, retnew).Decode(lock); err == nil {
		// 更新对应的日志记录
		if !lock.Log1.IsZero() {
			up = bson.D{{"$set", bson.D{{"state", int32(1)}}}, upNow}
			d.loginLog.UpdateOne(d.ctx, bson.D{{"_id", lock.Log1}}, up)
			// 删除玩家锁
			d.locker.DeleteOne(d.ctx, bson.D{{"_id", uid}, {"agent", zero64}, {"room", zero32}})
		}
		return true
	}
	return false
}

// 锁定用户到指定房间
func (d *driver) LockUserRoom(agent int64, uid model.UserId, kind int32, roomId int32, coinKey string, win int64, bet int64, round int32) (*model.User, error) {
	query := bson.D{
		{"_id", uid},
		{"agent", agent},
		{"room", bson.D{{"$in", []int32{zero32, roomId}}}},
	}
	newId := model.NewObjectId()
	up := bson.D{{"$set", bson.D{{"log2", newId}, {"kind", kind}, {"room", roomId}}}, upNow}

	lock := new(model.UserLocker)
	if err := d.locker.FindOneAndUpdate(d.ctx, query, up, returnNew).Decode(lock); err != nil {
		if err == model.ErrNoDocuments {
			if err = d.locker.FindOne(d.ctx, query[:1]).Decode(lock); err != nil {
				if err != model.ErrNoDocuments {
					return nil, err
				}
			}

			if lock.Id != uid || lock.Agent != agent {
				// 登录已过期
				return nil, protocol.ErrorLoginExpired
			} else if lock.Room != roomId {
				// 已登录其它房间
				return nil, errors.New(fmt.Sprintf("您已登录游戏[%v]房间[%v]", lock.Kind, lock.Room))
			}
			return nil, protocol.ErrorLoginExpired
		}
		return nil, err
	}

	replace := !lock.Log2.IsZero()
	if replace {
		// 旧的连接设置为强制退出
		up1 := bson.D{{"$set", bson.D{{"state", int32(2)}, {"win", win}, {"bet", bet}, {"round", round}}}, upNow}
		d.roomLog.UpdateOne(d.ctx, bson.D{{"_id", lock.Log2}, {"state", int32(0)}}, up1)
	}

	query = query[:1]
	user := new(model.User)
	if err := d.user.FindOne(d.ctx, query).Decode(user); err != nil {
		return nil, err
	}

	bagDealOp := options.FindOne().SetProjection(bson.D{{"_id", false}, {coinKey, true}})
	bag := model.CoinBag{}
	if err := d.bag.FindOne(d.ctx, query, bagDealOp).Decode(&bag); err == nil {
		user.Coin = bag[coinKey]
	}

	// 写登录房间日志
	newLock := bson.D{
		{"_id", newId},
		{"uid", user.Id},
		{"win", zero64},
		{"bet", zero64},
		{"round", zero32},
		{"state", zero32},
		{"kind", kind},
		{"room", roomId},
		{coinKey, user.Coin},
		{"ip", user.Ip},
		{"init", lock.Up},
		{"up", lock.Up},
	}
	if replace {
		newLock = append(newLock, bson.E{Key: "f", Value: lock.Log2})
	}
	d.roomLog.InsertOne(d.ctx, newLock)

	return user, nil
}

// 解锁用户从指定房间
func (d *driver) UnlockUserRoom(agent int64, uid model.UserId, roomId int32, win int64, bet int64, round int32) bool {
	query := bson.D{
		{"_id", uid},
		{"room", roomId},
	}
	up := bson.D{{"$set", bson.D{{"kind", zero32}, {"room", zero32}}}, upNow}
	lock := new(model.UserLocker)
	if err := d.locker.FindOneAndUpdate(d.ctx, query, up, retnew).Decode(lock); err == nil {
		if lock.Log2.IsZero() == false {
			// 更新对应的日志记录
			up1 := bson.D{{"$set", bson.D{{"state", int32(1)}, {"win", win}, {"bet", bet}, {"round", round}}}, upNow}
			_, err = d.roomLog.UpdateOne(d.ctx, bson.D{{"_id", lock.Log2}, {"state", int32(0)}}, up1)
			//log.Debugf("UnlockUserRoom:%v,err:%v", lock.Log2, err)

			// 删除玩家锁
			d.locker.DeleteOne(d.ctx, bson.D{{"_id", uid}, {"agent", zero64}, {"room", zero32}})
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
	d.user.FindOne(d.ctx, bson.D{{"act", account}}).Decode(&res)
	return res.Id
}

// 设置内容
func (d *driver) SetValue(id, v interface{}) error {
	_, err := d.user.UpdateOne(d.ctx, bson.D{{"_id", id}}, bson.D{{"$set", v}}, upsert2)
	return err
}

func (d *driver) UserLogin(req *protocol.LoginReq) (*protocol.LoginSuccessAck, error) {
	return nil, nil
}

// 获取agentId
func (d *driver) CheckUserAgent(uid model.UserId, agent int64) bool {
	query := bson.D{
		{"_id", uid},
		{"agent", agent},
	}
	if changed, err := d.locker.UpdateOne(d.ctx, query, bson.D{upNow}); err == nil {
		return changed.MatchedCount > 0
	}
	return false
}
