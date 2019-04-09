package internal

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

var (
	KingFishMap    = make(map[int32]*SpecialFish, 20) // 鱼王
	SanYuanFishMap = make(map[int32]*SpecialFish, 20) // 三元
	SiXiFishMap    = make(map[int32]*SpecialFish, 20) // 四喜
	BBXMap         = make(map[int32]*BBX, 100)        //
	FishMap        = make(map[int32]*Fish, 100)
)

// 鱼王/三元/四喜
type SpecialFish struct {
	Id               int32   `yaml:"Id"`               // ID
	Probability      float64 `yaml:"Probability"`      // 产生几率
	MaxScore         int32   `yaml:"MaxScore"`         // 最大倍率
	CatchProbability float64 `yaml:"CatchProbability"` // 捕获几率
	VisualScale      float64 `yaml:"VisualScale"`      // 缩放
	VisualId         int32   `yaml:"VisualId"`         // 视图ID
	BoundingBox      int32   `yaml:"BoundingBox"`
	LockLevel        int32   `yaml:"LockLevel"`
}

type BoundingBox struct {
	Radio float64 `yaml:"R"`
	X     float64 `yaml:"X"`
	Y     float64 `yaml:"Y"`
}

type BBX struct {
	Id     int32         `yaml:"Id"`
	BBList []BoundingBox `yaml:"BB"`
}

type Effect struct {
	Id    int32   `yaml:"Id"`
	Param []int32 `yaml:"Param"`
}

type Buffer struct {
	Id    int32   `yaml:"Id"`
	Param float64 `yaml:"Param"`
	Life  float64 `yaml:"Life"`
}

type Fish struct {
	Id          int32    `yaml:"Id"`
	Name        string   `yaml:"Name"`
	BroadCast   bool     `yaml:"BroadCast"`
	Probability float64  `yaml:"Probability"`
	VisualId    int32    `yaml:"VisualId"`
	Speed       float64  `yaml:"Speed"`
	BoundBox    int32    `yaml:"BoundBox"`
	ShowBingo   bool     `yaml:"ShowBingo"`
	Particle    string   `yaml:"Particle"`
	ShakeScree  bool     `yaml:"ShakeScree"`
	LockLevel   int32    `yaml:"LockLevel"`
	Effects     []Effect `yaml:"Effects"`
	Buffers     []Buffer `yaml:"Buffers"`
}

func LoadSpecialFish(fileName string) bool {
	var config struct {
		King    []SpecialFish `yaml:"King"`
		SanYuan []SpecialFish `yaml:"SanYuan"`
		SiXi    []SpecialFish `yaml:"SiXi"`
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

	for i := 0; i < len(config.King); i++ {
		v := &config.King[i]
		KingFishMap[v.Id] = v
	}
	for i := 0; i < len(config.SanYuan); i++ {
		v := &config.SanYuan[i]
		SanYuanFishMap[v.Id] = v
	}
	for i := 0; i < len(config.SiXi); i++ {
		v := &config.SiXi[i]
		SiXiFishMap[v.Id] = v
	}

	return true
}

func LoadBoundingBox(fileName string)bool{
	var config struct {
		BBX []BBX `yaml:"BBX"`
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

	for i := 0; i < len(config.BBX); i++ {
		v := &config.BBX[i]
		BBXMap[v.Id] = v
	}

	return true
}

func LoadFish(fileName string)bool {
	var config struct {
		Fish []Fish `yaml:"Fish"`
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

	for i := 0; i < len(config.Fish); i++ {
		v := &config.Fish[i]
		FishMap[v.Id] = v
	}

	return true
}