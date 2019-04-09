package internal

type Point struct {
	X float64
	Y float64
}

func (point *Point) Offset(x, y float64) {
	point.X += x
	point.Y += y
}

func (point *Point) SetPoint(x, y float64) {
	point.X = x
	point.Y = y
}

func (point *Point) Add(a Point) Point {
	return Point{
		X: point.X + a.X,
		Y: point.Y + a.Y,
	}
}

func (point *Point) Dec(a Point) Point {
	return Point{
		X: point.X - a.X,
		Y: point.Y - a.Y,
	}
}

func (point *Point) Mul(multip float64) Point {
	return Point{
		X: point.X * multip,
		Y: point.Y * multip,
	}
}

type MovePoint struct {
	Point
	Direction float64
}

type MovePoints = []MovePoint