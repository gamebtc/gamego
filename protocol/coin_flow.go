package protocol

// 游戏币流水
type CoinFlow struct {
	Sn     int64       `json:"sn" bson:"sn" msg:"sn"`                               //交易序列号,全局交易中唯一
	Uid    int32       `json:"u" bson:"u" msg:"u"`                                  //用户ID
	Add    int64       `json:"a" bson:"a" msg:"a"`                                  //游戏币变化量
	Coin   int64       `json:"c" bson:"c" msg:"c"`                                  //游戏币变化后的金额
	Expect int64       `json:"e,omitempty" bson:"e,omitempty" msg:"e,omitempty"`    //期望值
	Tax    int64       `json:"x,omitempty" bson:"x,omitempty" msg:"x,omitempty"`    //税收
	Kind   int32       `json:"-" bson:"-" msg:"-"`                                  //所在游戏
	Room   int32       `json:"r,omitempty" bson:"r,omitempty" msg:"r,omitempty"`    //所在房间ID
	Type   int32       `json:"t,omitempty" bson:"t,omitempty" msg:"t,omitempty"`    //原因
	State  int32       `json:"s" bson:"s" msg:"s"`                                  //状态
	Note   string      `json:"n,omitempty" bson:"n,omitempty" msg:"n,omitempty"`    //备注
	Att    interface{} `json:"at,omitempty" bson:"at,omitempty" msg:"at,omitempty"` //附件
}

// 用户行为
type UserAction struct {
	Start  int64	`json:"s" bson:"s" msg:"s"`		//开始时间
	Uid    int32	`json:"u" bson:"u" msg:"u"`		//用户ID
	Type   int32    `json:"t" bson:"t" msg:"t"` 	//类型
	Arg    interface{} `json:"a,omitempty" bson:"a,omitempty" msg:"a,omitempty"` //参数
}