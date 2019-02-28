// +build !mgo

package mongodb

import (
	"strconv"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/pkg/errors"

	"local.com/abc/game/model"
)

type Bag struct {
	Bag model.CoinBag `bson:"bag"` //玩家的钱包
}

var errorEmptySn = errors.New("sn is empty")
var errorLookRoom = errors.New("lock room fail")
var sessionOp = options.Session()

const bagPrefix = "bag."

var bagDealOp = options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.D{{"_id", false}, {"bag", true}})

func (d *driver) deal(coinKey string, flow *model.CoinFlow, safe bool) error {
	f := func(sc mongo.SessionContext) (err error) {
		filed := bagPrefix + coinKey
		collName := coinKey + "_" + strconv.Itoa(int(flow.Kind))
		logDeal := d.GetColl(collName)
		if err = sc.StartTransaction(); err == nil {
			add := flow.Add
			var query bson.D
			if safe && add < 0 {
				query = bson.D{{"_id", flow.Uid}, {filed, bson.D{{"$gte", -add}}}}
			} else {
				query = bson.D{{"_id", flow.Uid}}
			}
			up := bson.D{{"$inc", bson.D{{"bag.v", int64(1)}, {filed, flow.Add}}}}
			sr := d.user.FindOneAndUpdate(sc, query, up, bagDealOp)
			if err = sr.Err(); err == nil {
				b := Bag{}
				if err = sr.Decode(&b); err == nil && b.Bag != nil {
					if coin, ok := b.Bag[coinKey]; ok {
						flow.Coin = coin
						var ir *mongo.InsertOneResult
						if ir, err = logDeal.InsertOne(sc, flow); err == nil && ir.InsertedID != nil {
							return sc.CommitTransaction(sc)
						} else {
							sc.AbortTransaction(sc)
							//交易号重复，记录插入错误
							flow.Coin = flow.Expect
							d.GetColl(collName+"_rpt").InsertOne(sc, flow)
							return nil
						}
					}
				}
			}
			sc.AbortTransaction(sc)
		}
		return err
	}
	return d.Client().UseSessionWithOptions(d.ctx, sessionOp, f)
}

// 钱包交易
func (d *driver) BagDeal(coinKey string, flow *model.CoinFlow) error {
	return d.deal(coinKey, flow, false)
}

func (d *driver) BagDealSafe(coinKey string, flow *model.CoinFlow) error {
	return d.deal(coinKey, flow, true)
}

// 转账
func (d *driver) BagDealTransfer(from string, to string, flow *model.CoinFlow, lockRoom bool) error {
	if flow == nil || flow.Add <= 0 {
		return nil
	}
	f := func(sc mongo.SessionContext) (err error) {
		uid := flow.Uid
		add := flow.Add
		logFrom := d.Collection(from + "_" + strconv.Itoa(int(flow.Kind)))
		logTo := d.Collection(to + "_" + strconv.Itoa(int(flow.Kind)))
		if err = sc.StartTransaction(); err == nil {
			// 锁定玩家到大厅房间1，房间1代表钱包转账中
			if lockRoom {
				var ur *mongo.UpdateResult
				ur, err = d.locker.UpdateOne(sc, bson.D{{"_id", uid}, {"room", zeroInt32}}, bson.D{{"$set", bson.D{{"room", int32(1)}}}})
				if err != nil || ur == nil || ur.ModifiedCount != 1 {
					sc.AbortTransaction(sc)
					if err == nil {
						err = errorLookRoom
					}
					return
				}
			}
			fromFiled := bagPrefix + from
			toFiled := bagPrefix + to
			up := bson.D{{"$inc", bson.D{{"bag.v", int64(1)}, {fromFiled, -add}, {toFiled, add}}}}
			query := bson.D{{"_id", uid}, {fromFiled, bson.D{{"$gte", add}}}}
			sr := d.user.FindOneAndUpdate(sc, query, up, bagDealOp)
			if err = sr.Err(); err == nil {
				b := Bag{}
				if err = sr.Decode(&b); err == nil && b.Bag != nil {
					fromCoin, ok2 := b.Bag[from]
					toCoin, ok := b.Bag[to]
					if ok && ok2 {
						flow.Add = -add
						flow.Coin = fromCoin
						var ir1 *mongo.InsertOneResult
						if ir1, err = logFrom.InsertOne(sc, flow); err == nil && ir1.InsertedID != nil {
							flow.Add = add
							flow.Coin = toCoin
							var ir2 *mongo.InsertOneResult
							if ir2, err = logTo.InsertOne(sc, flow); err == nil && ir2.InsertedID != nil {
								if lockRoom {
									// 解锁
									var ur *mongo.UpdateResult
									ur, err = d.locker.UpdateOne(sc, bson.D{{"_id", uid}, {"room", int32(1)}}, bson.D{{"$set", bson.D{{"room", zeroInt32}}}})
									if err != nil || ur == nil || ur.ModifiedCount != 1 {
										sc.AbortTransaction(sc)
										if err == nil {
											err = errorLookRoom
										}
										return
									}
								}
								return sc.CommitTransaction(sc)
							}
						}
					}
				}
			}
			sc.AbortTransaction(sc)
		}
		return err
	}
	return d.Client().UseSessionWithOptions(d.ctx, sessionOp, f)
}
