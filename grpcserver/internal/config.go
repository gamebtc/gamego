package internal

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

type AppConfig struct {
	Consul   msg.ConsulConfig   `yaml:"consul"`
	Grpc     msg.GrpcConfig     `yaml:"grpc"`
	Database msg.DatabaseConfig `yaml:"database"`
	InitBag  model.CoinBag      `yaml:"initBag"`
	Pprof    string             `yaml:"pprof"`
	LogLevel string             `yaml:"logLevel"`
	SlowOp   uint32             `yaml:"slowOp"`
	AgentId  uint32             `yaml:"agentId"`
	Codec    string             `yaml:"codec"`
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
	return cfg
}

func (config *AppConfig)GetInitCoin() model.CoinBag {
	r1 := make(model.CoinBag, len(config.InitBag))
	if config.InitBag != nil {
		for k, v := range config.InitBag {
			r1[k] = v
		}
	}
	return r1
}
