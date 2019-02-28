package main


import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type dnsfixConfig struct {
	Dns      []string            `yaml:"dns"`
	Host     map[string][]string `yaml:"host"`
	Pprof    string              `yaml:"pprof"`
	LogLevel string              `yaml:"logLevel"`
	Period   int                 `yaml:"period"`
}

func InitConfig(path string)(*dnsfixConfig) {
	cfg := &dnsfixConfig{}
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

