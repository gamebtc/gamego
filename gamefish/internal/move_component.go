package internal

import (
	"math"
)

const (
	GAME_FPS = 5
)

type mover interface {
	InitMove(*MoveComponent)
	OnUpdate(float64)
}

type MoveComponent struct {
	MovePoint
	PathId    int32
	Path      MovePoints
	Speed     float64
	Pause     bool
	EndPath   bool
	Offset    Point
	Delay     float64
	BeginMove bool
}

type MoveByPath struct {
	*MoveComponent
	Troop     bool
	Duration float64
	Elapse   float64
}

func (m *MoveByPath) InitMove(x *MoveComponent) {
	m.MoveComponent = x
	path := GetPathData(m.PathId, m.Troop)
	m.Duration = float64(len(path))
	m.Path = path
	m.Elapse = 0
	m.EndPath = false
}

func (m *MoveByPath) OnUpdate(ms float64) {
	if m.Pause || m.EndPath || m.Path == nil {
		return
	}
	if ms < 0 {
		ms = 1000 / GAME_FPS
	}

	se := float64(ms) / 1000

	if m.Delay > 0 {
		m.Delay -= se
		return
	}

	if m.BeginMove == false && m.Elapse > 0 {
		m.BeginMove = true
	}
	m.Elapse += se * m.Speed

	time := math.Min(1, m.Elapse/m.Duration)

	fIndex := time * m.Duration
	index := int32(fIndex)
	diff := fIndex - float64(index)
	pathSize := int32(len(m.Path))

	if index >= pathSize {
		index = pathSize - 1
	} else if index < 0 || diff < 0 {
		index = 0
		diff = 0
	}

	mp := MovePoint{}
	if index < pathSize-1 {
		p1 := m.Path[index]
		p2 := m.Path[index+1]
		mp.X = p1.X*(1-diff) + (p2.X * diff)
		mp.Y = p1.Y*(1-diff) + (p2.Y * diff)
		mp.Direction = p1.Direction*(1-diff) + (p2.Direction * diff)

		if math.Abs(p1.Direction-p2.Direction) > math.Pi {
			mp.Direction = p1.Direction
		}
	} else {
		mp = m.Path[index]
		m.EndPath = true
	}

	m.X = mp.X + m.Offset.X
	m.Y = mp.Y + m.Offset.Y
	m.Direction = mp.Direction
}

type MoveByDirection struct {
	*MoveComponent
	dx    float64
	dy    float64
	angle float64
	Rebound   bool
	TargetId  int32
}

func (m *MoveByDirection) InitMove(x *MoveComponent) {
	m.MoveComponent = x
	angle := m.MovePoint.Direction
	m.angle = angle
	m.dx = math.Cos(angle - (math.Pi / 2))
	m.dy = math.Sin(angle - (math.Pi / 2))
	m.EndPath = false
}

func (m *MoveByDirection) OnUpdate(ms float64) {
	if m.Pause || m.EndPath {
		return
	}

	if ms < 0 {
		ms = 1000 / GAME_FPS
	}

	se := float64(ms) / 1000
	if m.Delay > 0 {
		m.Delay -= se
		return
	}

	if m.BeginMove == false {
		m.BeginMove = true
	}

	x := m.X + (m.Speed * m.dx * se)
	y := m.Y + (m.Speed * m.dy * se)

	if m.Rebound {
		if x < 0 {
			x = 0 + (0 - x)
			m.dx = -m.dx
			m.angle = -m.angle
		}
		if x > DefaultWidth {
			x = DefaultWidth - (x - DefaultWidth)
			m.dx = -m.dx
			m.angle = -m.angle
		}
		if y < 0 {
			y = 0 + (0 - y)
			m.dy = -m.dy
			m.angle = math.Pi - m.angle
		}
		if y < DefaultHeight {
			y = DefaultHeight - (y - DefaultHeight)
			m.dy = -m.dy
			m.angle = math.Pi - m.angle
		}
	} else {
		if x < 0 || x > DefaultWidth || y < 0 || y > DefaultHeight {
			m.EndPath = true
		}
	}

	m.X = x
	m.Y = y
}
