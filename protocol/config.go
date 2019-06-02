package protocol

type IP = int64
type UserId = int32

type ConsulConfig struct {
	Addr         string   `yaml:"addr"`
	Ttl          int      `yaml:"ttl"`
	ServerPrefix string   `yaml:"serverPrefix"`
	RoomPrefix   string   `yaml:"roomPrefix"`
	Services     []int32  `yaml:"services"`
}

type GrpcConfig struct {
	Listen     string            `yaml:"listen"`
	Name       string            `yaml:"name"`
	Addr       string            `yaml:"addr"`
	Port       int               `yaml:"port"`
	Tags       []string          `yaml:"tags"`
	Dcsa       uint              `yaml:"dcsa"`
	Interval   uint              `yaml:"interval"`
	MaxConnect int               `yaml:"maxConnect"`
	Meta       map[string]string `yaml:"meta"`
}

type TcpConfig struct {
	Listen   string `yaml:"listen"`
	ReadBuf  int    `yaml:"readBuf"`
	WriteBuf int    `yaml:"writeBuf"`
}

type UdpConfig struct {
	Listen   string `yaml:"listen"`
	ReadBuf  int    `yaml:"readBuf"`
	WriteBuf int    `yaml:"writeBuf"`
	Dscp     int    `yaml:"dscp"`
}

type KcpConfig struct {
	Sndwnd   int `yaml:"sndwnd"`
	Rcvwnd   int `yaml:"rcvwnd"`
	Mtu      int `yaml:"mtu"`
	Nodelay  int `yaml:"nodelay"`
	Interval int `yaml:"interval"`
	Resend   int `yaml:"resend"`
	Nc       int `yaml:"nc"`
}

type DatabaseConfig struct {
	Driver  string           `yaml:"driver"`
	Url     string           `yaml:"url"`
	Name    string           `yaml:"name"`
	Watch   []string         `yaml:"watch"`
	Refresh map[string]int32 `yaml:"refresh"`
}

// 房间服务配置
type RoomConfig struct {
	Id   int32    `yaml:"id" bson:"_id"`    //房间唯一ID
	Game int32    `yaml:"game" bson:"game"` //游戏分类
	Key  string   `yaml:"key" bson:"-"`     //服务器KEY
	Addr []string `yaml:"addr" bson:"addr"` //服务器地址
}