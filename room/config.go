package room

import (
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
)

// 房间配置
var(
	Config    model.RoomInfo
	RoomId    int32   //房间唯一ID
	KindId    int32   //游戏分类
	CoinKey   string  //金币类型
)

type AppConfig struct {
	Consul   protocol.ConsulConfig   `yaml:"consul"`
	Tcp      protocol.TcpConfig      `yaml:"tcp"`
	Udp      protocol.UdpConfig      `yaml:"udp"`
	Kcp      protocol.KcpConfig      `yaml:"kcp"`
	Grpc     protocol.GrpcConfig     `yaml:"grpc"`
	Database protocol.DatabaseConfig `yaml:"database"`
	Room     protocol.RoomConfig     `yaml:"room"` //房间配置
	Codec    string                  `yaml:"codec"`
	Pprof    string                  `yaml:"pprof"`
	LogLevel string                  `yaml:"logLevel"`
	LogFile  string                  `yaml:"logFile"`
}

func InitConfig(path string) *AppConfig {
	cfg := &AppConfig{}
	if data, err := ioutil.ReadFile(path); err != nil {
		log.Fatal("app config file not exists:%v", err)
	} else {
		if err = yaml.Unmarshal(data, cfg); err != nil {
			log.Fatal("app config file error:%v", err)
		}
	}
	if lv, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(lv)
	}

	coder = protocol.GetCoder(cfg.Codec)

	if cfg.Pprof != "" {
		go http.ListenAndServe(cfg.Pprof, nil)
	}

	if cfg.LogFile != "" {
		f, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE, 0755)
		if err == nil {
			mw := io.MultiWriter(os.Stdout, f)
			log.SetOutput(mw)
		}
	}
	return cfg
}
