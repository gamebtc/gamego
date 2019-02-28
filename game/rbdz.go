package main

import (
	log "github.com/sirupsen/logrus"

	. "local.com/abc/game/model"
)

// 红黑大战
// 游戏玩法：
// 游戏使用1副扑克牌，无大小王
// 红黑各派3张牌

// 0:红赢，1赔1，和黑全输
// 1:黑赢，1赔1，和红全输

//
const (
	leopard       = 10 * radix //三同10倍
	StraightFlush = 5 * radix  //顺金5倍
	Flush         = 3 * radix  //金花3倍
	Straight      = 2 * radix  //顺子2倍
	BigPair       = 1 * radix  //大对子(9-A)
)

type RbdzDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
	dealer *GoldenFlowerDealer
}

func NewRbdzDealer(config *RoomInfo) *RbdzDealer {
	d := &RbdzDealer{
		dealer: NewGoldenFlowerDealer(true),
	}
	return d
}

func (this *RbdzDealer) Deal() (a []byte, b []byte) {
	// 检查剩余牌数量
	if this.Offset >= len(this.Poker)*2/3 {
		log.Debugf("重新洗牌:%v", this.i)
		this.Poker = NewPoker(1, false, true)
		this.Offset = 0
	}
	// 红黑各取1张牌
	a = this.Poker[this.Offset : this.Offset+3]
	b = this.Poker[this.Offset+3 : this.Offset+6]
	this.Offset += 6
	return
}

func (this *RbdzDealer) Pk(a []byte, b []byte) (odds []int32) {
	ag := this.dealer.GetGroup(a)
	bg := this.dealer.GetGroup(b)

	redWin := int32(0)
	blackWin := int32(0)
	draw := int32(-1 * radix)
	if ag.Power > bg.Power {
		redWin = 1 * radix
		blackWin = -1 * radix
	} else if ag.Power < bg.Power {
		redWin = -1 * radix
		blackWin = 1 * radix
	}

	if ag.IsThreeKind() || bg.IsThreeKind() {
		draw = leopard
		log.Debug("三同10倍")
	} else if ag.IsStraightFlush() || bg.IsStraightFlush() {
		draw = StraightFlush
		log.Debug("顺金5倍")
	} else if ag.IsFlush() || bg.IsFlush() {
		draw = Flush
		log.Debug("金花3倍")
	} else if ag.IsStraight() || bg.IsStraight() {
		draw = Straight
		log.Debug("顺子2倍")
	} else if (ag.Key.Pair()>>8) >= 9 || (bg.Key.Pair()>>8) >= 9 {
		draw = BigPair
		log.Debug("大对子(9-A)")
	}

	return []int32{redWin, blackWin, draw}
}

func (this *RbdzDealer) Control(c interface{}) ([]byte, []byte) {
	// 增加控制参数
	return this.Deal()
}

func (this *RbdzDealer) Tax(i int) int64{
	// 所有投注项千分之50的税
	return 50
}

func (this *RbdzDealer) BetItem() int{
	// 可下注的选项数量(0:红赢,1:黑赢,2:幸运一击)
	return 3
}