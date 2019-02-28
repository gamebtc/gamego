package model


type Rank struct {
	Id    int32 `bson:"_id"` //ID
	Rank  int32 `bson:"r"`   //结算排行
	Coin int32 `bson:"s"`   //积分
}
