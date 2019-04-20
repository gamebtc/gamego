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

func (point *Point) Mul(mul float64) Point {
	return Point{
		X: point.X * mul,
		Y: point.Y * mul,
	}
}

type Points = []Point

type MovePoint struct {
	//Point     `yaml:",inline"`
	X float64 `yaml:"X"`
	Y float64 `yaml:"Y"`
	Direction float64 `yaml:"Direction"`
}

type MovePoints = []MovePoint
