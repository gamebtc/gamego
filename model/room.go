package model

import (
	"time"

	"local.com/abc/game/msg"
)

// 房间信息
type RoomInfo struct {
	Id      int32     `bson:"_id"`     //房间唯一ID
	Kind    int32     `bson:"kind"`    //游戏分类
	Name    string    `bson:"name"`    //房间名字
	Tax     int64     `bson:"tax"`     //房间税率(千分比)
	Level   int32     `bson:"level"`   //房间等级
	Cap     int32     `bson:"cap"`     //房间人数容量
	Ante    int64     `bson:"ante"`    //房间底注
	DoorMin int64     `bson:"doorMin"` //进入房间最低限度
	DoorMax int64     `bson:"doorMax"` //进入房间最高限度
	StayMin int64     `bson:"stayMin"` //停留在房间最低限度
	StayMax int64     `bson:"stayMax"` //停留在房间最高限度
	PlayMin int64     `bson:"playMin"` //玩游戏的最低限度
	PlayMax int64     `bson:"playMax"` //玩游戏的最高限度
	CoinKey string    `bson:"coinKey"` //货币类型
	Tab     int32     `bson:"tab"`     //桌子数目,0不限制
	Seat    int32     `bson:"seat"`    //每一桌的座位数目,0不限制
	Icon    int32     `bson:"icon"`    //房间图标
	Show    int32     `bson:"show"`    //是否显示
	Period  int       `bson:"period"`  //帧更新周期(毫秒)
	Pause   int32     `bson:"pause"`   //暂停
	State   int32     `bson:"state"`   //房间状态(0:不可用，1：可用)
	Jobs    []int32   `bson:"jobs"`    //只对特定人员开放(0:无锁任何人可进，其它：锁定，指定类型的玩家可以进入)
	Packs   []int32   `bson:"packs"`   //只对特定的包ID开放
	Lock    int32     `bson:"lock"`    //是否锁定(锁定后，只出不进)
	Close   int32     `bson:"close"`   //是否关闭中(0:任何人可进，1:设置为关闭状态,不再开始新的游戏)
	Sort    int32     `bson:"sort"`    //房间排序
	Addr    string    `bson:"addr"`    //服务器地址
	Key     string    `bson:"key"`     //服务器KEY
	Init    time.Time `bson:"init"`    //创建时间
	Up      time.Time `bson:"up"`      //更新时间
	Ver     int32     `bson:"ver"`     //房间版本
	Conf    Raw       `bson:"conf"`    //其它配置项
	Cache   Raw       `bson:"cache"`   //其它缓存
	WinRate int32     `bson:"winRate"` //必赢局概率千分比(0-1000)
}

func (room *RoomInfo) GetMsg() *msg.RoomInfo {
	return &msg.RoomInfo{
		Id:      room.Id,
		Kind:    room.Kind,
		Name:    room.Name,
		Ante:    room.Ante,
		DoorMin: room.DoorMin,
		DoorMax: room.DoorMax,
		StayMin: room.StayMin,
		StayMax: room.StayMax,
		PlayMin: room.PlayMin,
		PlayMax: room.PlayMax,
		CoinKey: room.CoinKey,
		Icon:    room.Icon,
	}
}