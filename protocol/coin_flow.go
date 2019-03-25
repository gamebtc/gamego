package protocol


// 用户行为
type UserAction struct {
	Start  int64	`json:"s" bson:"s" msg:"s"`		//开始时间
	Uid    int32	`json:"u" bson:"u" msg:"u"`		//用户ID
	Type   int32    `json:"t" bson:"t" msg:"t"` 	//类型
	Arg    interface{} `json:"a,omitempty" bson:"a,omitempty" msg:"a,omitempty"` //参数
}