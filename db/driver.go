package db

import (
	"time"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

var(
	Driver GameDriver
)


type GameDriver interface {
	// 获取包配置
	GetPackConf(id int32) *model.PackInfo
	// 获取渠道配置
	GetChanConf(code string) *model.ChanInfo
	// 获取IP配置
	GetIpInfo(ip model.IP) *model.IpInfo
	// 禁止机器码的到期时间
	GetMachineInfo(id string) *model.MachineInfo
	// 获取单个房间信息
	GetRoom(roomId int32, ver int32) (*model.RoomInfo, error)
	// 锁定房间服务
	LockRoomServer(room *msg.RoomConfig) (obj *model.RoomInfo, err error)
	// 获取所有房间
	GetAllRoom(query interface{}) ([]*model.RoomInfo, error)
	// 获取账号
	GetAccount(app int32, t int32, name string) (*model.Account, error)
	// 创建账号
	CreateAccount(*model.Account, *msg.LoginReq) error
	// 创建用户
	CreateUser(*model.User, *msg.LoginReq) error
	// 锁定用户
	LockUser(int64, *model.User, *msg.LoginReq) (*model.UserLocker, error)
	// 解锁用户
	UnlockUser(agent int64, userId int32) bool
	// 创建新的ID
	NewId() (id model.ObjectId, e error)
	// 获取数据库时间
	Now() time.Time
	// 获取用户
	GetUser(id int32) (user *model.User, err error)
	// 加载用户
	LoadUser(*model.User) error

    // 加载机器人
    LoadRobot(room int32, count int32)[]*model.User
	// 卸载机器人
	UnloadRobot(room int32, ids[]int32)
	// 清理机器人
	ClearRobot(room int32)

	// 锁定用户到指定房间
	LockUserRoom(agent int64, userId int32, kind int32, roomId int32) (*model.User, error)
	// 解锁用户从指定房间
	UnlockUserRoom(agent int64, userId int32, roomId int32) bool
	// 更新agentId
	CheckUserAgent(id int32, agent int64) bool

	// 根据消息号获取消息
	GetHint(code int32) string

	// 执行命令
	Exec(args []interface{}, result interface{}) *model.ExecError
	// 执行命令
	Exec2(args []interface{}, result interface{}) error

	// 更新监听
	Watch([]string) error
	// 刷新数据缓存
	Refresh(map[string]int32) error

	Close()

	// 获取新的交易序列号
	NewSN(key interface{}, count int64) int64

	// 钱包交易
	BagDeal(coinKey string, flow *model.CoinFlow) error
	// 钱包安全交易
	BagDealSafe(coinKey string, flow *model.CoinFlow) error
	// 转账
	BagDealTransfer(form string, to string, flow *model.CoinFlow, lockRoom bool) error
	// 保存日志
	SaveLog(collName string, value interface{}) (err error)

	// 获取机器人
	GetRobot(room int32, count int32) []*model.User
	// 机器人退出
	ExitRobot(room int32, users []int32)
}
