package model

import (
	"sort"
	"strings"
)

const fiveCardCount = 5

// https://zh.wikipedia.org/wiki/%E6%92%B2%E5%85%8B%E7%89%8C%E5%9E%8B
// 五张(同花顺>铁支>葫芦>同花>顺子>三条>二对>对子>散牌)
func fiveCardLess(a FiveCard, b FiveCard) bool {
	// 同花顺
	if i, j := a.StraightFlush(), b.StraightFlush(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 四条
	if i, j := a.FourKind(), b.FourKind(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 葫芦
	if i, j := a.FullHouse(), b.FullHouse(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 同花
	if i, j := a.Flush(), b.Flush(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 顺子
	if i, j := a.Straight(), b.Straight(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 三同
	if i, j := a.ThreeKind(), b.ThreeKind(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 两对
	if i, j := a.TwoPair(), b.TwoPair(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 一对
	if i, j := a.OnePair(), b.OnePair(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}

	// 单张
	if i, j := a.Zilch(), b.Zilch(); i < j {
		return true
	} else if i > j {
		return false
	} else if i != 0 {
		return a.flowerLess(b)
	}
	return false
}

// 五张
type FiveCard struct {
	p [fiveCardCount]byte
	f [fiveCardCount]byte
}

func (a FiveCard) ToString(split string) string {
	a0 := (a.p[0]) | (a.f[0])
	a1 := (a.p[1]) | (a.f[1])
	a2 := (a.p[2]) | (a.f[2])
	a3 := (a.p[3]) | (a.f[3])
	a4 := (a.p[4]) | (a.f[4])
	b := strings.Builder{}
	b.Grow(fiveCardCount * (4 + len(split)))
	b.WriteString(pokerMap[a0])
	b.WriteString(split)
	b.WriteString(pokerMap[a1])
	b.WriteString(split)
	b.WriteString(pokerMap[a2])
	b.WriteString(split)
	b.WriteString(pokerMap[a3])
	b.WriteString(split)
	b.WriteString(pokerMap[a4])
	return b.String()
}

func fiveCardNumberSort(a []byte) (byte, byte, byte, byte, byte) {
	a0, a1, a2, a3, a4 := reversePoker(a[0]), reversePoker(a[1]), reversePoker(a[2]), reversePoker(a[3]), reversePoker(a[4])
	// 让a0成为最大值
	if a0 < a1 {
		a0, a1 = a1, a0
	}
	if a0 < a2 {
		a0, a2 = a2, a0
	}
	if a0 < a3 {
		a0, a3 = a3, a0
	}
	if a0 < a4 {
		a0, a4 = a4, a0
	}
	// 让a1成为第2大
	if a1 < a2 {
		a1, a2 = a2, a1
	}
	if a1 < a3 {
		a1, a3 = a3, a1
	}
	if a1 < a4 {
		a1, a4 = a4, a1
	}
	// 让a2成为第3大
	if a2 < a3 {
		a2, a3 = a3, a2
	}
	if a2 < a4 {
		a2, a4 = a4, a2
	}
	// 让a3成为第4大
	if a3 < a4 {
		a3, a4 = a4, a3
	}

	return reversePoker(a0), reversePoker(a1), reversePoker(a2), reversePoker(a3), reversePoker(a4)
}

func NewFiveCard(a []byte) FiveCard {
	a0, a1, a2, a3, a4 := fiveCardNumberSort(a)
	return FiveCard{
		p: [5]byte{PokerPoint(a0), PokerPoint(a1), PokerPoint(a2), PokerPoint(a3), PokerPoint(a4)},
		f: [5]byte{PokerFlower(a0), PokerFlower(a1), PokerFlower(a2), PokerFlower(a3), PokerFlower(a4)},
	}
}

func FiveCardNumber(a []byte) uint64 {
	a0, a1, a2, a3, a4 := fiveCardNumberSort(a)
	return (uint64(a0) << 32) | (uint64(a1) << 24) | (uint64(a2) << 16) | (uint64(a3) << 8) | uint64(a4)
}

func (a FiveCard) GetNumber() uint64 {
	a0 := a.p[0] | (a.f[0])
	a1 := a.p[1] | (a.f[1])
	a2 := a.p[2] | (a.f[2])
	a3 := a.p[3] | (a.f[3])
	a4 := a.p[4] | (a.f[4])
	return (uint64(a0) << 32) | (uint64(a1) << 24) | (uint64(a2) << 16) | (uint64(a3) << 8) | uint64(a4)
}

// 同花顺点数
func (a FiveCard) StraightFlush() byte {
	if a.sameFlower() {
		return a.Straight()
	}
	return 0
}

// 四条（Four of a Game，亦称“铁支”、“四张”或“炸弹”）：有四张同一点数的牌。
func (a FiveCard) FourKind() uint32 {
	if a.p[1] == a.p[2] && a.p[2] == a.p[3] {
		if a.p[0] == a.p[1] {
			return (uint32(a.p[0]) << 8) | uint32(a.p[4])
		} else if a.p[3] == a.p[4] {
			return (uint32(a.p[4]) << 8) | uint32(a.p[0])
		}
	}
	return 0
}

// 葫芦（Fullhouse，亦称“俘虏”、“骷髅”、“夫佬”、“满堂红”、“富尔豪斯”）：三张同一点数的牌，加一对其他点数的牌。
func (a FiveCard) FullHouse() uint32 {
	if a.p[0] == a.p[1] && a.p[3] == a.p[4] {
		if a.p[1] == a.p[2] {
			return (uint32(a.p[2]) << 8) | uint32(a.p[3])
		} else if a.p[3] == a.p[2] {
			return (uint32(a.p[2]) << 8) | uint32(a.p[1])
		}
	}
	return 0
}

// 同花（Flush，简称“花”：五张同一花色的牌。
func (a FiveCard) Flush() uint32 {
	if a.sameFlower() {
		return a.Zilch()
	}
	return 0
}

// 顺子，五张顺连的牌
func (a FiveCard) Straight() byte {
	p0 := a.p[0]
	p1 := a.p[1]
	p2 := a.p[2]
	p3 := a.p[3]
	p4 := a.p[4]
	if p1 == (p2+1) && p2 == (p3+1) && p3 == (p4+1) {
		if p0 == (p1 + 1) {
			return p0
		} else if (p0 == PokerPoint(AA)) && (p1 == PokerPoint(AJ)) {
			//AJ098
			return p1
		}
	}
	return 0
}

// 三条（Three of a game，亦称“三张”）：有三张同一点数的牌。
func (a FiveCard) ThreeKind() uint32 {
	if a.p[2] == a.p[1] {
		if a.p[0] == a.p[1] {
			//前三张相同
			return (uint32(a.p[2]) << 16) | (uint32(a.p[3]) << 8) | uint32(a.p[4])
		} else if a.p[2] == a.p[3] {
			//中间三张相同
			return (uint32(a.p[2]) << 16) | (uint32(a.p[0]) << 8) | uint32(a.p[4])
		}
	} else if a.p[2] == a.p[3] && a.p[3] == a.p[4] {
		// 后面三张相同
		return (uint32(a.p[2]) << 16) | (uint32(a.p[0]) << 8) | uint32(a.p[1])
	}
	return 0
}

// 两对（Two Pairs，香港称“Two啤”）：两张相同点数的牌，加另外两张相同点数的牌
func (a FiveCard) TwoPair() uint32 {
	if a.p[0] == a.p[1] {
		if a.p[2] == a.p[3] {
			return (uint32(a.p[1]) << 16) | (uint32(a.p[3]) << 8) | uint32(a.p[4])
		} else if a.p[3] == a.p[4] {
			return (uint32(a.p[1]) << 16) | (uint32(a.p[3]) << 8) | uint32(a.p[2])
		}
	} else if a.p[1] == a.p[2] && a.p[3] == a.p[4] {
		return (uint32(a.p[1]) << 16) | (uint32(a.p[3]) << 8) | uint32(a.p[0])
	}
	return 0
}

// 一对（One Pair，香港称“啤”）：两张相同点数的牌。
func (a FiveCard) OnePair() uint32 {
	if a.p[0] == a.p[1] {
		return (uint32(a.p[1]) << 24) | (uint32(a.p[2]) << 16) | (uint32(a.p[3]) << 8) | uint32(a.p[4])
	} else if a.p[1] == a.p[2] {
		return (uint32(a.p[2]) << 24) | (uint32(a.p[0]) << 16) | (uint32(a.p[3]) << 8) | uint32(a.p[4])
	} else if a.p[2] == a.p[3] {
		return (uint32(a.p[3]) << 24) | (uint32(a.p[0]) << 16) | (uint32(a.p[1]) << 8) | uint32(a.p[4])
	} else if a.p[3] == a.p[4] {
		return (uint32(a.p[4]) << 24) | (uint32(a.p[0]) << 16) | (uint32(a.p[1]) << 8) | uint32(a.p[2])
	}
	return 0
}

// 单牌点数
func (a FiveCard) Zilch() uint32 {
	return (uint32(a.p[0]) << 16) | (uint32(a.p[1]) << 12) | (uint32(a.p[2]) << 8) | (uint32(a.p[3]) << 4) | uint32(a.p[4])
}

// 花色的点数
func (a FiveCard) getFlowerPower() uint64 {
	return (uint64(a.f[0]) << 32) | (uint64(a.f[1]) << 24) | (uint64(a.f[2]) << 16) | (uint64(a.f[3]) << 8) | uint64(a.f[4])
}

func (a FiveCard) flowerLess(b FiveCard) bool {
	return a.getFlowerPower() < b.getFlowerPower()
}

// 是否是相同的花色
func (a FiveCard) sameFlower() bool {
	return (a.f[0] == a.f[1]) && (a.f[1] == a.f[2]) && (a.f[2] == a.f[3]) && (a.f[3] == a.f[4])
}

type FiveCardItem struct {
	Weight int32
	Number uint64
	Key    FiveCard
}

type FiveCardGroup struct {
	FiveCardItem
	//Values [600]byte
}

// 同花顺16种[98264-98279]
func (g *FiveCardGroup) IsStraightFlush() bool {
	return g.Weight >= 98264 && g.Weight <= 98279
}

// 铁支168种[98096-98263]
func (g *FiveCardGroup) IsFourKind() bool {
	return g.Weight >= 98096 && g.Weight <= 98263
}

// 葫芦1008[97088-98095]
func (g *FiveCardGroup) IsFullHouse() bool {
	return g.Weight >= 97088 && g.Weight <= 98095
}

// 同花68[97020-97087]
func (g *FiveCardGroup) IsFlush() bool {
	return g.Weight >= 97020 && g.Weight <= 97087
}

// 顺子4080[92940-97019]
func (g *FiveCardGroup) IsStraight() bool {
	return g.Weight >= 92940 && g.Weight <= 97019
}

// 三条6720[86220-92939]
func (g *FiveCardGroup) IsThreeKind() bool {
	return g.Weight >= 86220 && g.Weight <= 92939
}

// 两对15120[71100-86219]
func (g *FiveCardGroup) IsTwoPair() bool {
	return g.Weight >= 71100 && g.Weight <= 86219
}

// 一对53760[17340-71099]
func (g *FiveCardGroup) IsOnePair() bool {
	return g.Weight >= 17340 && g.Weight <= 71099
}

// 散牌17340[0-17339]
func (g *FiveCardGroup) Zilch() bool {
	return g.Weight <= 17339
}

// (同花顺>铁支>葫芦>同花>顺子>三条>二对>对子>散牌)
type FiveCardDealer struct {
	Groups        map[uint64]*FiveCardGroup
	All           []*FiveCardGroup // 全部牌型98280[0-98279]
	StraightFlush []*FiveCardGroup // 同花顺16种[98264-98279]0.016
	FourKind      []*FiveCardGroup // 铁支168种[98096-98263] 0.17
	FullHouse     []*FiveCardGroup // 葫芦1008[97088-98095]  1.0
	Flush         []*FiveCardGroup // 同花68[97020-97087]    0.069
	Straight      []*FiveCardGroup // 顺子4080[92940-97019]  4.15
	ThreeKind     []*FiveCardGroup // 三条6720[86220-92939]  6.83
	TwoPair       []*FiveCardGroup // 两对15120[71100-86219] 15.38
	OnePair       []*FiveCardGroup // 一对53760[17340-71099] 54.7
	Zilch         []*FiveCardGroup // 散牌17340[0-17339]     17.64
	Know          []byte           // 已知牌
}

func (d *FiveCardDealer) Len() int {
	return len(d.All)
}
func (d *FiveCardDealer) Swap(i, j int) {
	d.All[i], d.All[j] = d.All[j], d.All[i]
}

func (d *FiveCardDealer) Less(i, j int) bool {
	return fiveCardLess(d.All[i].Key, d.All[j].Key)
}

func (d *FiveCardDealer) GetGroup(a []byte) *FiveCardGroup {
	num := FiveCardNumber(a)
	return d.Groups[num]
}

func (d *FiveCardDealer) SumPower() (i int) {
	for _, v := range d.All {
		i += int(v.Weight)
	}
	return
}

func (d *FiveCardDealer) Game() {
	i := len(d.All) - 1

	end := i
	for ; i >= 0 && d.All[i].Key.StraightFlush() != 0; i-- {
	}
	d.StraightFlush = d.All[i+1 : end+1]

	end = i
	for ; i >= 0 && d.All[i].Key.FourKind() != 0; i-- {
	}
	d.FourKind = d.All[i+1 : end+1]

	end = i
	for ; i >= 0 && d.All[i].Key.FullHouse() != 0; i-- {
	}
	d.FullHouse = d.All[i+1 : end+1]

	end = i
	for ; i >= 0 && d.All[i].Key.Flush() != 0; i-- {
	}
	d.Flush = d.All[i+1 : end+1]

	end = i
	for ; i >= 0 && d.All[i].Key.Straight() != 0; i-- {
	}
	d.Straight = d.All[i+1 : end+1]

	end = i
	for ; i >= 0 && d.All[i].Key.ThreeKind() != 0; i-- {
	}
	d.ThreeKind = d.All[i+1 : end+1]

	end = i
	for ; i >= 0 && d.All[i].Key.TwoPair() != 0; i-- {
	}
	d.TwoPair = d.All[i+1 : end+1]

	end = i
	for ; i >= 0 && d.All[i].Key.OnePair() != 0; i-- {
	}
	d.OnePair = d.All[i+1 : end+1]

	end = i
	d.Zilch = d.All[0 : end+1]
}

func (d *FiveCardDealer) Pk(a *FiveCardDealer) (win int, lost int, draw int) {
	for _, item1 := range d.All {
		dp := item1.Weight
		for _, item2 := range a.All {
			ap := item2.Weight
			if dp > ap {
				win += 1
			} else if dp < ap {
				lost += 1
			} else {
				draw += 1
			}
		}
	}
	return
}

func (d *FiveCardDealer) WinRate(b *FiveCardDealer) (float32, float32) {
	win1, lost1, draw1 := d.Pk(b)
	f1 := float32(win1*10000/(win1+lost1+draw1)) / 100

	a2 := CreateFiveCardDealer(d.Know[1:], b.Know)
	b2 := CreateFiveCardDealer(b.Know, d.Know[1:])
	win2, lost2, draw2 := a2.Pk(b2)
	f2 := float32(win2*10000/(win2+lost2+draw2)) / 100
	return f1, f2
}

var initFiveCard = [28]byte{
	A8, A9, A10, AJ, AQ, AK, AA,
	B8, B9, B10, BJ, BQ, BK, BA,
	C8, C9, C10, CJ, CQ, CK, CA,
	D8, D9, D10, DJ, DQ, DK, DA,
}

var globalDealer *FiveCardDealer

func InitGlobalFiveCard() {
	globalDealer = NewFiveCardDealer()
}

func NewFiveCardDealer() *FiveCardDealer {
	valOff := make(map[uint64]int, 98280*2)
	dealer := &FiveCardDealer{
		Groups: make(map[uint64]*FiveCardGroup, 98280*2),
	}
	const maxIndex = 28
	for i := 0; i < maxIndex; i++ {
		a0 := initFiveCard[i]
		for t := 0; t < maxIndex; t++ {
			a1 := initFiveCard[t]
			if a0 == a1 {
				continue
			}
			for j := 0; j < maxIndex; j++ {
				a2 := initFiveCard[j]
				if a2 == a1 || a2 == a0 {
					continue
				}
				for m := 0; m < maxIndex; m++ {
					a3 := initFiveCard[m]
					if a3 == a2 || a3 == a1 || a3 == a0 {
						continue
					}
					for n := 0; n < maxIndex; n++ {
						a4 := initFiveCard[n]
						if a4 == a3 || a4 == a2 || a4 == a1 || a4 == a0 {
							continue
						}
						val := []byte{a0, a1, a2, a3, a4}
						key := NewFiveCard(val)
						number := key.GetNumber()
						if g, ok := dealer.Groups[number]; ok {
							off := valOff[number]
							//g.Values[off], g.Values[off+1], g.Values[off+2], g.Values[off+3], g.Values[off+4] = a0, a1, a2, a3, a4
							valOff[number] = off + fiveCardCount
						} else {
							g = new(FiveCardGroup)
							g.Key, g.Number = key, number
							//g.Values[0], g.Values[1], g.Values[2], g.Values[3], g.Values[4] = a0, a1, a2, a3, a4
							dealer.Groups[number] = g
							valOff[number] = fiveCardCount
							dealer.All = append(dealer.All, g)
						}
					}
				}
			}
		}
		poker := initPoker[i]
		pokerMap[poker] = pokerChar[i]
	}
	sort.Sort(dealer)
	for i := 0; i < len(dealer.All); i++ {
		dealer.All[i].Weight = int32(i)
	}
	dealer.Game()
	return dealer
}

func FiveCardExists(b []byte, a byte) bool {
	for _, k := range b {
		if k == a {
			return true
		}
	}
	return false
}

func CreateFiveCardDealer(a []byte, b []byte) *FiveCardDealer {
	// 获取A所有可能的组合，组合的平均牌力
	const maxIndex = 28
	a0 := a[0]
	l := len(a)
	dealer := &FiveCardDealer{
		Groups: make(map[uint64]*FiveCardGroup, 98280/l),
		Know:   make([]byte, 0, 5),
	}
	dealer.Know = append(dealer.Know, a...)
	for t := 0; t < maxIndex; t++ {
		var a1 byte
		if l >= 2 && a[1] > 0 {
			t = maxIndex
			a1 = a[1]
		} else {
			a1 = initFiveCard[t]
			if a0 == a1 || FiveCardExists(b, a1) {
				continue
			}
		}
		for j := 0; j < maxIndex; j++ {
			var a2 byte
			if l >= 3 && a[2] > 0 {
				j = maxIndex
				a2 = a[2]
			} else {
				a2 = initFiveCard[j]
				if a2 == a1 || a2 == a0 || FiveCardExists(b, a2) {
					continue
				}
			}
			for m := 0; m < maxIndex; m++ {
				var a3 byte
				if l >= 4 && a[3] > 0 {
					m = maxIndex
					a3 = a[3]
				} else {
					a3 = initFiveCard[m]
					if a3 == a2 || a3 == a1 || a3 == a0 || FiveCardExists(b, a3) {
						continue
					}
				}
				for n := 0; n < maxIndex; n++ {
					var a4 byte
					if l >= 5 && a[4] > 0 {
						m = maxIndex
						a4 = a[4]
					} else {
						a4 = initFiveCard[n]
						if a4 == a3 || a4 == a2 || a4 == a1 || a4 == a0 || FiveCardExists(b, a4) {
							continue
						}
					}
					val := []byte{a0, a1, a2, a3, a4}
					key := NewFiveCard(val)
					number := key.GetNumber()
					if g, ok := dealer.Groups[number]; !ok {
						g = new(FiveCardGroup)
						g.Key, g.Number = key, number
						g.Weight = globalDealer.Groups[number].Weight
						dealer.Groups[number] = g
						dealer.All = append(dealer.All, g)
					}
				}
			}
		}
	}
	sort.Sort(dealer)
	dealer.Game()
	return dealer
}

func FiveCardPk(a, b []byte) (float32, float32) {
	d1 := CreateFiveCardDealer(a, b)
	d2 := CreateFiveCardDealer(b, a)
	f1, f2 := d1.WinRate(d2)
	return f1, f2
}
