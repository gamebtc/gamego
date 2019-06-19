package internal

import (
	"local.com/abc/game/model"
)

// 百人牛牛
type BullBullDealer struct {
	i int32
	wild bool
}

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
	b := allPoker[model.BullCount : 5*model.BullCount]
	odds := bullBullPk(a, b, this.wild)
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
			b = allPoker[offset+model.BullCount : offset+(5*model.BullCount)]
			odds = bullBullPk(a, b, this.wild)
		}
	}
	note := model.PokerArrayString(a) + "|" + model.PokerArrayString(b)
	poker := allPoker[offset : offset+5*model.BullCount]
	return poker, odds, note, cheat
}

func bullBullPk(a []byte, b []byte, wild bool) (odds []int32) {
	ag := model.GetBullPower(a, wild)
	// 可下注的选项4个
	odds = []int32{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		bg := model.GetBullPower(b[i*model.BullCount:(i+1)*model.BullCount], wild)
		if ag > bg {

		}
	}
	return odds
}
