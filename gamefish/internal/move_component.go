package internal

import (
	"math"
)

const (
	GAME_FPS = 5
)

func (fish *Fish) moveByPath(second float64) {
	if fish.BeginMove == false && fish.Elapse > 0 {
		fish.BeginMove = true
	}
	fish.Elapse += second * fish.Speed

	pathSize := int32(len(fish.Path))
	time := math.Min(1.0, fish.Elapse/float64(pathSize))
	fIndex := time * float64(pathSize)
	index := int32(fIndex)
	diff := fIndex - float64(index)

	if index >= pathSize {
		index = pathSize - 1
	} else if index < 0 || diff < 0 {
		index = 0
		diff = 0
	}

	mp := MovePoint{}
	if index < pathSize-1 {
		p1 := fish.Path[index]
		p2 := fish.Path[index+1]
		mp.X = p1.X*(1-diff) + (p2.X * diff)
		mp.Y = p1.Y*(1-diff) + (p2.Y * diff)
		mp.Direction = p1.Direction*(1-diff) + (p2.Direction * diff)

		if math.Abs(p1.Direction-p2.Direction) > math.Pi {
			mp.Direction = p1.Direction
		}
	} else {
		mp = fish.Path[index]
		fish.EndPath = true
	}

	fish.X = mp.X + fish.Offset.X
	fish.Y = mp.Y + fish.Offset.Y
	fish.Direction = mp.Direction
}
func (fish *Fish) initAngle() {
	fish.EndPath = false
	a := fish.Direction
	fish.angle = a
	a -= math.Pi / 2
	fish.dx = math.Cos(a)
	fish.dy = math.Sin(a)
}

func (fish *Fish) moveByDirection(second float64) {
	if fish.BeginMove == false {
		fish.BeginMove = true
	}

	if fish.TargetId != 0 {
		target := fish.table.fishes[fish.TargetId]
		if target!=nil && target.State < ObjectState_Dead && target.InSide {
			if CalcDistance(target.X, target.Y, fish.X, fish.Y) > 10 {
				fish.Direction = CalcAngle(target.X, target.Y, fish.X, fish.Y)
				fish.initAngle()
			} else {
				fish.X = target.X
				fish.Y = target.Y
				fish.Direction = target.Direction
				return
			}
		} else {
			fish.TargetId = 0
		}
	}
	x := fish.X + (fish.Speed * fish.dx * second)
	y := fish.Y + (fish.Speed * fish.dy * second)
	if x < 0 || x > SystemConf.ScreenWidth || y < 0 || y >  SystemConf.ScreenHeight {
		fish.EndPath = true
	}

	fish.Direction = fish.angle - (math.Pi / 2)
	fish.X = x
	fish.Y = y
}
