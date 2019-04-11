package internal

type Point struct {
	X float64 `yaml:"X"`
	Y float64 `yaml:"Y"`
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
	Point     `yaml:",inline"`
	Direction float64 `yaml:"Direction"`
}

type MovePoints = []MovePoint
