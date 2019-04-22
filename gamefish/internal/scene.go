package internal

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	SceneConfigs []SceneConfig
	SystemConf SystemConfig
)

type SystemConfig struct {
	SnakeHead       int32     `yaml:"SnakeHead"`
	SnakeTail       int32     `yaml:"SnakeTail"`
	ScreenWidth     float64   `yaml:"ScreenWidth"`     // 屏幕宽度
	ScreenHeight    float64   `yaml:"ScreenHeight"`    // 屏幕高度
	RobotProbMul    float64   `yaml:"RobotProbMul"`    // 机器人捕鱼概率
	BulletSpeed     float64   `yaml:"BulletSpeed"`     // 子弹速度
	BulletInterval  float64   `yaml:"BulletInterval"`  // 开炮时间最小间隔(毫秒)
	MaxBulletCount  int32     `yaml:"MaxBulletCount"`  // 每个人最大发弹数
	SwitchSceneTime float64   `yaml:"SwitchSceneTime"` // 切换场景的时间5秒
	NoticeLevel     int64     `yaml:"NoticeLevel"`     // 消息广播等级
	FirstFire       FirstFire `yaml:"FirstFire"`
}

type DistributeFishConfig struct {
	FishWeight  `yaml:",inline"`
	Time        float64 `yaml:"Time"`
	MinCount    int32   `yaml:"MinCount"`
	MaxCount    int32   `yaml:"MaxCount"`
	RefreshType int32   `yaml:"RefreshType"`
	OffsetX     float64 `yaml:"OffsetX"`
	OffsetY     float64 `yaml:"OffsetY"`
	OffsetTime  float64 `yaml:"OffsetTime"`
}

type TroopConfig struct {
	BeginTime float64 `yaml:"BeginTime"`
	EndTime   float64 `yaml:"EndTime"`
	TroopId   int32   `yaml:"TroopId"`
}

type SceneConfig struct {
	Id             int32                  `yaml:"Id"`
	Next           int32                  `yaml:"Next"`
	Image          string                 `yaml:"Image"`
	Time           float64                `yaml:"Time"`
	Troop          TroopConfig            `yaml:"Troop"`
	DistributeFish []DistributeFishConfig `yaml:"DistributeFish"`
}

type FirstFire struct {
	FishWeight `yaml:",inline"`
	Level      int32 `yaml:"Level"`
	CreatCount int32 `yaml:"CreatCount"`
}


func (t *SceneConfig) fixWeight() {
	for i := 0; i < len(t.DistributeFish); i++ {
		t.DistributeFish[i].fixWeight()
	}
}

func (t *SceneConfig) findTroop(time float64) *TroopConfig {
	if time >= t.Troop.BeginTime && time < t.Troop.EndTime {
		return &t.Troop
	}
	return nil
}

func LoadScene(fileName string) bool {
	var config struct {
		Scene []SceneConfig `yaml:"Scene"`
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("scene config file not exists:%v", err)
		return false
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("scene config file error:%v", err)
		return false
	}

	for i := 0; i < len(config.Scene); i++ {
		config.Scene[i].fixWeight()
	}
	SceneConfigs = config.Scene

	return true
}

func LoadSystem(fileName string) bool {
	var config struct {
		System SystemConfig `yaml:"System"`
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("system config file not exists:%v", err)
		return false
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("system config file error:%v", err)
		return false
	}
	SystemConf = config.System
	SystemConf.FirstFire.fixWeight()
	return true
}

func findSceneConfig(id int32)*SceneConfig {
	for i := 0; i < len(SceneConfigs); i++ {
		if t := &SceneConfigs[i]; t.Id == id {
			return t
		}
	}
	return nil
}