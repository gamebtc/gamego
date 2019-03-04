package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)


// 骰宝
type SbaoDealer struct {
	*rand.Rand
	i      int32
}

var(
	// 执行步骤
	sbaoSchedule = []Plan{
		{f: gameReady, d: second},
		{f: sbaoOpen, d: time.Microsecond},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoBet, d: second},
		{f: sbaoClose, d: second},
		{f: sbaoDeal, d: 2*second},
	}

	// 骰宝税率千分比
	sbaoTaxs = []int64{50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50,50}
)

func NewSbaoDealer(config *model.RoomInfo) *SbaoDealer {
	d := &SbaoDealer{
	}
	d.newRand()
	return d
}

// 开始
func sbaoOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = 2
	log.Debugf("骰宝开始下注:%v", table.CurId)
}

func sbaoBet(table *Table) {
	for _, role := range table.Roles {
		if role.Session == nil && (role.Id%5) == rand.Int31n(5) {
			bet := msg.BetReq{
				Item: rand.Int31n(31),
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
func sbaoClose (table *Table) {
	table.State = 3
	log.Debugf("停止下注:%v", table.CurId)
}

// 发牌结算
func sbaoDeal(table *Table){
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌结算
	table.dealer.Deal(table)
}

func(this *SbaoDealer) Schedule()[]Plan{
	return sbaoSchedule
}

func (this *SbaoDealer) newRand(){
	bin := make([]byte, 8)
	crand.Read(bin)
	seed := binary.LittleEndian.Uint64(bin)
	this.Rand = rand.New(rand.NewSource(int64(seed)))
}

func (this *SbaoDealer) Deal(table *Table) {
	if this.i > 32 {
		this.i = 0
		this.newRand()
	}
	this.i++
	a := []byte{
		byte(this.Int31n(6) + 1),
		byte(this.Int31n(6) + 1),
		byte(this.Int31n(6) + 1),
	}
	// 48=字符串‘0’
	note := string([]byte{
		a[0] + 48,
		a[1] + 48,
		a[2] + 48,
	})
	round := table.round
	round.Odds = sbaoPk(a)
	round.Poker = a
	round.Note = note
	log.Debugf("开牌:%v,%v", note, round.Odds)

	for _, role := range table.Roles {
		if flow := role.Balance(sbaoTaxs); flow != nil {
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

func findPoint(a []byte, b byte) int32 {
	c := int32(0)
	if a[0] == b {
		c++
	}
	if a[1] == b {
		c++
	}
	if a[2] == b {
		c++
	}
	return c
}

func sbaoPk(a []byte) (odds []int32) {
	var i0 int32 = lostRadix  //小(1:1)，点数4-10
	var i1 int32 = lostRadix  //大(1:1)，点数11-17
	var i2 int32 = lostRadix  //单(1:1)，相加总数为单(5,7,9,11,13,15,17)
	var i3 int32 = lostRadix  //双(1:1)，相加总数为双(4,6,8,10,12,14,16)
	var i4 int32 = lostRadix  //单点1(1:1-3)
	var i5 int32 = lostRadix  //单点2(1:1-3)
	var i6 int32 = lostRadix  //单点3(1:1-3)
	var i7 int32 = lostRadix  //单点4(1:1-3)
	var i8 int32 = lostRadix  //单点5(1:1-3)
	var i9 int32 = lostRadix  //单点6(1:1-3)
	var i10 int32 = lostRadix //豹子全围(1:24),大小单双通吃
	var i11 int32 = lostRadix //3个1(1:150)
	var i12 int32 = lostRadix //3个2(1:150)
	var i13 int32 = lostRadix //3个3(1:150)
	var i14 int32 = lostRadix //3个4(1:150)
	var i15 int32 = lostRadix //3个5(1:150)
	var i16 int32 = lostRadix //3个6(1:150)
	var i17 int32 = lostRadix //和4(1:50)
	var i18 int32 = lostRadix //和5(1:18)
	var i19 int32 = lostRadix //和6(1:14)
	var i20 int32 = lostRadix //和7(1:12)
	var i21 int32 = lostRadix //和8(1:8)
	var i22 int32 = lostRadix //和9(1:6)
	var i23 int32 = lostRadix //和10(1:6)
	var i24 int32 = lostRadix //和11(1:6)
	var i25 int32 = lostRadix //和12(1:6)
	var i26 int32 = lostRadix //和13(1:8)
	var i27 int32 = lostRadix //和14(1:12)
	var i28 int32 = lostRadix //和15(1:14)
	var i29 int32 = lostRadix //和16(1:18)
	var i30 int32 = lostRadix //和17(1:50)

	// 三个点的和
	sum := a[0] + a[1] + a[2]
	// 大小
	if sum >= 4 && sum <= 10 {
		i0 = radix*1 + radix
	} else if sum >= 11 && sum <= 17 {
		i1 = radix*1 + radix
	}

	// 单双
	if sum%2 == 1 {
		i2 = radix*1 + radix
	} else {
		i3 = radix*1 + radix
	}

	// 单点1-6
	if c := findPoint(a, 1); c > 0 {
		i4 = radix*c + radix
	}
	if c := findPoint(a, 2); c > 0 {
		i5 = radix*c + radix
	}
	if c := findPoint(a, 3); c > 0 {
		i6 = radix*c + radix
	}
	if c := findPoint(a, 4); c > 0 {
		i7 = radix*c + radix
	}
	if c := findPoint(a, 5); c > 0 {
		i8 = radix*c + radix
	}
	if c := findPoint(a, 6); c > 0 {
		i9 = radix*c + radix
	}

	// 豹子全围
	if a[0] == a[1] && a[1] == a[2] {
		// 大小单双通吃
		i0 = lostRadix
		i1 = lostRadix
		i2 = lostRadix
		i3 = lostRadix
		i10 = radix*24 + radix
		switch a[0] {
		case 1:
			i11 = radix*150 + radix
		case 2:
			i12 = radix*150 + radix
		case 3:
			i13 = radix*150 + radix
		case 4:
			i14 = radix*150 + radix
		case 5:
			i15 = radix*150 + radix
		case 6:
			i16 = radix*150 + radix
		}
	}

	//和
	switch sum {
	case 4:
		i17 = radix*50 + radix
	case 5:
		i18 = radix*18 + radix
	case 6:
		i19 = radix*14 + radix
	case 7:
		i20 = radix*12 + radix
	case 8:
		i21 = radix*8 + radix
	case 9:
		i22 = radix*6 + radix
	case 10:
		i23 = radix*6 + radix
	case 11:
		i24 = radix*6 + radix
	case 12:
		i25 = radix*6 + radix
	case 13:
		i26 = radix*8 + radix
	case 14:
		i27 = radix*12 + radix
	case 15:
		i28 = radix*14 + radix
	case 16:
		i29 = radix*18 + radix
	case 17:
		i30 = radix*50 + radix
	}
	return []int32{i0, i1, i2, i3, i4, i5, i6, i7, i8, i9,
		i10, i11, i12, i13, i14, i15, i16, i17, i18, i19,
		i20, i21, i22, i23, i24, i25, i26, i27, i28, i29, i30,
	}
}

func (this *SbaoDealer) BetItem() int{
	// 可下注的选项数量
	return 31
}