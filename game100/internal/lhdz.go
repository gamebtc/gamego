package internal

import (
	"local.com/abc/game/model"
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
	dragonWin = []int32{1*radix + radix, lostRadix, lostRadix}
	// 1:虎赢，1赔1，和龙全输
	tigerWin = []int32{lostRadix, 1*radix + radix, lostRadix}
	// 2和1赔8, 龙虎全退
	dtTie = []int32{0 + radix, 0 + radix, 8*radix + radix}
	// 龙虎点数映射表
	dtPoint = [66]byte{}
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
	dtPoint[model.BlackJoker], dtPoint[model.RedJoker] = 14, 15
}

// 龙虎大战发牌
type LhdzDealer struct {
	i      int32
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

func NewLhdzDealer() Dealer {
	d := &LhdzDealer{}
	return d
}

func (this *LhdzDealer) Deal(table *Table) ([]byte, []int32, string, bool) {
	// 检查剩余牌数量
	offset := this.Offset
	if offset >= len(this.Poker)/2 {
		this.Poker = model.NewPoker(8, false, true)
		offset = 0
	}
	// 龙虎各取1张牌
	a := this.Poker[offset]
	b := this.Poker[offset+1]
	odds := lhdzPk(a, b)

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
			a = this.Poker[offset]
			b = this.Poker[offset+1]
			odds = lhdzPk(a, b)
		}
	}
	this.Offset = offset + 2

	//
	note := model.PokerString(a) + "|" + model.PokerString(b)
	poker := []byte{a, b}
	return poker, odds, note, cheat
}

func lhdzPk(a byte, b byte) (odds []int32) {
	// 龙虎比较大小(0:龙赢,1:虎赢,2:和)
	pa := dtPoint[a]
	pb := dtPoint[b] //b[0] & 0X0f
	if pa > pb {
		odds = dragonWin
	} else if pa < pb {
		odds = tigerWin
	} else {
		odds = dtTie
	}
	return
}
