package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)

// 龙虎大战
// 游戏玩法：
// 游戏使用8副扑克牌，无大小王
// 龙派第一张牌，虎派第二张牌，无须补牌
// 根据牌面点数定输赢，点数相同则为和
// 牌型大小：牌点从大到小依次为：K\Q\J\10\9\8\7\6\5\4\3\2\A,不区分花色
// 赔率：下注龙:压中1赔1,开和全退;下注虎:压中1赔1,开和全退;下注和:1赔8

var (
	// 0:龙赢，1赔1，和虎全输
	dragonWin = []int32{radix*1 + radix, lostRadix, lostRadix}
	// 1:虎赢，1赔1，和龙全输
	tigerWin = []int32{lostRadix, radix*1 + radix, lostRadix}
	// 2和1赔8, 龙虎全退
	dtTie = []int32{0 + radix, 0 + radix, radix*8 + radix}
	// 龙虎点数映射表
	dtPoint = [64]byte{}
	// 龙虎税率千分比
	lhdzTaxs = []int64{50, 50, 50}
)

func init() {
	dtPoint[model.AA], dtPoint[model.BA], dtPoint[model.CA], dtPoint[model.DA] = 1, 1, 1, 1
	dtPoint[model.A2], dtPoint[model.B2], dtPoint[model.C2], dtPoint[model.D2] = 2, 2, 2, 2
	dtPoint[model.A3], dtPoint[model.B3], dtPoint[model.C3], dtPoint[model.D3] = 3, 3, 3, 3
	dtPoint[model.A4], dtPoint[model.B4], dtPoint[model.C4], dtPoint[model.D4] = 4, 4, 4, 4
	dtPoint[model.A5], dtPoint[model.B5], dtPoint[model.C5], dtPoint[model.D5] = 5, 5, 5, 5
	dtPoint[model.A6], dtPoint[model.B6], dtPoint[model.C6], dtPoint[model.D6] = 6, 6, 6, 6
	dtPoint[model.A7], dtPoint[model.B7], dtPoint[model.C7], dtPoint[model.D7] = 7, 7, 7, 7
	dtPoint[model.A8], dtPoint[model.B8], dtPoint[model.C8], dtPoint[model.D8] = 8, 8, 8, 8
	dtPoint[model.A9], dtPoint[model.B9], dtPoint[model.C9], dtPoint[model.D9] = 9, 9, 9, 9
	dtPoint[model.A10], dtPoint[model.B10], dtPoint[model.C10], dtPoint[model.D10] = 10, 10, 10, 10
	dtPoint[model.AJ], dtPoint[model.BJ], dtPoint[model.CJ], dtPoint[model.DJ] = 11, 11, 11, 11
	dtPoint[model.AQ], dtPoint[model.BQ], dtPoint[model.CQ], dtPoint[model.DQ] = 12, 12, 12, 12
	dtPoint[model.AK], dtPoint[model.BK], dtPoint[model.CK], dtPoint[model.DK] = 13, 13, 13, 13
}

// 龙虎大战发牌
type LhdzDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

var(
	// 执行步骤
	lhdzSchedule = []Plan{
		{f: gameReady, d: second},
		{f: lhdzOpen, d: time.Microsecond},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzBet, d: second},
		{f: lhdzClose, d: second},
		{f: lhdzDeal, d: 2*second},
	}
)

func NewLhdzDealer(config *model.RoomInfo) *LhdzDealer {
	d := &LhdzDealer{
	}
	return d
}

// 开始
func lhdzOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = 2
	log.Debugf("龙虎大战开始下注:%v", table.CurId)
}

func lhdzBet(table *Table) {
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
func lhdzClose (table *Table) {
	table.State = 3
	log.Debugf("停止下注:%v", table.CurId)
}

// 发牌结算
func lhdzDeal(table *Table){
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌PK
	table.dealer.Deal(table)
}

func(this *LhdzDealer) Schedule()[]Plan{
	return lhdzSchedule
}

func (this *LhdzDealer) Deal(table *Table) {
	// 检查剩余牌数量
	if this.Offset >= len(this.Poker)*2/3 {
		log.Debugf("重新洗牌:%v", this.i)
		this.Poker = model.NewPoker(8, false, true)
		this.Offset = 0
	}
	// 龙虎各取1张牌
	a := this.Poker[this.Offset : this.Offset+1]
	b := this.Poker[this.Offset+1 : this.Offset+2]
	this.Offset += 2

	note := model.PokerArrayString(a) + "|" + model.PokerArrayString(b)
	round := table.round
	round.Odds = lhdzPk(a, b)
	round.Poker = []byte{a[0], b[0]}
	round.Note = note
	// log.Debugf("发牌:%v,%v", note, round.Odds)

	for _, role := range table.Roles {
		if flow := role.Balance(lhdzTaxs); flow != nil {
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

	return
}

func lhdzPk(a []byte, b []byte) (odds []int32) {
	// 龙虎比较大小(0:龙赢,1:虎赢,2:和)
	pa := dtPoint[a[0]]
	pb := dtPoint[b[0]] //b[0] & 0X0f
	if pa > pb {
		odds = dragonWin
	} else if pa < pb {
		odds = tigerWin
	} else {
		odds = dtTie
	}
	return
}

func (this *LhdzDealer) BetItem() int{
	// 可下注的选项数量(0:龙赢,1:虎赢,2:和)
	return 3
}