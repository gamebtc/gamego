package internal

import (
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
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
	dragonWin = []int32{1*radix + radix, lostRadix, lostRadix}
	// 1:虎赢，1赔1，和龙全输
	tigerWin = []int32{lostRadix, 1*radix + radix, lostRadix}
	// 2和1赔8, 龙虎全退
	dtTie = []int32{0 + radix, 0 + radix, 8*radix + radix}
	// 龙虎点数映射表
	dtPoint = [64]byte{}
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

func NewLhdzDealer() GameDriver {
	d := &LhdzDealer{
	}
	return d
}

// 准备游戏, 状态1
func(this *LhdzDealer)Ready(table *Table){

}
// 开始下注, 状态1
func(this *LhdzDealer)Open(table *Table){

}
// 游戏中
func(this *LhdzDealer)Play(table *Table){

}
// 停止下注
func(this *LhdzDealer)Stop(table *Table){

}

func (this *LhdzDealer) Deal(table *Table) {
	a, b, odds := this.GetPokers(table)
	note := model.PokerArrayString(a) + "|" + model.PokerArrayString(b)
	round := table.round
	round.Odds = odds
	round.Poker = []byte{a[0], b[0]}
	round.Note = note
	log.Debugf("发牌:%v,%v", note, odds)
	for _, role := range table.Roles {
		if flow := role.Balance(); flow != nil {
			room.WriteCoin(flow)
			if !role.IsRobot() {
				log.Debugf("结算:%v", flow)
			}
		}
	}
}

func (this *LhdzDealer)GetPokers(table *Table)([]byte,[]byte,[]int32) {
	// 检查剩余牌数量
	offset := this.Offset
	if offset >= len(this.Poker)/2 {
		log.Debugf("重新洗牌:%v", this.i)
		this.Poker = model.NewPoker(8, false, true)
		offset = 0
	}
	// 龙虎各取1张牌
	a := this.Poker[offset : offset+1]
	b := this.Poker[offset+1 : offset+2]

	odds := lhdzPk(a, b)

	// 系统必须赢
	if table.MustWin() {
		for table.CheckWin(odds) < 0 {
			offset += 1
			if offset >= len(this.Poker)/2 {
				log.Debugf("重新洗牌:%v", this.i)
				this.Poker = model.NewPoker(8, false, true)
				offset = 0
			}
			a = this.Poker[offset : offset+1]
			b = this.Poker[offset+1 : offset+2]
			odds = lhdzPk(a, b)
			log.Debugf("系统无敌:%v,%v,%v", a, b, odds)
		}
	}

	this.Offset = offset + 2
	return a, b, odds
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