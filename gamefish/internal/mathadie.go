package internal

import (
	"math"
)

func Factorial(number int) int {
	factorial := 1
	temp := number
	for i := 0; i < number; i++ {
		factorial *= temp
		temp--
	}
	return factorial
}

func Combination(count int, r int) int {
	return Factorial(count) / (Factorial(r) * Factorial(count-r))
}

func CalcDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func CalcAngle(x1, y1, x2, y2 float64) float64 {
	fDistance := CalcDistance(x1, y1, x2, y2)
	if fDistance == 0 {
		return 0
	}

	cosValue := (x1 - x2) / fDistance
	angle := math.Acos(cosValue)
	if y1 < y2 {
		angle = 2*math.Pi - angle
	}
	angle += math.Pi / 2
	return angle
}

func BuildLinear(initX, initY []float64, initCount int,  fDistance float64) (points []Point) {
	points = make([]Point, 0, initCount)

	if initCount < 2 {
		return
	}

	if fDistance <= 0 {
		return
	}

	disTotal := CalcDistance(initX[initCount-1], initY[initCount-1], initX[0], initY[0])
	if disTotal <= 0.0 {
		return
	}

	cosValue := math.Abs(initY[initCount-1]-initY[0]) / disTotal
	angle := math.Acos(cosValue)

	point := Point{
		X: initX[0],
		Y: initY[0],
	}
	points = append(points, point)

	tfDis := float64(0)
	for tfDis < disTotal {
		size := float64(len(points))
		if initX[initCount-1] < initX[0] {
			point.X = initX[0] - math.Sin(angle)*(fDistance*size)
		} else {
			point.X = initX[0] + math.Sin(angle)*(fDistance*size)
		}

		if initY[initCount-1] < initY[0] {
			point.Y = initY[0] - math.Cos(angle)*(fDistance*size)
		} else {
			point.Y = initY[0] + math.Cos(angle)*(fDistance*size)
		}
		points = append(points, point)
		tfDis = CalcDistance(point.X, point.Y, initX[0], initY[0])
	}

	tPoint := &points[len(points)-1]
	tPoint.X = initX[initCount-1]
	tPoint.Y = initY[initCount-1]
	return
}

func BuildLinear2(initX, initY []float64, initCount int32, fDistance float64) (points []MovePoint) {
	points = make([]MovePoint, 0, initCount)
	if initCount < 2 {
		return
	}

	if fDistance <= 0 {
		return
	}

	disTotal := CalcDistance(initX[0], initY[0], initX[initCount-1], initY[initCount-1])
	if disTotal <= 0 {
		return
	}

	tAngle := CalcAngle(initX[initCount-1], initY[initCount-1], initX[0], initY[0]) - (math.Pi / 2)
	point := MovePoint{
		Point: Point{
			X: initX[0],
			Y: initY[0],
		},
		Direction: tAngle,
	}
	points = append(points, point)
	tfDis := float64(0)

	for tfDis < disTotal {
		size := float64(len(points))

		point.X = initX[0] + math.Cos(tAngle)*(fDistance*size)
		point.Y = initY[0] + math.Sin(tAngle)*(fDistance*size)
		point.Direction = tAngle

		points = append(points, point)
		tfDis = CalcDistance(point.X, point.Y, initX[0], initY[0])
	}

	tPoint := &points[len(points)-1]
	tPoint.X = initX[initCount-1]
	tPoint.X = initY[initCount-1]
	return
}

func BuildBezier(initX, initY []float64, initCount int32, fDistance float64) (points []MovePoint) {
	if initCount < 3 {
		return
	}
	index := int32(0)
	tPos0 := MovePoint{}
	t := float64(0)
	count := initCount - 1
	tfDis := fDistance
	tPos := MovePoint{}

	points = make([]MovePoint, 0, int(1.0/0.00001))
	for t < 1 {
		tPos.X = 0
		tPos.Y = 0
		index = 0
		for index <= count {
			tempValue := math.Pow(t, float64(index)) * math.Pow(1-t, float64(count-index)) * float64(Combination(int(count), int(index)))
			tPos.X += initX[index] * tempValue
			tPos.Y += initY[index] * tempValue
			index++
		}

		fSpace := float64(0)
		if len(points) > 0 {
			backPos := &points[len(points)-1]
			fSpace = CalcDistance(backPos.X, backPos.Y, tPos.X, tPos.Y)
		}

		if fSpace >= tfDis || len(points) == 0 {
			if len(points) > 0 {
				temp_dis := CalcDistance(tPos.X, tPos.Y, tPos0.X, tPos0.Y)
				if temp_dis != 0 {
					tempValue := (tPos.X - tPos0.X) / temp_dis
					if tPos.Y-tPos0.Y >= 0 {
						tPos.Direction = math.Acos(tempValue)
					} else {
						tPos.Direction = -math.Acos(tempValue)
					}
				} else {
					tPos.Direction = 1
				}
			} else {
				tPos.Direction = 1.
			}
			points = append(points, tPos)
			tPos0.X = tPos.X
			tPos0.Y = tPos.Y
		}
		t += 0.00001
	}
	return
}

func BuildCircle( centerX,  centerY,  radius float64,  fishCount int32)(points []MovePoint) {
	if fishCount <= 0 || radius == 0 {
		return
	}
	cell_radian := 2 * math.Pi / float64(fishCount)
	for i := int32(0); i < fishCount; i++ {
		pp := MovePoint{
			Point: Point{
				X: centerX + radius*math.Cos(float64(i)*cell_radian),
				Y: centerY + radius*math.Sin(float64(i)*cell_radian),
			},
			Direction: cell_radian,
		}
		points = append(points, pp)
	}
	return
}

func BuildCirclePath( centerX,  centerY,  radius float64,  begin float64,  fAngle float64,  nStep int32,  fAdd float64)(points []MovePoint) {
	if fAngle == 0 || radius == 0 {
		return
	}
	if nStep < 1 {
		nStep = 1
	}
	nCir := float64(int(2 * math.Pi * radius / float64(nStep)))
	nCount := int(nCir * math.Abs(fAngle) / (2 * math.Pi))
	cell_radian := 2 * math.Pi / nCir * fAngle / math.Abs(fAngle)
	pLast := Point{}
	for i := 0; i < nCount; i++ {
		pp := MovePoint{
			Point: Point{
				X: centerX + radius*math.Cos(begin+float64(i)*cell_radian),
				Y: centerY + radius*math.Sin(begin+float64(i)*cell_radian),
			},
		}
		if i == 0 {
			pp.Direction = begin + float64(i)*cell_radian + (math.Pi / 2)
		} else {
			pp.Direction = CalcAngle(pLast.X, pLast.Y, pp.X, pp.Y) + (math.Pi / 2)
		}
		pLast = pp.Point
		if fAdd != 0 {
			radius += fAdd
		}
		points = append(points, pp)
	}
	return
}

func GetRotationPosByOffest( xPos,  yPos,  xOffest,  yOffest,  dir,  fHScale,  fVScale float64)(Point) {
	pt := Point{}
	r := math.Sqrt(xOffest*xOffest + yOffest*yOffest)
	fd := CalcAngle(0, 0, xOffest, yOffest) - (math.Pi / 2) + dir
	pt.X = (xPos - r*math.Cos(fd)) * fHScale
	pt.Y = (yPos - r*math.Sin(fd)) * fVScale
	return pt
}