package internal

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	Cannons   []CannonPos
	EffectPos      = CannonPos{}
	LockInfo        = CannonLock{}
	JettonPos      = Point{}
)

// 筹码
type Jetton struct {
	Point `yaml:",inline"`
	Max   int32 `yaml:"Max"`
}

type CannonPart struct {
	Point      `yaml:",inline"`
	ResName    string  `yaml:"ResName"`
	ResType    int32   `yaml:"ResType"`
	Type       int32   `yaml:"Type"`
	FireOffset int32   `yaml:"FireOffset"`
	RoateSpeed float64 `yaml:"RoateSpeed"`
}

// 锁定
type CannonLock struct {
	Point `yaml:",inline"`
	Icon  string `yaml:"Icon"`
	Line  string `yaml:"Line"`
	Flag  string `yaml:"Flag"`
}

type CannonIon struct {
	Point   `yaml:",inline"`
	IonFlag string `yaml:"IonFlag"`
}

// 炮台位置
type CannonPos struct {
	Id        int32 `yaml:"Id"`
	MovePoint `yaml:",inline"`
}

// 子弹
type CannonBullet struct {
	ResName string  `yaml:"ResName"`
	ResType int32   `yaml:"ResType"`
	Scale   float64 `yaml:"Scale"`
}

// 鱼网
type CannonNet struct {
	CannonBullet `yaml:",inline"`
	Point        `yaml:",inline"`
}

type CannonSet struct {
	Id     int32        `yaml:"Id"`
	Cannon CannonPart   `yaml:"Cannon"`
	Bullet CannonBullet `yaml:"Bullet"`
	Net    CannonNet    `yaml:"Net"`
}

type CannonSetS struct {
	Id         int32       `yaml:"Id"`
	Normal     int32       `yaml:"Normal"`
	Ion        int32       `yaml:"Ion"`
	Double     int32       `yaml:"Double"`
	Rebound    bool        `yaml:"Rebound"`
	CannonList []CannonSet `yaml:"CannonList"`
}

func LoadCannon(fileName string) bool {
	var config struct {
		Cannon    []CannonPos  `yaml:"CannonPos"`
		Effect    CannonPos    `yaml:"CannonEffect"`
		Jetton    Jetton       `yaml:"Jetton"`
		Lock      CannonLock   `yaml:"Lock"`
		CannonSet []CannonSetS `yaml:"CannonSet"`
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
	Cannons = config.Cannon
	return true
}
