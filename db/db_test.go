package db

import (
	"testing"

	"local.com/abc/game/db/mongodb"
	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

// go test -v -run="指定函数名"
// https://blog.csdn.net/hjmnasdkl/article/details/81304329
// go test -v -run="none" -bench=.    不允许单元测试，运行所有的基准测试
// -benchmem 表示分配内存的次数和字节数，-benchtime="3s" 表示持续3秒

//// 钱包交易
//BagDeal(bagGold string, flow *model.CoinFlow) error
//// 银行交易
//BankDeal(bagGold string, flow *model.CoinFlow) error
//// 银行转到钱包
//Bank2Bag(bagGold string, flow *model.CoinFlow) error
//// 钱包转到银行
//Bag2Bank(bagGold string, flow *model.CoinFlow) error

var driver GameDriver
const bagGold = "gc1"
const bankGold = "gc2"
const kind = 0

func getTestDrive(t *testing.T)GameDriver {
	var testUrl = "mongodb://127.0.0.1:27088,127.0.0.1:27089,127.0.0.1:27090/admin?replicaSet=gameRs"
	conf := new(msg.DatabaseConfig)
	conf.Url = testUrl
	conf.Name = "game"
	conf.Driver = "mongodb"
	d, err := mongodb.NewGameDriver(conf)
	if err != nil {
		t.Fatal(err)
	}
	driver = d
	return d
}

func TestBagDeal(t *testing.T) {
	d := getTestDrive(t)
	flow := &model.CoinFlow{
		Sn:    d.NewSN(kind, 1),
		Uid:   2743405,
		Add:   333,
		Coin:  0,
		Kind:  kind,
		Room:  1001,
		Type: 1,
		State: 0,
		Note:  "BagDeal",
	}

	if err := d.BagDeal(bagGold, flow); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkBagDeal(b *testing.B) {
	d := getTestDrive(nil)
	b.ReportAllocs()
	b.ResetTimer()
	flow := &model.CoinFlow{
		Uid:   2743405,
		Add:   444,
		Coin:  0,
		Kind:  0,
		Room:  0,
		Type: 1,
		State: 0,
		Note:  "BagDeal",
	}
	sn := d.NewSN(kind,int64(b.N))
	flow.Sn = sn
	var err error
	for i := 0; i < b.N; i++ {
		flow.Sn = sn + int64(i)
		flow.Room += 1
		flow.Coin = 0
		err = d.BagDeal(bagGold, flow)
	}
	if err != nil {
		b.Fatal(err)
	}
}

func TestBagDealSafe(t *testing.T) {
	d := getTestDrive(t)
	flow := &model.CoinFlow{
		Sn:    d.NewSN(kind,1),
		Uid:   2743405,
		Add:   555,
		Coin:  0,
		Kind:  kind,
		Room:  1001,
		Type: 2,
		State: 0,
		Note:  "BagDealSafe",
	}
	if err := d.BagDealSafe(bankGold, flow); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkBagDealSafe(b *testing.B) {
	d := getTestDrive(nil)
	b.ReportAllocs()
	b.ResetTimer()
	flow := &model.CoinFlow{
		Uid:   2743405,
		Add:   666,
		Coin:  0,
		Kind:  0,
		Room:  0,
		Type: 1,
		State: 0,
		Note:  "BagDealSafe",
	}
	sn := d.NewSN(kind,int64(b.N))
	for i := 0; i < b.N; i++ {
		flow.Sn = sn + int64(i)
		flow.Room += 1
		if err := d.BagDealSafe(bankGold, flow); err != nil {
			b.Fatal(err)
		}
	}
}

func TestBank2Bag(t *testing.T) {
	d := getTestDrive(t)
	flow := &model.CoinFlow{
		Sn:    d.NewSN(kind,1),
		Uid:   2743405,
		Add:   777,
		Coin:  0,
		Kind:  kind,
		Room:  1001,
		Type:  3,
		State: 0,
		Note:  "Bank2Bag",
	}
	if err := d.BagDealTransfer(bankGold, bagGold, flow, false); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkBank2Bag(b *testing.B) {
	d := getTestDrive(nil)
	b.ReportAllocs()
	b.ResetTimer()
	flow := &model.CoinFlow{
		Uid:   2743405,
		Add:   888,
		Coin:  0,
		Kind:  0,
		Room:  0,
		Type:  1,
		State: 0,
		Note:  "Bank2Bag",
	}
	sn := d.NewSN(kind,int64(b.N))
	for i := 0; i < b.N; i++ {
		flow.Sn = sn + int64(i)
		flow.Room += 1
		if err := d.BagDealTransfer(bankGold, bagGold, flow, false); err != nil {
			b.Fatal(err)
		}
	}
}

func TestBag2Bank(t *testing.T) {
	d := getTestDrive(t)
	flow := &model.CoinFlow{
		Sn:    d.NewSN(kind,1),
		Uid:   2743405,
		Add:   999,
		Coin:  0,
		Kind:  kind,
		Room:  1001,
		Type:  4,
		State: 0,
		Note:  "Bag2Bank",
	}
	if err := d.BagDealTransfer(bagGold, bankGold, flow, true); err != nil {
		t.Fatal(err)
	}
}


func BenchmarkBag2Bank(b *testing.B) {
	d := getTestDrive(nil)
	b.ReportAllocs()
	b.ResetTimer()
	flow := &model.CoinFlow{
		Uid:   2743405,
		Add:   333,
		Coin:  0,
		Kind:  kind,
		Room:  0,
		Type:  1,
		State: 0,
		Note:  "Bag2Bank",
	}
	sn := d.NewSN(kind, int64(b.N))
	for i := 0; i < b.N; i++ {
		flow.Sn = sn + int64(i)
		flow.Room += 1
		if err := d.BagDealTransfer(bagGold, bankGold, flow, true); err != nil {
			b.Fatal(err)
		}
	}
}

//go test -v -run="TestLoadUser"
func TestLoadUser(t *testing.T) {

	d := getTestDrive(t)

	if u1, err := d.GetUser(3313004); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("user:%v", u1)
	}

	//user := &model.User{}
	//user.Id = 3313004
	//user.Ip = 3313004
	//
	//if err := d.LoadUser(user); err != nil {
	//	t.Fatal(err)
	//}
}