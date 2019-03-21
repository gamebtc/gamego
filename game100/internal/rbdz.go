package internal

import (
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/room"
)

// 红黑大战
// 游戏玩法：
// 游戏使用1副扑克牌，无大小王
// 红黑各派3张牌

// 0:红赢，1赔1，和黑全输
// 1:黑赢，1赔1，和红全输

//
const (
	Winleopard       = 10*radix + radix //三同10倍
	WinStraightFlush = 5*radix + radix  //顺金5倍
	WinFlush         = 3*radix + radix  //金花3倍
	WinStraight      = 2*radix + radix  //顺子2倍
	WinBigPair       = 1*radix + radix  //大对子(9-A)
)

type RbdzDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

var(
	dealer = model.NewGoldenFlowerDealer(true)
)

func NewRbdzDealer() Dealer {
	d := &RbdzDealer{
	}
	return d
}


func (this *RbdzDealer) Deal(table *Table) {
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
			if !role.IsRobot()  {
				log.Debugf("结算:%v", flow)
			}
		}
	}
}

func (this *RbdzDealer)GetPokers(table *Table)([]byte,[]byte,[]int32) {
	// 检查剩余牌数量
	offset := this.Offset
	if offset >= len(this.Poker)/2 {
		log.Debugf("重新洗牌:%v", this.i)
		this.Poker = model.NewPoker(1, false, true)
		offset = 0
	}
	// 红黑各取3张牌
	a := this.Poker[offset : offset+3]
	b := this.Poker[offset+3 : offset+6]
	odds := rbdzPk(a, b)

	// 系统必须赢
	if table.MustWin() {
		for table.CheckWin(odds) < 0 {
			offset += 1
			if offset >= len(this.Poker)/2 {
				log.Debugf("重新洗牌:%v", this.i)
				this.Poker = model.NewPoker(1, false, true)
				offset = 0
			}
			a = this.Poker[offset : offset+3]
			b = this.Poker[offset+3 : offset+6]
			odds = rbdzPk(a, b)
			log.Debugf("系统无敌:%v,%v,%v", a, b, odds)
		}
	}
	this.Offset = offset + 6
	return a, b, odds
}

func rbdzPk(a []byte, b []byte) (odds []int32) {
	// 可下注的选项数量(0:红赢,1:黑赢,2:幸运一击)
	ag := dealer.GetGroup(a)
	bg := dealer.GetGroup(b)

	redWin := int32(lostRadix)
	blackWin := int32(lostRadix)
	if ag.Power > bg.Power {
		redWin = 1*radix + radix
	} else if ag.Power < bg.Power {
		blackWin = 1*radix + radix
	}

	lucky := int32(lostRadix)
	if ag.IsThreeKind() || bg.IsThreeKind() {
		lucky = Winleopard
		log.Debug("三同10倍")
	} else if ag.IsStraightFlush() || bg.IsStraightFlush() {
		lucky = WinStraightFlush
		log.Debug("顺金5倍")
	} else if ag.IsFlush() || bg.IsFlush() {
		lucky = WinFlush
		log.Debug("金花3倍")
	} else if ag.IsStraight() || bg.IsStraight() {
		lucky = WinStraight
		log.Debug("顺子2倍")
	} else if (ag.Key.Pair()>>8) >= 9 || (bg.Key.Pair()>>8) >= 9 {
		lucky = WinBigPair
		log.Debug("大对子(9-A)")
	}
	return []int32{redWin, blackWin, lucky}
}
