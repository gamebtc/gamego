package internal

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	SceneMap = make(map[int32]*SceneSet, 20)
)

type DistrubFishSet struct {
	Time        float64 `yaml:"Time"`
	MinCount    int32   `yaml:"MinCount"`
	MaxCount    int32   `yaml:"MaxCount"`
	FishList    []int32 `yaml:"FishList"`
	Weight      []int32 `yaml:"Weight"`
	RefreshType int32   `yaml:"RefreshType"`
	OffsetX     float64 `yaml:"OffsetX"`
	OffsetY     float64 `yaml:"OffsetY"`
	OffsetTime  float64 `yaml:"OffsetTime"`
}

type TroopSet struct {
	BeginTime float64 `yaml:"BeginTime"`
	EndTime   float64 `yaml:"EndTime"`
	TroopID   int32   `yaml:"Id"`
}

type SceneSet struct {
	Id          int32            `yaml:"Id"`
	Next        int32            `yaml:"Next"`
	Image       string           `yaml:"Image"`
	Time        float64          `yaml:"Time"`
	TroopList   []TroopSet       `yaml:"TroopList"`
	DistrubFish []DistrubFishSet `yaml:"DistrubFish"`
}

func LoadScene(fileName string) bool {
	var config struct {
		Scene []SceneSet `yaml:"scene"`
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("path config file not exists:%v", err)
		return false
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("path config file error:%v", err)
		return false
	}

	for i := 0; i < len(config.Scene); i++ {
		v := &config.Scene[i]
		SceneMap[v.Id] = v
	}

	return true
}
