package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
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
	Winleopard       = radix*10 + radix //三同10倍
	WinStraightFlush = radix*5 + radix  //顺金5倍
	WinFlush         = radix*3 + radix  //金花3倍
	WinStraight      = radix*2 + radix  //顺子2倍
	WinBigPair       = radix*1 + radix  //大对子(9-A)
)

type RbdzDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

var(
	// 执行步骤
	rbdzSchedule = []Plan{
		{f: gameReady, d: second},
		{f: rbdzOpen, d: time.Microsecond},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzBet, d: second},
		{f: rbdzClose, d: second},
		{f: rbdzDeal, d: 2*second},
	}

	dealer = model.NewGoldenFlowerDealer(true)

	// 红黑税率千分比
	rbdzTaxs = []int64{50,50,50}
)

func NewRbdzDealer(config *model.RoomInfo) *RbdzDealer {
	d := &RbdzDealer{
	}
	return d
}

// 开始
func rbdzOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = 2
	log.Debugf("龙虎大战开始下注:%v", table.CurId)
}

func rbdzBet(table *Table) {
	for _, role := range table.Roles {
		if role.Session == nil && (role.Id%5) == rand.Int31n(5) {
			bet := msg.BetReq{
				Item: rand.Int31n(3),
				Coin: 100 + rand.Int31n(100)*100,
			}
			if role.RobotCanBet(bet.Item) {
				role.AddBet(bet)
				//log.Debugf("机器人下注:%v,%v", bet, r)
			}
		}
	}
}

// 停止下注
func rbdzClose (table *Table) {
	table.State = 3
	log.Debugf("停止下注:%v", table.CurId)
}

// 发牌结算
func rbdzDeal(table *Table){
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌PK
	table.dealer.Deal(table)
}


func(this *RbdzDealer) Schedule()[]Plan{
	return rbdzSchedule
}

func (this *RbdzDealer) Deal(table *Table) {
	// 检查剩余牌数量
	if this.Offset >= len(this.Poker)*2/3 {
		log.Debugf("重新洗牌:%v", this.i)
		this.Poker = model.NewPoker(1, false, true)
		this.Offset = 0
	}
	// 红黑各取1张牌
	a := this.Poker[this.Offset : this.Offset+3]
	b := this.Poker[this.Offset+3 : this.Offset+6]
	this.Offset += 6

	note := model.PokerArrayString(a) + "|" + model.PokerArrayString(b)
	round := table.round
	round.Odds = rbdzPk(a, b)
	round.Poker = []byte{a[0], a[1], a[2], b[0], b[1], b[2]}
	round.Note = note
	// log.Debugf("发牌:%v,%v", note, round.Odds)

	for _, role := range table.Roles {
		if flow := role.Balance(rbdzTaxs); flow != nil {
			room.WriteCoin(flow)
			if role.Session != nil {
				log.Debugf("结算:%v", flow)
			}
		}
	}
	// 结算结果发给玩家
	table.LastId = table.CurId
	round.End = room.Now()

	room.SaveLog(round)
}

func rbdzPk(a []byte, b []byte) (odds []int32) {
	ag := dealer.GetGroup(a)
	bg := dealer.GetGroup(b)

	redWin := int32(0)
	blackWin := int32(0)
	tie := int32(lostRadix)
	if ag.Power > bg.Power {
		redWin = radix*1 + radix
		blackWin = lostRadix
	} else if ag.Power < bg.Power {
		redWin = lostRadix
		blackWin = radix*1 + radix
	}

	if ag.IsThreeKind() || bg.IsThreeKind() {
		tie = Winleopard
		log.Debug("三同10倍")
	} else if ag.IsStraightFlush() || bg.IsStraightFlush() {
		tie = WinStraightFlush
		log.Debug("顺金5倍")
	} else if ag.IsFlush() || bg.IsFlush() {
		tie = WinFlush
		log.Debug("金花3倍")
	} else if ag.IsStraight() || bg.IsStraight() {
		tie = WinStraight
		log.Debug("顺子2倍")
	} else if (ag.Key.Pair()>>8) >= 9 || (bg.Key.Pair()>>8) >= 9 {
		tie = WinBigPair
		log.Debug("大对子(9-A)")
	}

	return []int32{redWin, blackWin, tie}
}

func (this *RbdzDealer) BetItem() int{
	// 可下注的选项数量(0:红赢,1:黑赢,2:幸运一击)
	return 3
}