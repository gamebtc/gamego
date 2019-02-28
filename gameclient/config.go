package main

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"local.com/abc/game/model"
	. "local.com/abc/game/msg"
)

// 房间配置
var Config *model.RoomInfo

type AppConfig struct {
	Room         RoomConfig   `yaml:"room"` //房间配置
	Tcp          TcpConfig    `yaml:"tcp"`
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
	Server       string       `yaml:"server"`
}

func InitConfig(path string)(*AppConfig) {
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
	return cfg
}