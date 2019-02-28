package model

import (
	"sort"
	"strings"
)

const goldenFlowerCount = 3

// 炸金花
func goldenFlowerLess(a GoldenFlower, b GoldenFlower, flower bool) bool {
	// 三同
	if i, j := a.ThreeKind(), b.ThreeKind(); i < j {
		return true
	} else if i > j {
		return false
	} else if flower && (i != 0) {
		return a.flowerLess(b)
	}

	// 顺金
	if i, j := a.StraightFlush(), b.StraightFlush(); i < j {
		return true
	} else if i > j {
		return false
	} else if flower && (i != 0) {
		return a.flowerLess(b)
	}

	// 金花
	if i, j := a.Flush(), b.Flush(); i < j {
		return true
	} else if i > j {
		return false
	} else if flower && (i != 0) {
		return a.flowerLess(b)
	}

	// 顺子
	if i, j := a.Straight(), b.Straight(); i < j {
		return true
	} else if i > j {
		return false
	} else if flower && (i != 0) {
		return a.flowerLess(b)
	}

	// 对子
	if i, j := a.Pair(), b.Pair(); i < j {
		return true
	} else if i > j {
		return false
	} else if flower && (i != 0) {
		return a.flowerLess(b)
	}

	// 单张
	if i, j := a.Zilch(), b.Zilch(); i < j {
		return true
	} else if i > j {
		return false
	} else if flower && (i != 0) {
		return a.flowerLess(b)
	}

	return false
}

// 金花
type GoldenFlower struct {
	p [goldenFlowerCount]byte
	f [goldenFlowerCount]byte
}

func (a GoldenFlower) ToString(split string) string {
	a0 := (a.p[0]) | (a.f[0])
	a1 := (a.p[1]) | (a.f[1])
	a2 := (a.p[2]) | (a.f[2])
	b := strings.Builder{}
	b.Grow(goldenFlowerCount * (4 + len(split)))
	b.WriteString(pokerMap[a0])
	b.WriteString(split)
	b.WriteString(pokerMap[a1])
	b.WriteString(split)
	b.WriteString(pokerMap[a2])
	return b.String()
}

func reversePoker(a byte) byte {
	return ((a & 0x0f) << 4) | (a >> 4)
}

func goldenFlowerNumberSort(a []byte) (byte, byte, byte) {
	a0, a1, a2 := reversePoker(a[0]), reversePoker(a[1]), reversePoker(a[2])
	// 让a0成为最大值
	if a0 < a1 {
		a0, a1 = a1, a0
	}
	if a0 < a2 {
		a0, a2 = a2, a0
	}
	// 让a1成为第二大
	if a1 < a2 {
		a1, a2 = a2, a1
	}
	return reversePoker(a0), reversePoker(a1), reversePoker(a2)
}

func NewGoldenFlower(a []byte) GoldenFlower {
	a0, a1, a2 := goldenFlowerNumberSort(a)
	return GoldenFlower{
		p: [3]byte{PokerPoint(a0), PokerPoint(a1), PokerPoint(a2)},
		f: [3]byte{PokerFlower(a0), PokerFlower(a1), PokerFlower(a2)},
	}
}

func GoldenFlowerNumber(a []byte) uint32 {
	a0, a1, a2 := goldenFlowerNumberSort(a)
	return (uint32(a0) << 16) | (uint32(a1) << 8) | uint32(a2)
}

func (a GoldenFlower) GetNumber() uint32 {
	a0 := a.p[0] | (a.f[0])
	a1 := a.p[1] | (a.f[1])
	a2 := a.p[2] | (a.f[2])
	return (uint32(a0) << 16) | (uint32(a1) << 8) | uint32(a2)
}

// 三同点数
func (a GoldenFlower) ThreeKind() byte {
	if a.p[0] == a.p[1] && a.p[1] == a.p[2] {
		return a.p[0]
	}
	return 0
}

// 顺金点数
func (a GoldenFlower) StraightFlush() byte {
	if a.sameFlower() {
		return a.Straight()
	}
	return 0
}

// 金花点数
func (a GoldenFlower) Flush() uint32 {
	if a.sameFlower() {
		return a.Zilch()
	}
	return 0
}

// 顺子
func (a GoldenFlower) Straight() byte {
	p0 := a.p[0]
	p1 := a.p[1]
	p2 := a.p[2]
	if p1 == (p2 + 1) {
		if p0 == (p1 + 1) {
			return p0
		} else if (p0 == PokerPoint(AA)) && (p1 == PokerPoint(A3)) {
			//A32
			return p1
		}
	}
	return 0
}

// 对子
func (a GoldenFlower) Pair() uint32 {
	if a.p[0] == a.p[1] {
		return (uint32(a.p[1]) << 8) | uint32(a.p[2])
	} else if a.p[1] == a.p[2] {
		return (uint32(a.p[1]) << 8) | uint32(a.p[0])
	}
	return 0
}

// 单牌点数
func (a GoldenFlower) Zilch() uint32 {
	return (uint32(a.p[0]) << 16) | (uint32(a.p[1]) << 8) | uint32(a.p[2])
}

// 花色的点数
func (a GoldenFlower) getFlowerPower() uint32 {
	return (uint32(a.f[0]) << 16) | (uint32(a.f[1]) << 8) | uint32(a.f[2])
}

func (a GoldenFlower) flowerLess(b GoldenFlower) bool {
	return a.getFlowerPower() < b.getFlowerPower()
}

// 不同花色的532
func (a GoldenFlower) Special() bool {
	if a.p[0] == 5 && a.p[1] == 3 && a.p[2] == 2 {
		return !a.sameFlower()
	}
	return false
}

// 是否是相同的花色
func (a GoldenFlower) sameFlower() bool {
	return (a.f[0] == a.f[1]) && (a.f[1] == a.f[2])
}

type GoldenFlowerGroup struct {
	Power  uint32
	Number uint32
	Key    GoldenFlower
	Values [18]byte
}

func (g *GoldenFlowerGroup) IsThreeKind() bool {
	return g.Power >= 22048 && g.Power <= 22099
}

func (g *GoldenFlowerGroup) IsStraightFlush() bool {
	return g.Power >= 22000 && g.Power <= 22047
}

func (g *GoldenFlowerGroup) IsFlush() bool {
	return g.Power >= 20904 && g.Power <= 21999
}

func (g *GoldenFlowerGroup) IsStraight() bool {
	return g.Power >= 20184 && g.Power <= 20903
}

func (g *GoldenFlowerGroup) IsPair() bool {
	return g.Power >= 16440 && g.Power <= 20183
}

func (g *GoldenFlowerGroup) Zilch() bool {
	return g.Power <= 16439
}

func (g *GoldenFlowerGroup) Special() bool {
	return g.Power <= 59
}

type GoldenFlowerDealer struct {
	Flower        bool
	Groups        map[uint32]*GoldenFlowerGroup
	All           []*GoldenFlowerGroup // 全部牌型22100[0-22099]
	ThreeKind     []*GoldenFlowerGroup // 三同52种[22048-22099]
	StraightFlush []*GoldenFlowerGroup // 顺金48种[22000-22047]
	Flush         []*GoldenFlowerGroup // 金花1096[20904-21999]
	Straight      []*GoldenFlowerGroup // 顺子720[20184-20903]
	Pair          []*GoldenFlowerGroup // 对子3744[16440-20183]
	Zilch         []*GoldenFlowerGroup // 散牌16440[0-16439]
	Special       []*GoldenFlowerGroup // 特殊牌[0-59]60
}

func (d *GoldenFlowerDealer) Len() int {
	return len(d.All)
}
func (d *GoldenFlowerDealer) Swap(i, j int) {
	d.All[i], d.All[j] = d.All[j], d.All[i]
}

func (d *GoldenFlowerDealer) Less(i, j int) bool {
	return goldenFlowerLess(d.All[i].Key, d.All[j].Key, d.Flower)
}

func (d *GoldenFlowerDealer) GetGroup(a []byte) *GoldenFlowerGroup {
	num := GoldenFlowerNumber(a)
	return d.Groups[num]
}

func (d *GoldenFlowerDealer) Kind() {
	i := len(d.All) - 1
	//[22048-22099]52
	end := i
	for ; i >= 0 && d.All[i].Key.ThreeKind() != 0; i-- {
	}
	d.ThreeKind = d.All[i+1 : end+1]

	//[22000-22047]48
	end = i
	for ; i >= 0 && d.All[i].Key.StraightFlush() != 0; i-- {
	}
	d.StraightFlush = d.All[i+1 : end+1]

	//[20904-21999]1096
	end = i
	for ; i >= 0 && d.All[i].Key.Flush() != 0; i-- {
	}
	d.Flush = d.All[i+1 : end+1]

	//[20184-20903]720
	end = i
	for ; i >= 0 && d.All[i].Key.Straight() != 0; i-- {
	}
	d.Straight = d.All[i+1 : end+1]

	//[16440-20183]3744
	end = i
	for ; i >= 0 && d.All[i].Key.Pair() != 0; i-- {
	}
	d.Pair = d.All[i+1 : end+1]

	//[0-16439]16440
	end = i
	d.Zilch = d.All[0 : end+1]

	//[0-59]60
	for i = 0; i < len(d.All) && d.All[i].Key.Special(); i++ {
	}
	d.Special = d.All[0:i]
}

func NewGoldenFlowerDealer(flower bool) *GoldenFlowerDealer {
	valOff := make(map[uint32]int, 50000)
	dealer := &GoldenFlowerDealer{
		Flower: flower,
		Groups: make(map[uint32]*GoldenFlowerGroup, 50000),
	}
	const maxIndex = 52
	for i := 0; i < maxIndex; i++ {
		a0 := initPoker[i]
		for t := 0; t < maxIndex; t++ {
			a1 := initPoker[t]
			if a0 == a1 {
				continue
			}
			for j := 0; j < maxIndex; j++ {
				a2 := initPoker[j]
				if a2 == a1 || a2 == a0 {
					continue
				}
				val := []byte{a0, a1, a2}
				key := NewGoldenFlower(val)
				number := key.GetNumber()
				if g, ok := dealer.Groups[number]; ok {
					off := valOff[number]
					g.Values[off], g.Values[off+1], g.Values[off+2] = a0, a1, a2
					valOff[number] = off + goldenFlowerCount
				} else {
					g = &GoldenFlowerGroup{
						Key:    key,
						Number: number,
					}
					g.Values[0], g.Values[1], g.Values[2] = a0, a1, a2
					dealer.Groups[number] = g
					valOff[number] = goldenFlowerCount
					dealer.All = append(dealer.All, g)
				}
			}
		}
		poker := initPoker[i]
		pokerMap[poker] = pokerChar[i]
	}
	sort.Sort(dealer)
	for i := 0; i < len(dealer.All); i++ {
		dealer.All[i].Power = uint32(i)
	}
	dealer.Kind()
	return dealer
}
