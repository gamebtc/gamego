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

// 游戏默认的窗口大小
var (
	DefaultWidth  = float64(2160)
	DefaultHeight = float64(1350)
)

type SPATH struct {
	Id         int32      `yaml:"Id"`
	Type       int32      `yaml:"Type"`
	PosX       [4]float64 `yaml:"PosX"`
	PosY       [4]float64 `yaml:"PosY"`
	Next       int32      `yaml:"Next"`
	Delay      int32      `yaml:"Delay"`
	PointCount int32      `yaml:"-"`
}

type ShapeBase struct {
	Count    int32   `yaml:"Count"`
	Same     bool    `yaml:"Same"`
	Path     int32   `yaml:"Path"`
	Speed    float64 `yaml:"Speed"`
	Interval float64 `yaml:"Interval"`
	FishList []int32 `yaml:"FishList"`
	Weight   []int32 `yaml:"Weight"`
}

type ShapeLine struct {
	X          [2]float64 `yaml:"X"`
	Y          [2]float64 `yaml:"Y"`
	Count      int32      `yaml:"Count"`
	PriceCount int32      `yaml:"PriceCount"`
	Path       int32      `yaml:"Path"`
	Same       bool       `yaml:"Same"`
	Speed      float64    `yaml:"Speed"`
	Interval   float64    `yaml:"Interval"`
	FishType   []int32    `yaml:"FishType"`
	Weight     []int32    `yaml:"Weight"`
}

type ShapeCircle struct {
	X          float64 `yaml:"X"`
	Y          float64 `yaml:"Y"`
	R          float64 `yaml:"R"`
	Count      int32   `yaml:"Count"`
	PriceCount int32   `yaml:"PriceCount"`
	Same       bool    `yaml:"Same"`
	Path       int32   `yaml:"Path"`
	Speed      float64 `yaml:"Speed"`
	Interval   float64 `yaml:"Interval"`
	FishType   []int32 `yaml:"FishType"`
	Weight     []int32 `yaml:"Weight"`
}

type ShapePoint struct {
	X        float64 `yaml:"X"`
	Y        float64 `yaml:"Y"`
	Count    int32   `yaml:"Count"`
	Path     int32   `yaml:"Path"`
	Speed    float64 `yaml:"Speed"`
	Same     bool    `yaml:"Same"`
	Interval float64 `yaml:"Interval"`
	FishType []int32 `yaml:"FishType"`
	Weight   []int32 `yaml:"Weight"`
}

type TroopData struct {
	Id         int32         `yaml:"Id"`
	Describes  []string      `yaml:"Describes"`
	LineData   []ShapeLine   `yaml:"LineData"`
	CircleData []ShapeCircle `yaml:"CircleData"`
	PointData  []ShapePoint  `yaml:"PointData"`
}

type Troop struct {
	Id        int32        `yaml:"Id"`
	Describes []string     `yaml:"Describes"`
	Step      []int32      `yaml:"Step"`
	Shape     []ShapePoint `yaml:"Shape"`
}

var (
	LoadPath         = false
	NormalPaths      = make([]SPATH, 0, 64)
	NormalPathVector = make([]MovePoints, 0, 64*2*2*2*2)

	TroopPath    = make(map[int32]*SPATH)
	TroopMap     = make(map[int32]*Troop)
	TroopDataMap = make(map[int32]*TroopData)
	TroopPathMap = make(map[int32]MovePoints)
)

func GetNormalPath(id int32) *SPATH {
	if len(NormalPaths) <= 0 || id < 0 {
		return nil
	}
	return &NormalPaths[id%int32(len(NormalPaths))]
}

func GetPathData(id int32, troop bool) MovePoints {
	if troop {
		if r, ok := TroopPathMap[id]; ok {
			return r
		} else {
			return nil
		}
	} else {
		return NormalPathVector[id%int32(len(NormalPathVector))]
	}
}

func GetTroopPath(id int32) *SPATH {
	if r, ok := TroopPath[id]; ok {
		return r
	}
	return nil
}

func GetTroop(id int32) *Troop {
	if r, ok := TroopMap[id]; ok {
		return r
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
			tt.Count = sl.PriceCount
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
			tt.Count = sc.PriceCount
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

func SetPointCount(p []SPATH) {
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
		Path  []SPATH     `yaml:"Path"`
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

	for i := 0; i < len(config.Troop); i++ {
		td := &config.Troop[i]
		TroopDataMap[td.Id] = td
		trp := CreatTroopByData(td)
		TroopMap[td.Id] = trp
	}

	for i := 0; i < len(config.Path); i++ {
		pd := &config.Path[i]
		TroopPath[pd.Id] = pd
	}

	size := len(TroopPath)
	exclude := make([]int32, 0, size)
	for k, sph := range TroopPath {
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
			if _, ok := TroopPath[nxt]; ok == false {
				break
			}
			exclude = append(exclude, nxt)
			nxt = TroopPath[nxt].Next
		}
		path := CreatePathByData(sph, false, false, false, false, true)
		TroopPathMap[k] = path
	}
	return true
}

func LoadNormalPath(fileName string) bool {
	// TODO:json to NormalPaths
	var config struct {
		Path []SPATH `yaml:"Path"`
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

	NormalPaths = config.Path
	size := int32(len(NormalPaths))
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
		sph := &NormalPaths[i]
		nxt := sph.Next
		for nxt > 0 && nxt < size {
			exclude = append(exclude, nxt)
			nxt = NormalPaths[nxt].Next
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

func CreatePathByData(sp *SPATH, xMirror, yMirror, xyMirror, Not, troop bool) (out MovePoints) {
	out = make(MovePoints, 0, 1000)
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
			x[n] *= DefaultWidth
			y[n] *= DefaultHeight
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
			if s, ok := TroopPath[nxt]; ok {
				sp = s
			} else {
				break
			}
		} else {
			if sp.Next > 0 && sp.Next < int32(len(NormalPaths)) {
				sp = &NormalPaths[sp.Next]
			} else {
				break
			}
		}
	}
	return
}
