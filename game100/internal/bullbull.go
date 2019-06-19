package internal

import (
	"local.com/abc/game/model"
)

// 百人牛牛
type BullBullDealer struct {
	i int32
	wild bool
}

var (
	// 亚赔的方式
	bullRate = []int32{1*radix,
		1*radix,
		2*radix,
		3*radix,
		4*radix,
		5*radix,
		6*radix,
		7*radix,
		8*radix,
		9*radix,
		10*radix,
		10*radix,
		10*radix,
	}
)

func NewBullBullDealer() Dealer {
	d := &BullBullDealer{
		wild:true,
	}
	return d
}

func (this *BullBullDealer) Deal(table *Table) ([]byte, []int32, string, bool) {
	offset := 0
	allPoker := model.NewPoker(1, this.wild, true)
	// 庄家取前面5张，闲家取后面20张
	a := allPoker[:model.BullCount]
	bs := allPoker[model.BullCount : 5*model.BullCount]
	odds := bullBullPk(a, bs, this.wild)
	// 系统必须赢
	cheat := false
	if table.mustWin() {
		for table.checkWin(odds) < 0 {
			cheat = true
			offset++
			if offset >= 25 {
				offset = 0
				allPoker = model.NewPoker(1, this.wild, true)
			}
			// 庄家取前面5张，闲家取后面20张
			a = allPoker[offset : offset+model.BullCount]
			bs = allPoker[offset+model.BullCount : offset+(5*model.BullCount)]
			odds = bullBullPk(a, bs, this.wild)
		}
	}
	note := model.PokerArrayString(a) + "|" + model.PokerArrayString(bs)
	poker := allPoker[offset : offset+5*model.BullCount]
	return poker, odds, note, cheat
}

func bullBullPk(a []byte, bs []byte, wild bool) (odds []int32) {
	ag := model.GetBullPower(a, wild)
	// 可下注的选项4个
	odds = []int32{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		player := bs[i*model.BullCount : (i+1)*model.BullCount]
		bg := model.GetBullPower(player, wild)
		bankWin := ag > bg
		if ag == bg {
			bankWin = bullBullBankWin(a, player)
		}
		if bankWin {
			odds[i] = (-bullRate[ag]) + radix
		} else {
			odds[i] = bullRate[bg] + radix
		}
	}
	return odds
}

// 比较单张牌的大小定胜负(先比点数，点数相同比花色)
// 牌点在dtPoint中,从大到小依次为：大王\小王\K\Q\J\10\9\8\7\6\5\4\3\2\A
func bullBullBankWin(a []byte, b []byte) bool {
	// 取a的最大牌
	maxPoker := a[0]
	maxPoint := dtPoint[maxPoker]
	for i := 1; i < model.BullCount; i++ {
		pk := a[i]
		pi := dtPoint[pk]
		if pi > maxPoint || (pi == maxPoint && pk > maxPoker) {
			maxPoker, maxPoint = pk, pi
		}
	}
	// 跟b的每一张牌比较
	for i := 0; i < model.BullCount; i++ {
		pk := b[i]
		pi := dtPoint[pk]
		if pi > maxPoint || (pi == maxPoint && pk > maxPoker) {
			return false
		}
	}
	return true
}

