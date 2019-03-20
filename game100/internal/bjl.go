package internal

import (
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/room"
)

var(
	// 百家乐点数映射表
	bjlPoint = [64]byte{}
)
func init() {
	bjlPoint[model.AA], bjlPoint[model.BA], bjlPoint[model.CA], bjlPoint[model.DA] = 1, 1, 1, 1
	bjlPoint[model.A2], bjlPoint[model.B2], bjlPoint[model.C2], bjlPoint[model.D2] = 2, 2, 2, 2
	bjlPoint[model.A3], bjlPoint[model.B3], bjlPoint[model.C3], bjlPoint[model.D3] = 3, 3, 3, 3
	bjlPoint[model.A4], bjlPoint[model.B4], bjlPoint[model.C4], bjlPoint[model.D4] = 4, 4, 4, 4
	bjlPoint[model.A5], bjlPoint[model.B5], bjlPoint[model.C5], bjlPoint[model.D5] = 5, 5, 5, 5
	bjlPoint[model.A6], bjlPoint[model.B6], bjlPoint[model.C6], bjlPoint[model.D6] = 6, 6, 6, 6
	bjlPoint[model.A7], bjlPoint[model.B7], bjlPoint[model.C7], bjlPoint[model.D7] = 7, 7, 7, 7
	bjlPoint[model.A8], bjlPoint[model.B8], bjlPoint[model.C8], bjlPoint[model.D8] = 8, 8, 8, 8
	bjlPoint[model.A9], bjlPoint[model.B9], bjlPoint[model.C9], bjlPoint[model.D9] = 9, 9, 9, 9
	bjlPoint[model.A10], bjlPoint[model.B10], bjlPoint[model.C10], bjlPoint[model.D10] = 10, 10, 10, 10
	bjlPoint[model.AJ], bjlPoint[model.BJ], bjlPoint[model.CJ], bjlPoint[model.DJ] = 20, 20, 20, 20
	bjlPoint[model.AQ], bjlPoint[model.BQ], bjlPoint[model.CQ], bjlPoint[model.DQ] = 30, 30, 30, 30
	bjlPoint[model.AK], bjlPoint[model.BK], bjlPoint[model.CK], bjlPoint[model.DK] = 40, 40, 40, 40
}

// 百家乐
type BjlDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

func NewBjlDealer() GameDriver {
	d := &BjlDealer{
	}
	return d
}

// 等待
func(this *BjlDealer) Wait(table *Table){

}

// 准备游戏, 状态1
func(this *BjlDealer)Ready(table *Table){

}

// 开始下注, 状态1
func(this *BjlDealer)Open(table *Table){

}

// 游戏中
func(this *BjlDealer)Play(table *Table){

}

func (this *BjlDealer) Deal(table *Table) {
	a, b, odds := this.GetPokers(table)
	note := model.PokerArrayString(a) + "|" + model.PokerArrayString(b)
	round := table.round
	round.Odds = odds
	round.Poker = []byte{a[0], a[1], a[2], b[0], b[1], b[2]}
	round.Note = note
	log.Debugf("发牌:%v,%v", note, odds)

	for _, role := range table.Roles {
		if flow := role.Balance(); flow != nil {
			room.WriteCoin(flow)
			log.Debugf("结算:%v", flow)
		}
	}
	for _, role := range table.Robot {
		if flow := role.Balance(); flow != nil {
			//room.WriteCoin(flow)
			//log.Debugf("结算:%v", flow)
		}
	}
}

func getBjlPoint(a []byte)byte {
	return (bjlPoint[a[0]] + bjlPoint[a[1]] + bjlPoint[a[2]]) % 10
}

// 是否补牌
func(this *BjlDealer)RepairCard(a, b []byte, offset int) int {
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
			aa = bjlPoint[a[2]] % 10
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

func (this *BjlDealer)GetPokers(table *Table)([]byte,[]byte,[]int32) {
	// 检查剩余牌数量
	offset := this.Offset
	if offset >= len(this.Poker)/2 {
		log.Debugf("重新洗牌:%v", this.i)
		this.Poker = model.NewPoker(8, false, true)
		offset = 0
	}
	// 闲庄先各取2张牌
	a := []byte{this.Poker[offset], this.Poker[offset+2], 0}
	b := []byte{this.Poker[offset+1], this.Poker[offset+3], 0}
	count := this.RepairCard(a, b, offset+4)
	odds := bjlPk(a, b)

	// 系统必须赢
	if table.MustWin() {
		for table.CheckWin(odds) < 0 {
			offset += 1
			if offset >= len(this.Poker)/2 {
				log.Debugf("重新洗牌:%v", this.i)
				this.Poker = model.NewPoker(8, false, true)
				offset = 0
			}
			// 闲庄先各取2张牌
			a[0], a[1], a[2] = this.Poker[offset], this.Poker[offset+2], 0
			b[0], b[1], b[2] = this.Poker[offset+1], this.Poker[offset+3], 0
			count = this.RepairCard(a, b, offset+4)
			odds = bjlPk(a, b)
			log.Debugf("系统无敌:%v,%v,%v", a, b, odds)
		}
	}
	this.Offset = offset + 4 + count
	return a, b, odds
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
	if (bjlPoint[a[0]]) == (bjlPoint[a[1]]) {
		odds[3] = 11*radix + radix
		log.Debug("闲对")
	}
	//庄对
	if (bjlPoint[b[0]]) == (bjlPoint[b[1]]) {
		odds[4] = 11*radix + radix
		log.Debug("庄对")
	}
	return odds
}