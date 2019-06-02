package mongodb

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"local.com/abc/game/model"
)

//go test -v -run="TestMsg"
func TestMsg(t *testing.T) {
	val := new(model.Account)
	registry := bson.NewRegistryBuilder().Build()

	buf := make([]byte, 0, 256)
	b, err := bson.MarshalAppendWithRegistry(registry, buf, val)
	if err != nil {
		t.Fatalf("[err%v]", err)
	}
	t.Logf("[b%v]", b)
}

type Point struct {
	X int
	Y int
	Z int
}

//go test -v -run="TestSession"
func TestSession(t *testing.T) {
	url := "mongodb://127.0.0.1:27088,127.0.0.1:27089,127.0.0.1:27090/admin?replicaSet=gameRs"
	if client, err := mongo.NewClient(options.Client().ApplyURI(url)); err != nil {
		t.Fatal(err)
	} else {
		if err = client.Connect(nil); err != nil {
			t.Fatal(err)
			return
		}
		op := options.Session()
		ctx := context.WithValue(context.Background(), "game", "test")

		f1 := func(s mongo.SessionContext) error {
			point := &Point{6566, 2, 3}
			db := client.Database("game")
			c1 := db.Collection("coll01")
			c2 := db.Collection("coll02")

			s.StartTransaction()
			c1.InsertOne(s, point)
			c2.InsertOne(s, point)
			if point.X > 100 {
				s.CommitTransaction(s)
			} else {
				s.AbortTransaction(s)
			}

			return nil
		}

		f2 := func(s mongo.SessionContext) error {
			point := &Point{34, 2, 3}
			db := client.Database("game")
			c1 := db.Collection("coll01")
			c2 := db.Collection("coll02")

			s.StartTransaction()
			c1.InsertOne(s, point)
			c2.InsertOne(s, point)

			client.UseSessionWithOptions(ctx, op, f1)

			if point.X > 100 {
				s.CommitTransaction(s)
			} else {
				s.AbortTransaction(s)
			}
			return nil
		}

		client.UseSessionWithOptions(ctx, op, f2)
	}
}

func TestIn(t *testing.T) {
	url := "mongodb://127.0.0.1:27088,127.0.0.1:27089,127.0.0.1:27090/admin?replicaSet=gameRs"
	if client, err := mongo.NewClient(options.Client().ApplyURI(url)); err != nil {
		t.Fatal(err)
	} else {
		if err = client.Connect(nil); err != nil {
			t.Fatal(err)
			return
		}

		db := client.Database("game")
		locker := db.Collection("userLocker")

		query := bson.D{
			{"_id", 4743924},
			{"agent", 5593465379},
			{"room", bson.D{{"$in", []int32{0, 11011}}}},
		}
		newId := model.NewObjectId()
		up := bson.D{{"$set", bson.D{{"log2", newId}, {"game", 432}, {"room", 11014}}}, upNow}
		if _, err := locker.UpdateOne(nil, query, up); err != nil {

			t.Fatal(err)
		}

	}
}
