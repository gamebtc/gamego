package internal

import (
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
)

// 百家乐
type BjlDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

func NewBjlDealer() Dealer {
	d := &BjlDealer{}
	return d
}

func (this *BjlDealer) Deal(table *Table) ([]byte, []int32, string, bool) {
	// 检查剩余牌数量
	offset := this.Offset
	if offset >= len(this.Poker)/2 {
		this.Poker = model.NewPoker(8, false, true)
		offset = 0
	}
	// 闲庄先各取2张牌
	a := []byte{this.Poker[offset], this.Poker[offset+2], 0}
	b := []byte{this.Poker[offset+1], this.Poker[offset+3], 0}
	count := this.repairCard(a, b, offset+4)
	odds := bjlPk(a, b)

	// 系统必须赢
	cheat := false
	if table.mustWin() {
		for table.checkWin(odds) < 0 {
			cheat = true
			offset += 1
			if offset >= len(this.Poker)/2 {
				this.Poker = model.NewPoker(8, false, true)
				offset = 0
			}
			// 闲庄先各取2张牌
			a[0], a[1], a[2] = this.Poker[offset], this.Poker[offset+2], 0
			b[0], b[1], b[2] = this.Poker[offset+1], this.Poker[offset+3], 0
			count = this.repairCard(a, b, offset+4)
			odds = bjlPk(a, b)
		}
	}
	this.Offset = offset + 4 + count

	//
	note := model.PokerArrayString(a) + "|" + model.PokerArrayString(b)
	poker := []byte{a[0], a[1], a[2], b[0], b[1], b[2]}
	return poker, odds, note, cheat
}

func getBjlPoint(a []byte) byte {
	return (model.BjlPoint[a[0]] + model.BjlPoint[a[1]] + model.BjlPoint[a[2]]) % 10
}

// 是否补牌
func (this *BjlDealer) repairCard(a, b []byte, offset int) int {
	count := 0
	pa := getBjlPoint(a)
	pb := getBjlPoint(b)
	// 检查是否补牌
	if pa >= 8 || pb >= 8 || (pa >= 6 && pb >= 6) {
		//任何一家拿到9点（天生赢家），牌局就算结束，不再补牌
	} else {
		var aa byte = 255 //闲家补的牌点数,255表示没有补牌
		//闲家0-5必须博牌
		if pa <= 5 {
			a[2] = this.Poker[offset]
			aa = model.BjlPoint[a[2]] % 10
			offset++
			count++
		}
		addB := false
		// 庄家0-2必须博牌
		if pb <= 2 {
			addB = true
		} else {
			switch pb {
			case 3:
				// 如果闲家补得第三张牌（非三张牌点数相加）是8点，不须补牌，其他则需补牌
				addB = aa != 8
			case 4:
				// 如果闲家补得第三张牌是0,1,8,9点，不须补牌，其他则需补牌
				addB = (aa != 0) && (aa != 1) && (aa != 8) && (aa != 9)
			case 5:
				// 如果闲家补得第三张牌是0,1,2,3,8,9点，不须补牌，其他则需补牌
				addB = (aa != 0) && (aa != 1) && (aa != 2) && (aa != 3) && (aa != 8) && (aa != 9)
			case 6:
				// 如果闲家需补牌（即前提是闲家为1至5点）而补得第三张牌是6或7点，补一张牌，其他则不需补牌
				addB = (aa == 6) || (aa == 7)
			}
		}
		if addB {
			b[2] = this.Poker[offset]
			count++
		}
	}
	return count
}

func bjlPk(a []byte, b []byte) (odds []int32) {
	pa := getBjlPoint(a)
	pb := getBjlPoint(b)
	// 可下注的选项数量(0:闲赢,1:庄赢,2:和,3:闲对,4:庄对)
	odds = []int32{lostRadix, lostRadix, lostRadix, lostRadix, lostRadix}
	if pa > pb {
		// 闲赢
		odds[0] = 1*radix + radix
	} else if pa < pb {
		// 庄赢
		odds[1] = 1*radix + radix
	} else {
		// 和(1赔8，闲/庄退回)
		odds[0], odds[1], odds[2] = radix, radix, 8*radix+radix
	}
	//闲对
	if (model.BjlPoint[a[0]]) == (model.BjlPoint[a[1]]) {
		odds[3] = 11*radix + radix
		log.Debug("闲对")
	}
	//庄对
	if (model.BjlPoint[b[0]]) == (model.BjlPoint[b[1]]) {
		odds[4] = 11*radix + radix
		log.Debug("庄对")
	}
	return odds
}
