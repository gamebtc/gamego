package conf

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	. "local.com/abc/game/protocol"
)

type AppConfig struct {
	Consul       ConsulConfig `yaml:"consul"`
	Tcp          TcpConfig    `yaml:"tcp"`
	Udp          UdpConfig    `yaml:"udp"`
	Kcp          KcpConfig    `yaml:"kcp"`
	MaxConnect   int32        `yaml:"maxConnect"`
	SendChanLen  int          `yaml:"sendChanLen"`
	RecvChanLen  int          `yaml:"recvChanLen"`
	ReadTimeout  int          `yaml:"readTimeout"`
	WriteTimeout int          `yaml:"writeTimeout"`
	RpmLimit     int32        `yaml:"rpmLimit"`
	Pprof        string       `yaml:"pprof"`
	LogLevel     string       `yaml:"logLevel"`
	SlowOp       uint32       `yaml:"slowOp"`
	AgentId      uint32       `yaml:"agentId"`
	SameIp       uint32       `yaml:"sameIp"`
	Seed         int32        `yaml:"seed"`
	Codec        string       `yaml:"codec"`
	VerSign 	 bool         `yaml:"verSign"`
}

func InitConfig(path string)(*AppConfig) {
	cfg := &AppConfig{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("app config file not exists:%v", err)
		return nil
	}

	if err = yaml.Unmarshal(data, cfg); err != nil {
		log.Fatalf("app config file error:%v", err)
		return nil
	}

	if lv, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(lv)
	}
	return cfg
}