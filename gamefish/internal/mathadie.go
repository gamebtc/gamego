package internal

import (
	"math"
)

func Factorial(n int) int {
	val := 1
	for i := 1; i <= n; i++ {
		val *= i
	}
	return val
}

func Combination(count int, r int) int {
	return Factorial(count) / (Factorial(r) * Factorial(count-r))
}

func CalcDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func CalcAngle(x1, y1, x2, y2 float64) float64 {
	distance := CalcDistance(x1, y1, x2, y2)
	if distance == 0 {
		return 0
	}

	cosValue := (x1 - x2) / distance
	angle := math.Acos(cosValue)
	if y1 < y2 {
		angle = 2*math.Pi - angle
	}
	angle += math.Pi / 2
	return angle
}

func BuildLinear(initX, initY []float64, initCount int, distance float64) (points Points) {
	points = make(Points, 0, initCount)

	if initCount < 2 {
		return
	}

	if distance <= 0 {
		return
	}

	totalDist := CalcDistance(initX[initCount-1], initY[initCount-1], initX[0], initY[0])
	if totalDist <= 0.0 {
		return
	}

	cosValue := math.Abs(initY[initCount-1]-initY[0]) / totalDist
	angle := math.Acos(cosValue)

	point := Point{
		X: initX[0],
		Y: initY[0],
	}
	points = append(points, point)

	dis := float64(0)
	for dis < totalDist {
		dis := float64(len(points)) * distance
		if initX[initCount-1] < initX[0] {
			point.X = initX[0] - math.Sin(angle)*dis
		} else {
			point.X = initX[0] + math.Sin(angle)*dis
		}

		if initY[initCount-1] < initY[0] {
			point.Y = initY[0] - math.Cos(angle)*dis
		} else {
			point.Y = initY[0] + math.Cos(angle)*dis
		}
		points = append(points, point)
		dis = CalcDistance(point.X, point.Y, initX[0], initY[0])
	}
	points[len(points)-1].X = initX[initCount-1]
	points[len(points)-1].Y = initY[initCount-1]
	return
}

func BuildLinear2(initX, initY []float64, initCount int32, distance float64) (points MovePoints) {
	points = make(MovePoints, 0, initCount)
	if initCount < 2 {
		return
	}

	if distance <= 0 {
		return
	}

	totalDist := CalcDistance(initX[0], initY[0], initX[initCount-1], initY[initCount-1])
	if totalDist <= 0 {
		return
	}

	angle := CalcAngle(initX[initCount-1], initY[initCount-1], initX[0], initY[0]) - (math.Pi / 2)
	point := MovePoint{
		X:         initX[0],
		Y:         initY[0],
		Direction: angle,
	}
	points = append(points, point)
	dis := float64(0)
	for dis < totalDist {
		dis := float64(len(points)) * distance
		point.X = initX[0] + math.Cos(angle)*dis
		point.Y = initY[0] + math.Sin(angle)*dis
		//point.Direction = Direction

		points = append(points, point)
		dis = CalcDistance(point.X, point.Y, initX[0], initY[0])
	}
	points[len(points)-1].X = initX[initCount-1]
	points[len(points)-1].Y = initY[initCount-1]
	return
}

func BuildBezier(initX, initY []float64, initCount int32, distance float64) MovePoints {
	if initCount < 3 {
		return make(MovePoints, 0, 0)
	}
	pos0 := Point{}
	t := float64(0)
	count := int(initCount - 1)
	points := make(MovePoints, 0, int(1.0/0.00001))
	for t < 1 {
		x, y := float64(0), float64(0)
		for i := 0; i <= count; i++ {
			tempValue := math.Pow(t, float64(i)) * math.Pow(1-t, float64(count-i)) * float64(Combination(count, i))
			x += initX[i] * tempValue
			y += initY[i] * tempValue
		}

		space := float64(0)
		if len(points) > 0 {
			backPos := &points[len(points)-1]
			space = CalcDistance(backPos.X, backPos.Y, x, y)
		}

		if space >= distance || len(points) == 0 {
			var direction float64
			if len(points) > 0 {
				if dist := CalcDistance(x, y, pos0.X, pos0.Y); dist != 0 {
					tempValue := (x - pos0.X) / dist
					if y-pos0.Y >= 0 {
						direction = math.Acos(tempValue)
					} else {
						direction = -math.Acos(tempValue)
					}
				} else {
					direction = 1
				}
			} else {
				direction = 1
			}
			points = append(points, MovePoint{X: x, Y: y, Direction: direction})
			pos0.X = x
			pos0.Y = y
		}
		t += 0.00001
	}
	return points
}

func BuildCircle(centerX, centerY, radius float64, fishCount int32) MovePoints {
	if fishCount <= 0 || radius == 0 {
		return make(MovePoints, 0, 0)
	}
	cellRadian := 2 * math.Pi / float64(fishCount)
	points := make(MovePoints, 0, fishCount)
	for i := int32(0); i < fishCount; i++ {
		pp := MovePoint{
			X:         centerX + radius*math.Cos(float64(i)*cellRadian),
			Y:         centerY + radius*math.Sin(float64(i)*cellRadian),
			Direction: cellRadian,
		}
		points = append(points, pp)
	}
	return points
}

func BuildCirclePath(centerX, centerY, radius float64, begin float64, angle float64, step int32, add float64) MovePoints {
	if angle == 0 || radius == 0 {
		return make(MovePoints, 0, 0)
	}
	if step < 1 {
		step = 1
	}
	cir := float64(int(2 * math.Pi * radius / float64(step)))
	count := int(cir * math.Abs(angle) / (2 * math.Pi))
	cellRadian := 2 * math.Pi / cir * angle / math.Abs(angle)
	lastX, lastY := float64(0), float64(0)
	points := make(MovePoints, 0, count)
	for i := 0; i < count; i++ {
		pp := MovePoint{
			X: centerX + radius*math.Cos(begin+float64(i)*cellRadian),
			Y: centerY + radius*math.Sin(begin+float64(i)*cellRadian),
		}
		if i == 0 {
			pp.Direction = begin + float64(i)*cellRadian + (math.Pi / 2)
		} else {
			pp.Direction = CalcAngle(lastX, lastY, pp.X, pp.Y) + (math.Pi / 2)
		}
		lastX = pp.X
		lastY = pp.Y
		if add != 0 {
			radius += add
		}
		points = append(points, pp)
	}
	return points
}

func GetRotationPosByOffset(xPos, yPos, xOffset, yOffset, dir, hScale, vScale float64) Point {
	r := math.Sqrt(xOffset*xOffset + yOffset*yOffset)
	fd := CalcAngle(0, 0, xOffset, yOffset) - (math.Pi / 2) + dir
	return Point{
		X: (xPos - r*math.Cos(fd)) * hScale,
		Y: (yPos - r*math.Sin(fd)) * vScale,
	}
}
