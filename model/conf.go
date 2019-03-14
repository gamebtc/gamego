package model

import(
	"time"

	"local.com/abc/game/msg"
)

//玩家类型(0:用户；1:代理；10:测试；11:管理；12:机器人)
const JobPlayer = 0
//const JobAgent = 1
const JobTest = 10
const JobManager = 11
const JobRobot = 12

type UserId = msg.UserId
type IP = msg.IP
type CoinBag = map[string]int64
type ConfType = map[string]string
type TagType = map[string]int32
type BanType = []int32

type ExecError struct{
	Code    int32
	Msg     string
}

func (self *ExecError)Error() string{
	return self.Msg
}

// 应用程序信息
type AppInfo struct {
	Id   int32     `bson:"_id"`  //应用ID
	Code string    `bson:"code"` //应用编码，唯一
	Name string    `bson:"name"` //应用名称
	Pack []int32   `bson:"pack"` //应用下面的包
	Url  string    `bson:"url"`  //应用下载地址
	Conf ConfType  `bson:"conf"` //应用配置
	Up   time.Time `bson:"up"`   //更新时间
	Ver  int32     `bson:"ver"`  //版本
}

// 包信息
type PackInfo struct {
	Id     int32     `bson:"_id"`    //包ID
	Code   string    `bson:"code"`   //包编码，唯一
	Name   string    `bson:"name"`   //包名称
	App    int32     `bson:"app"`    //所属应用ID(同一应用的客户端可以互通)
	Tmp    int32     `bson:"tmp"`    //短信模板ID
	State  int32     `bson:"state"`  //包状态
	Url    string    `bson:"url"`    //新包下载地址
	Conf   ConfType  `bson:"conf"`   //包配置
	CanVer []string  `bson:"canVer"` //应用可用版本
	Ban    BanType   `bson:"ban"`    //禁止注册/登录
	Note   string    `bson:"note"`   //备注
	Up     time.Time `bson:"up"`     //更新时间
	Ver    int32     `bson:"ver"`    //版本
}

// 渠道信息
type ChanInfo struct {
	Id       int32     `bson:"_id"`      //包ID
	Code     string    `bson:"code"`     //包编码，唯一
	Name     string    `bson:"name"`     //包名称
	App      int32     `bson:"app"`      //所属应用ID(同一应用的客户端可以互通)
	Parent   int32     `bson:"parent"`   //上级
	State    int32     `bson:"state"`    //包状态
	Conf     ConfType  `bson:"conf"`     //包配置
	Ban      BanType   `bson:"ban"`      //禁止注册/登录
	CanPlay  []int32   `bson:"canPlay"`  //游戏客户端显示的游戏
	Note     string    `bson:"note"`     //备注
	Init     time.Time `bson:"init"`     //初始时间
	Up       time.Time `bson:"up"`       //更新时间
	Ver      int32     `bson:"ver"`      //版本
}

// 提示信息
type HintInfo struct {
	Id  int32  `bson:"_id"` //包ID
	Msg string `bson:"msg"` //包ID
	Ver int32  `bson:"ver"` //版本
}

// IP信息
type IpInfo struct {
	Id      int64     `bson:"_id"`     //IP
	Risk    int32     `bson:"risk"`    //风险
	Proxy   int32     `bson:"risk"`    //是否是代理
	Ports   []int32   `bson:"ports"`   //开放的端口
	Tag     TagType   `bson:"tag"`     //标签
	Ban     BanType   `bson:"ban"`     //禁止注册/登录
	Country string    `bson:"country"` //国家
	Region  string    `bson:"region"`  //省份
	City    string    `bson:"city"`    //城市
	TalReg  int32     `bson:"talReg"`  //总注册人数
	DayReg  int32     `bson:"dayReg"`  //当天注册人数
	LastReg time.Time `bson:"lastReg"` //最后一个注册成功的时间
	Init    time.Time `bson:"init"`    //初始时间
	Up      time.Time `bson:"up"`      //更新时间
	Ver     int32     `bson:"ver"`     //版本
}

// 机器信息
type MachineInfo struct {
	Id   string    `bson:"_id"`  //IP
	Risk int32     `bson:"risk"` //风险
	Tag  TagType   `bson:"tag"`  //标签
	Ban  BanType   `bson:"ban"`  //禁止注册/登录
	Init time.Time `bson:"init"` //初始时间
	Up   time.Time `bson:"up"`   //更新时间
	Ver  int32     `bson:"ver"`  //版本
}

type RoomToken struct {
	Id      int32     `bson:"_id"`     //房间唯一ID
	Kind    int32     `bson:"kind"`    //游戏分类
	Key     string    `bson:"key"`     //服务器KEY
	Addr    string    `bson:"addr"`    //服务器地址
}