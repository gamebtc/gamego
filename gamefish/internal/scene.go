package internal

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	SceneSets []SceneSet
)

type DistrubFishSet struct {
	FishWeight  `yaml:",inline"`
	Time        float64 `yaml:"Time"`
	MinCount    int32   `yaml:"MinCount"`
	MaxCount    int32   `yaml:"MaxCount"`
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

func (t *SceneSet) fixWeight() {
	for i := 0; i < len(t.DistrubFish); i++ {
		t.DistrubFish[i].fixWeight()
	}
}

func (t *SceneSet) findTroopSet(time float64)(int, *TroopSet) {
	for i := 0; i < len(t.TroopList); i++ {
		if time >= t.TroopList[i].BeginTime && time < t.TroopList[i].EndTime {
			return i, &t.TroopList[i]
		}
	}
	return -1, nil
}

func LoadScene(fileName string) bool {
	var config struct {
		Scene []SceneSet `yaml:"Scene"`
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
		config.Scene[i].fixWeight()
	}
	SceneSets = config.Scene

	return true
}
