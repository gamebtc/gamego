package main

import (
	log "github.com/sirupsen/logrus"

	. "local.com/abc/game/model"
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
	dragonWin = []int32{1 * radix, -radix, -radix}
	// 1:虎赢，1赔1，和龙全输
	tigerWin = []int32{-radix, 1 * radix, -radix}
	// 2和1赔8, 龙虎全退
	dtDraw = []int32{0, 0, 8 * radix}
	// 龙虎点数映射表
	dtPoint = [64]byte{}
)

func init() {
	dtPoint[AA], dtPoint[BA], dtPoint[CA], dtPoint[DA] = 1, 1, 1, 1
	dtPoint[A2], dtPoint[B2], dtPoint[C2], dtPoint[D2] = 2, 2, 2, 2
	dtPoint[A3], dtPoint[B3], dtPoint[C3], dtPoint[D3] = 3, 3, 3, 3
	dtPoint[A4], dtPoint[B4], dtPoint[C4], dtPoint[D4] = 4, 4, 4, 4
	dtPoint[A5], dtPoint[B5], dtPoint[C5], dtPoint[D5] = 5, 5, 5, 5
	dtPoint[A6], dtPoint[B6], dtPoint[C6], dtPoint[D6] = 6, 6, 6, 6
	dtPoint[A7], dtPoint[B7], dtPoint[C7], dtPoint[D7] = 7, 7, 7, 7
	dtPoint[A8], dtPoint[B8], dtPoint[C8], dtPoint[D8] = 8, 8, 8, 8
	dtPoint[A9], dtPoint[B9], dtPoint[C9], dtPoint[D9] = 9, 9, 9, 9
	dtPoint[A10], dtPoint[B10], dtPoint[C10], dtPoint[D10] = 10, 10, 10, 10
	dtPoint[AJ], dtPoint[BJ], dtPoint[CJ], dtPoint[DJ] = 11, 11, 11, 11
	dtPoint[AQ], dtPoint[BQ], dtPoint[CQ], dtPoint[DQ] = 12, 12, 12, 12
	dtPoint[AK], dtPoint[BK], dtPoint[CK], dtPoint[DK] = 13, 13, 13, 13
}

// 龙虎大战发牌
type LhdzDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

func NewLhdzDealer(config *RoomInfo) *LhdzDealer {
	return new(LhdzDealer)
}

func (this *LhdzDealer) Deal() (a []byte, b []byte) {
	// 检查剩余牌数量
	if this.Offset >= len(this.Poker)*2/3 {
		log.Debugf("重新洗牌:%v", this.i)
		this.Poker = NewPoker(8, false, true)
		this.Offset = 0
	}
	// 龙虎各取1张牌
	a = this.Poker[this.Offset : this.Offset+1]
	b = this.Poker[this.Offset+1 : this.Offset+2]
	this.Offset += 2
	return
}

func (this *LhdzDealer) Pk(a []byte, b []byte) (odds []int32) {
	// 龙虎比较大小(0:龙赢,1:虎赢,2:和)
	pa := dtPoint[a[0]]
	pb := dtPoint[b[0]] //b[0] & 0X0f
	if pa > pb {
		odds = dragonWin
	} else if pa < pb {
		odds = tigerWin
	} else {
		odds = dtDraw
	}
	return
}

func (this *LhdzDealer) Control(c interface{}) ([]byte, []byte) {
	// 增加控制参数
	return this.Deal()
}

func (this *LhdzDealer) Tax(i int) int64{
	// 所有投注项千分之50的税
	return 50
}

func (this *LhdzDealer) BetItem() int{
	//switch config.Kind {
	//case model.GameKind_HHDZ:
	//	betItem = 3 // 可下注的选项数量(0:红赢,1:黑赢,2:幸运一击)
	//case model.GameKind_LHDZ:
	//	betItem = 3 // 可下注的选项数量(0:龙赢,1:虎赢,2:和)
	//case model.GameKind_BJL:
	//	betItem = 5 // 可下注的选项数量(0:庄赢,1:闲赢,2:和,3:庄对,4:闲对)
	//case model.GameKind_SBAO:
	//	betItem = 31 // 可下注的选项数量(0:大赢,1:小赢,2:和,3:庄对,4:闲对)
	//}

	// 可下注的选项数量(0:龙赢,1:虎赢,2:和)
	return 3
}