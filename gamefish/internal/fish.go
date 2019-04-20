package internal

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

const (
	SpecialFishType_Normal      = 0
	SpecialFishType_King        = 1
	SpecialFishType_KingAndQuan = 2
	SpecialFishType_Sanyuan     = 3
	SpecialFishType_Sixi        = 4
	SpecialFishType_Max         = 5
)

var (
	kingFishMap        = make(map[int32]*SpecialFishTemplate, 20) // 鱼王
	sanYuanFishMap     = make(map[int32]*SpecialFishTemplate, 20) // 三元
	siXiFishMap        = make(map[int32]*SpecialFishTemplate, 20) // 四喜
	bbxMap             = make(map[int32]*BBX, 100)                //
	fishTemplateMap    = make(map[int32]*FishTemplate, 100)
	bulletTemplateList []BulletTemplate
)

// 鱼王/三元/四喜
type SpecialFishTemplate struct {
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
	Id    int32         `yaml:"Id"`
	Boxes []BoundingBox `yaml:"BB"`
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

type FishTemplate struct {
	Id          int32         `yaml:"Id"`
	Name        string        `yaml:"Name"`
	BroadCast   bool          `yaml:"BroadCast"`
	Probability float64       `yaml:"Probability"`
	VisualId    int32         `yaml:"VisualId"`
	Speed       float64       `yaml:"Speed"`
	BoundBox    int32         `yaml:"BoundBox"`
	ShowBingo   bool          `yaml:"ShowBingo"`
	Particle    string        `yaml:"Particle"`
	ShakeScree  bool          `yaml:"ShakeScree"`
	LockLevel   int32         `yaml:"LockLevel"`
	Effects     []Effect      `yaml:"Effects"`
	Buffers     []Buffer      `yaml:"Buffers"`
	Boxes       []BoundingBox `yaml:"-"`
}

// 子弹
type BulletTemplate struct {
	Multiple    int32 `yaml:"Multiple"`    // 倍率
	Speed       int32 `yaml:"Speed"`       // 速度
	MaxCatch    int32 `yaml:"MaxCatch"`    // 最大捕鱼数量
	CatchRadius int32 `yaml:"CatchRadius"` // 捕获半径
	CannonType  int32 `yaml:"CannonType"`  // 炮管类型
}

func LoadConfig() {
	LoadBullet("bullet.yaml")
	LoadCannon("cannon.yaml")
	LoadScene("scene.yaml")
	LoadBoundingBox("bbox.yaml")
	LoadFish("fish.yaml")
	LoadSpecialFish("special.yaml")
	LoadTroop("troop.yaml")
	LoadNormalPath("path.yaml")
}

func LoadSpecialFish(fileName string) bool {
	var config struct {
		King    []SpecialFishTemplate `yaml:"King"`
		SanYuan []SpecialFishTemplate `yaml:"SanYuan"`
		SiXi    []SpecialFishTemplate `yaml:"SiXi"`
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
		kingFishMap[v.Id] = v
	}
	for i := 0; i < len(config.SanYuan); i++ {
		v := &config.SanYuan[i]
		sanYuanFishMap[v.Id] = v
	}
	for i := 0; i < len(config.SiXi); i++ {
		v := &config.SiXi[i]
		siXiFishMap[v.Id] = v
	}

	return true
}

func LoadBoundingBox(fileName string) bool {
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
		bbxMap[v.Id] = v
	}

	return true
}

func LoadFish(fileName string) bool {
	var config struct {
		Fish []FishTemplate `yaml:"Fish"`
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
		if bbList, ok := bbxMap[v.BoundBox]; ok {
			v.Boxes = bbList.Boxes
		}
		fishTemplateMap[v.Id] = v
	}
	return true
}

func LoadBullet(fileName string) bool {
	var config struct {
		Bullet []BulletTemplate `yaml:"Bullet"`
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
	bulletTemplateList = config.Bullet
	return true
}
