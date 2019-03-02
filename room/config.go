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
	. "local.com/abc/game/msg"
)

// 房间配置
var(
	Config    model.RoomInfo
	RoomId    int32   //房间唯一ID
	KindId    int32   //游戏分类
	CoinKey   string  //金币类型
)

type AppConfig struct {
	Consul   ConsulConfig   `yaml:"consul"`
	Tcp      TcpConfig      `yaml:"tcp"`
	Udp      UdpConfig      `yaml:"udp"`
	Kcp      KcpConfig      `yaml:"kcp"`
	Grpc     GrpcConfig     `yaml:"grpc"`
	Pprof    string         `yaml:"pprof"`
	LogLevel string         `yaml:"logLevel"`
	LogFile  string         `yaml:"logFile"`
	Database DatabaseConfig `yaml:"database"`
	Room     RoomConfig     `yaml:"room"` //房间配置
	Codec    string         `yaml:"codec"`
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

	coder = GetCoder(cfg.Codec)

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
