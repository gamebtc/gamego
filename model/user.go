package model

import (
	"time"
)

// 支付信息
type PayInfo struct {
	RealName string `bson:"rn"`   //实名
	Tel      string `bson:"tel"`  //电话号
	Ali      string `bson:"ali"`  //支付宝账号
	AliName  string `bson:"an"`   //支付宝姓名
	BankCard string `bson:"card"` //银行卡号
	CardName string `bson:"cn"`   //持卡人姓名
	BankName string `bson:"bn"`   //银行名称
	BankCode string `bson:"bc"`   //银行名称
}

// 账号，用于登录，一个账号下可以创建多个玩家号
type Account struct {
	Id    AccountId `bson:"_id"`   //账号唯一ID
	App   int32     `bson:"app"`   //所属应用(同一应用类型的客户端可以互通)
	Type  int32     `bson:"type"`  //账号类型(0:机器码+密码；1:用户名+密码；2:手机号+密码；其它：第三方登录)
	Name  string    `bson:"name"`  //账号名称
	Pwd   string    `bson:"pwd"`   //登录密码
	Token string    `bson:"token"` //账号会话
	Phone string    `bson:"phone"` //绑定的手机号
	Pay   PayInfo   `bson:"pay"`   //支付信息
	State int32     `bson:"state"` //账号状态(0:正常)
	Users []UserId  `bson:"users"` //账号下关联的玩家ID
	Pack  int32     `bson:"pack"`  //所属包
	Chan  int32     `bson:"chan"`  //所属渠道
	Ip0   IP        `bson:"ip0"`   //创建时的IP
	Udid  string    `bson:"udid"`  //创建时的机器码
	Tag   TagType   `bson:"tag"`   //标签
	Init  time.Time `bson:"init"`  //创建时间
	Up    time.Time `bson:"up"`    //更新时间
}

// User struct: 玩家基本信息
type User struct {
	Id     UserId    `bson:"_id"`    //唯一ID
	Icon   int32     `bson:"icon"`   //头像
	Vip    int32     `bson:"vip "`   //vip等级
	Name   string    `bson:"name"`   //玩家昵称
	App    int32     `bson:"app"`    //所属应用(同一应用类型的客户端可以互通)
	Sex    int32     `bson:"sex"`    //性别
	State  int32     `bson:"state"`  //账号状态(0:正常)
	Job    int32     `bson:"job "`   //玩家类型(0:用户；1:测试；2:管理；10:机器人)
	Risk   int32     `bson:"risk"`   //用户风险
	Act    AccountId `bson:"act"`    //关联的账号
	Pack   int32     `bson:"pack"`   //包ID
	Chan   int32     `bson:"chan"`   //渠道ID
	Ip0    IP        `bson:"ip0"`    //创建时的IP
	Last   time.Time `bson:"last"`   //最后登录时间
	Ip     IP        `bson:"ip"`     //最后登录时间
	Tag    TagType   `bson:"tag"`    //标签
	Init   time.Time `bson:"init"`   //创建时间
	Up     time.Time `bson:"up"`     //更新时间
	Ver    int64     `bson:"ver"`    //币的版本
	Bag    CoinBag   `bson:"-"`      //玩家的背包
	Coin   int64     `bson:"-"`      //携带金币
	Online bool      `bson:"-"`      //是否在线
	FlowSn int64     `bson:"-"`      //最后的写分序号,返回时用于验证
}

// 在线玩家锁定信息
type UserLocker struct {
	Id    UserId    `bson:"_id"`   //唯一ID
	Agent int64     `bson:"agent"` //前端代理号
	Ip    IP        `bson:"ip"`    //登录IP
	Kind  int32     `bson:"kind"`  //所在游戏
	Room  int32     `bson:"room"`  //所在房间ID
	Tab   int32     `bson:"tab"`   //所在桌子ID
	Init  time.Time `bson:"init"`  //创建时间
	Up    time.Time `bson:"up"`    //更新时间
	Log1  ObjectId  `bson:"log1"`  //登录大厅日志ID
	Log2  ObjectId  `bson:"log2"`  //登录房间日志ID
}
