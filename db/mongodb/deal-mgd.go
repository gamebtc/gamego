package mongodb

import (
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/abc/game/model"
)

var errorEmptySn = errors.New("sn is empty")
var errorLookRoom = errors.New("lock room fail")
var sessionOp = options.Session()

func (d *driver) deal(coinKey string, flow *model.CoinFlow, safe bool) error {
	f := func(sc mongo.SessionContext) (err error) {
		collName := coinKey + "_" + strconv.Itoa(int(flow.Kind))
		logDeal := d.GetColl(collName)
		if err = sc.StartTransaction(); err != nil {
			return err
		}

		add := flow.Add
		var query bson.D
		if safe && add < 0 {
			query = bson.D{{"_id", flow.Uid}, {coinKey, bson.D{{"$gte", -add}}}}
		} else {
			query = bson.D{{"_id", flow.Uid}}
		}
		up := bson.D{{"$inc", bson.D{{"ver", int64(1)}, {coinKey, flow.Add}}}, upNow}
		bagDealOp := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.D{{"_id", false}, {coinKey, true}})
		bag := model.CoinBag{}
		if err = d.bag.FindOneAndUpdate(sc, query, up, bagDealOp).Decode(&bag); err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		tmpNew := flow.New
		flow.New = bag[coinKey]
		var ir *mongo.InsertOneResult
		ir, err = logDeal.InsertOne(sc, flow)
		if err != nil || ir.InsertedID == nil {
			sc.AbortTransaction(sc)
			//交易号重复，记录插入错误
			flow.New = tmpNew
			d.GetColl(collName+"_rpt").InsertOne(sc, flow)
			return err
		}

		return sc.CommitTransaction(sc)
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
		if err = sc.StartTransaction(); err != nil {
			return err
		}

		// 锁定玩家到大厅房间1，房间1代表钱包转账中
		if lockRoom {
			var ur *mongo.UpdateResult
			ur, err = d.locker.UpdateOne(sc, bson.D{{"_id", uid}, {"room", zero32}}, bson.D{{"$set", bson.D{{"room", int32(1)}}}})
			if err != nil || ur == nil || ur.ModifiedCount != 1 {
				sc.AbortTransaction(sc)
				if err == nil {
					err = errorLookRoom
				}
				return
			}
		}

		query := bson.D{{"_id", uid}, {from, bson.D{{"$gte", add}}}}
		up := bson.D{{"$inc", bson.D{{"ver", int64(1)}, {from, -add}, {to, add}}}, upNow}
		bagDealOp := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.D{{"_id", false}, {from, true}, {to, true}})
		bag := model.CoinBag{}
		err = d.bag.FindOneAndUpdate(sc, query, up, bagDealOp).Decode(&bag)
		if err != nil {
			sc.AbortTransaction(sc)
			return err
		}

		fromCoin, toCoin := bag[from], bag[to]
		flow.Add = -add
		flow.New = fromCoin
		flow.Old = fromCoin + add
		var ir1 *mongo.InsertOneResult
		ir1, err = logFrom.InsertOne(sc, flow)
		if err != nil || ir1.InsertedID == nil {
			sc.AbortTransaction(sc)
			return err
		}

		flow.Add = add
		flow.New = toCoin
		flow.Old = toCoin - add
		var ir2 *mongo.InsertOneResult
		ir2, err = logTo.InsertOne(sc, flow)
		if err != nil || ir2.InsertedID == nil {
			sc.AbortTransaction(sc)
			return err
		}

		if lockRoom {
			// 解锁
			var ur *mongo.UpdateResult
			ur, err = d.locker.UpdateOne(sc, bson.D{{"_id", uid}, {"room", int32(1)}}, bson.D{{"$set", bson.D{{"room", zero32}}}})
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
	return d.Client().UseSessionWithOptions(d.ctx, sessionOp, f)
}
