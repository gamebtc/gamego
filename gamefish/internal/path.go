package internal

import (
	"io/ioutil"
	"math"

	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

// NormalPathType
const (
	NPT_LINE   = 0
	NPT_BEZIER = 1
	NPT_CIRCLE = 2
)

var (
	LoadPath         = false
	normalPaths      = make([]Path, 0, 64)
	NormalPathVector = make([]MovePoints, 0, 64*2*2*2*2)
	troopPaths       = make(map[int32]*Path)
	troopPoints      = make(map[int32]MovePoints)
	troops           []Troop
	troopDatas       []TroopData
)

type Path struct {
	Id         int32      `yaml:"Id"`
	Type       int32      `yaml:"Type"`
	PosX       [4]float64 `yaml:"PosX"`
	PosY       [4]float64 `yaml:"PosY"`
	Next       int32      `yaml:"Next"`
	Delay      int32      `yaml:"Delay"`
	PointCount int32      `yaml:"-"`
}

type FishWeight struct {
	FishType []int32 `yaml:"FishType"`
	Weight   []int32 `yaml:"Weight"`
	fix      bool    `yaml:"-"`
}

func (fw *FishWeight) fixWeight() {
	if fw.fix {
		return
	}
	fw.fix = true
	l := len(fw.FishType)
	if l < len(fw.Weight) {
		l = len(fw.Weight)
	}
	if l == 0 {
		fw.FishType = nil
		fw.Weight = nil
		return
	}
	fw.FishType = fw.FishType[0:l]
	fw.Weight = fw.Weight[0:l]
	for i := 1; i < len(fw.Weight); i++ {
		fw.Weight[i] += fw.Weight[i-1]
	}
}

func (fw *FishWeight) RandFish() int32 {
	if fw.fix == false {
		fw.fixWeight()
	}
	l := int32(len(fw.Weight))
	if l > 0 {
		r := gameRand.Int31n(l)
		for i, w := range fw.Weight {
			if r < w {
				return fw.FishType[i]
			}
		}
	}
	return 0
}

type ShapePoint struct {
	FishWeight `yaml:",inline"`
	Count      int32   `yaml:"Count"`
	Path       int32   `yaml:"Path"`
	Speed      float64 `yaml:"Speed"`
	Same       bool    `yaml:"Same"`
	Interval   float64 `yaml:"Interval"`
	X          float64 `yaml:"X"`
	Y          float64 `yaml:"Y"`
}

type ShapeLine struct {
	FishWeight `yaml:",inline"`
	Count      int32      `yaml:"Count"`
	Path       int32      `yaml:"Path"`
	Speed      float64    `yaml:"Speed"`
	Same       bool       `yaml:"Same"`
	Interval   float64    `yaml:"Interval"`
	X          [2]float64 `yaml:"X"`
	Y          [2]float64 `yaml:"Y"`
	CreatCount int32      `yaml:"CreatCount"`
}

type ShapeCircle struct {
	ShapePoint `yaml:",inline"`
	R          float64 `yaml:"R"`
	CreatCount int32   `yaml:"CreatCount"`
}

type TroopData struct {
	Id         int32         `yaml:"Id"`
	Describes  []string      `yaml:"Describes"`
	LineData   []ShapeLine   `yaml:"LineData"`
	CircleData []ShapeCircle `yaml:"CircleData"`
	PointData  []ShapePoint  `yaml:"PointData"`
}

func (t *TroopData) fixWeight() {
	for i := 0; i < len(t.LineData); i++ {
		t.LineData[i].fixWeight()
	}
	for i := 0; i < len(t.CircleData); i++ {
		t.CircleData[i].fixWeight()
	}
	for i := 0; i < len(t.PointData); i++ {
		t.PointData[i].fixWeight()
	}
}

type Troop struct {
	Id        int32        `yaml:"Id"`
	Describes []string     `yaml:"Describes"`
	Step      []int32      `yaml:"Step"`
	Shape     []ShapePoint `yaml:"Shape"`
}

func GetNormalPath(id int32) *Path {
	if len(normalPaths) <= 0 || id < 0 {
		return nil
	}
	return &normalPaths[id%int32(len(normalPaths))]
}

func GetTroopPath(id int32) *Path {
	return troopPaths[id]
}

func GetPathData(id int32, troop bool) MovePoints {
	if troop {
		return troopPoints[id]
	} else {
		return NormalPathVector[id%int32(len(NormalPathVector))]
	}
}

func GetTroop(id int32) *Troop {
	for i := len(troops) - 1; i >= 0; i-- {
		t := &troops[i]
		if t.Id == id {
			return t
		}
	}
	return nil
}

func GetRandNormalPathID() int32 {
	return gameRand.Int31n(int32(len(NormalPathVector)))
}

func CreatTroopByData(td *TroopData) (tp *Troop) {
	tp = &Troop{
		Id: td.Id,
	}
	lineCount := len(td.LineData)
	tp.Describes = td.Describes
	tp.Shape = make([]ShapePoint, 0, lineCount)
	tp.Step = make([]int32, 0, lineCount)
	for _, sl := range td.LineData {
		nc := sl.Count - 1
		if nc <= 0 {
			nc = 1
		}
		traceVectors := BuildLinear2(sl.X[:], sl.Y[:], 2, CalcDistance(sl.X[0], sl.Y[0], sl.X[1], sl.Y[1])/float64(nc))
		nc = int32(len(traceVectors))
		for i := int32(0); i < nc; i++ {
			tt := ShapePoint{}
			tt.X = traceVectors[i].X
			tt.Y = traceVectors[i].Y
			tt.Same = sl.Same
			tt.Count = sl.CreatCount
			tt.Path = sl.Path
			tt.Interval = sl.Interval
			tt.Speed = sl.Speed

			nt := len(sl.FishType)
			if nt > len(sl.Weight) {
				nt = len(sl.Weight)
			}

			for j := 0; j < nt; j++ {
				tt.FishType = append(tt.FishType, sl.FishType[j])
				tt.Weight = append(tt.Weight, sl.Weight[j])
			}
			tp.Shape = append(tp.Shape, tt)
		}
		tp.Step = append(tp.Step, nc)
	}

	//
	for _, sc := range td.CircleData {
		nc := sc.Count
		if nc <= 0 {
			nc = 1
		}
		traceVectors := BuildCircle(sc.X, sc.Y, sc.R, nc)
		nc = int32(len(traceVectors))
		for i := int32(0); i < nc; i++ {
			tt := ShapePoint{}
			tt.X = traceVectors[i].X
			tt.Y = traceVectors[i].Y
			tt.Same = sc.Same
			tt.Count = sc.CreatCount
			tt.Path = sc.Path
			tt.Interval = sc.Interval
			tt.Speed = sc.Speed
			tt.FishType = make([]int32, len(sc.FishType))
			tt.Weight = make([]int32, len(sc.Weight))
			copy(tt.FishType, sc.FishType)
			copy(tt.Weight, sc.Weight)
			tp.Shape = append(tp.Shape, tt)
		}
		tp.Step = append(tp.Step, nc)
	}

	//
	for _, ip := range td.PointData {
		tp.Shape = append(tp.Shape, ip)
		tp.Step = append(tp.Step, 1)
	}
	return
}

func SetPointCount(p []Path) {
	size := int32(len(p))
	for i := int32(0); i < size; i++ {
		path := &p[i]
		if path.Type == NPT_LINE {
			path.PointCount = 2
		} else if path.Type == NPT_BEZIER {
			if path.PosX[3] == 0 && path.PosY[3] == 0 {
				path.PointCount = 3
			} else {
				path.PointCount = 4
			}
		} else {
			path.PointCount = 4
		}
	}
}

func LoadTroop(fileName string) bool {
	// TODO:json to TroopPath
	var config struct {
		Troop []TroopData `yaml:"Troop"`
		Path  []Path      `yaml:"Path"`
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
	SetPointCount(config.Path)

	troopDatas = config.Troop
	troops = make([]Troop,0,len(config.Troop))
	for i := 0; i < len(config.Troop); i++ {
		td := &config.Troop[i]
		//td.Fix
		trp := CreatTroopByData(td)
		troops = append(troops, *trp)
	}


	for i := 0; i < len(config.Path); i++ {
		pd := &config.Path[i]
		troopPaths[pd.Id] = pd
	}

	size := len(troopPaths)
	exclude := make([]int32, 0, size)
	for k, sph := range troopPaths {
		find := false
		for _, tt := range exclude {
			if k == tt {
				find = true
				break
			}
		}
		if find {
			continue
		}

		nxt := sph.Next
		for nxt > 0 {
			if tmpPath := troopPaths[nxt]; tmpPath == nil {
				break
			}else {
				exclude = append(exclude, nxt)
				nxt = tmpPath.Next
			}
		}
		path := CreatePathByData(sph, false, false, false, false, true)
		troopPoints[k] = path
	}
	return true
}

func LoadNormalPath(fileName string) bool {
	// TODO:json to NormalPaths
	var config struct {
		Path []Path `yaml:"Path"`
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
	SetPointCount(config.Path)

	normalPaths = config.Path
	size := int32(len(normalPaths))
	exclude := make([]int32, 0, size)
	for i := int32(0); i < size; i++ {
		find := false
		for _, tt := range exclude {
			if i == tt {
				find = true
				break
			}
		}
		if find {
			continue
		}
		sph := &normalPaths[i]
		nxt := sph.Next
		for nxt > 0 && nxt < size {
			exclude = append(exclude, nxt)
			nxt = normalPaths[nxt].Next
		}
		for x := 0; x < 2; x++ {
			for y := 0; y < 2; y++ {
				for xy := 0; xy < 2; xy++ {
					for not := 0; not < 2; not++ {
						path := CreatePathByData(sph, x == 0, y == 0, xy == 0, not == 0, false)
						NormalPathVector = append(NormalPathVector, path)
					}
				}
			}
		}
	}
	return true
}

func CreatePathByData(sp *Path, xMirror, yMirror, xyMirror, Not, troop bool) MovePoints {
	out := make(MovePoints, 0, 1000)
	for sp != nil {
		x := [4]float64{}
		y := [4]float64{}
		copy(x[:], sp.PosX[:])
		copy(y[:], sp.PosY[:])
		pointCount := sp.PointCount

		if xMirror {
			if sp.Type == NPT_CIRCLE {
				x[0] = 1 - x[0]
				x[2] = math.Pi - x[2]
				y[2] = -y[2]
			} else {
				for n := int32(0); n < pointCount; n++ {
					x[n] = 1 - x[n]
				}
			}
		}
		if yMirror {
			if sp.Type == NPT_CIRCLE {
				y[0] = 1 - y[0]
				x[2] = 2*math.Pi - x[2]
				y[2] = -y[2]
			} else {
				for n := int32(0); n < pointCount; n++ {
					y[n] = 1 - y[n]
				}
			}
		}

		if xyMirror {
			if sp.Type == NPT_CIRCLE {
				t := x[0]
				x[0] = 1.0 - y[0]
				y[0] = 1.0 - t
				x[2] += math.Pi / 2
			} else {
				for n := int32(0); n < pointCount; n++ {
					t := x[n]
					x[n] = y[n]
					y[n] = t
				}
			}
		}
		//取反
		if Not {
			if sp.Type == NPT_CIRCLE {
				x[2] += y[2]
				y[2] = -y[2]
			} else {
				for n := int32(0); n < pointCount/2; n++ {
					t := x[n]
					x[n] = x[pointCount-1-n]
					x[pointCount-1-n] = t

					t = y[n]
					y[n] = y[pointCount-1-n]
					y[pointCount-1-n] = t
				}
			}
		}

		for n := int32(0); n < pointCount; n++ {
			x[n] *= SystemConf.ScreenWidth
			y[n] *= SystemConf.ScreenHeight
			if sp.Type == NPT_CIRCLE {
				break
			}
		}

		var path MovePoints
		if sp.Type == NPT_LINE {
			path = BuildLinear2(x[:], y[:], pointCount, 1.0)
		} else if sp.Type == NPT_BEZIER {
			path = BuildBezier(x[:], y[:], pointCount, 1.0)
		} else if sp.Type == NPT_CIRCLE {
			path = BuildCirclePath(x[0], y[0], x[1], x[2], y[2], 1, y[1])
		}
		out = append(out, path...)

		if sp.Delay > 0 && len(path) > 0 {
			last := out[len(out)-1]
			for i := int32(0); i < sp.Delay; i++ {
				out = append(out, last)
			}
		}

		if troop {
			nxt := sp.Next
			if s := troopPaths[nxt]; s != nil {
				sp = s
			} else {
				break
			}
		} else {
			if sp.Next > 0 && sp.Next < int32(len(normalPaths)) {
				sp = &normalPaths[sp.Next]
			} else {
				break
			}
		}
	}
	return out
}
