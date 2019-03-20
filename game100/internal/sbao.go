package internal

import (
	log "github.com/sirupsen/logrus"
	"math/rand"

	"local.com/abc/game/room"
)

// 骰宝
type SbaoDealer struct {
	*rand.Rand
	offset int32
}

func NewSbaoDealer() GameDriver {
	d := &SbaoDealer{
		Rand: room.NewRand(),
	}
	return d
}

// 准备游戏, 状态1
func(this *SbaoDealer)Ready(table *Table){

}
// 开始下注, 状态1
func(this *SbaoDealer)Open(table *Table){

}
// 游戏中
func(this *SbaoDealer)Play(table *Table){

}
// 停止下注
func(this *SbaoDealer)Stop(table *Table){

}

func (this *SbaoDealer) Deal(table *Table) {
	a, odds := this.GetPokers(table)
	// 48=字符串‘0’
	note := string([]byte{a[0] + 48, a[1] + 48, a[2] + 48})
	round := table.round
	round.Odds = odds
	round.Poker = a
	round.Note = note
	log.Debugf("开牌:%v,%v", note, odds)

	for _, role := range table.Roles {
		if flow := role.Balance(); flow != nil {
			room.WriteCoin(flow)
			if !role.IsRobot() {
				log.Debugf("结算:%v", flow)
			}
		}
	}
}

func (this *SbaoDealer)GetPokers(table *Table)([]byte,[]int32) {
	if this.offset > 32 {
		this.offset = 0
		this.Rand = room.NewRand()
	}
	this.offset++
	a := []byte{
		byte(this.Int31n(6) + 1),
		byte(this.Int31n(6) + 1),
		byte(this.Int31n(6) + 1),
	}
	odds := sbaoPk(a)

	// 系统必须赢
	if table.MustWin() {
		for table.CheckWin(odds) < 0 {
			a[0] = byte(this.Int31n(6) + 1)
			a[1] = byte(this.Int31n(6) + 1)
			a[2] = byte(this.Int31n(6) + 1)
			odds = sbaoPk(a)
			log.Debugf("系统无敌:%v,%v", a, odds)
		}
	}
	return a, odds
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

func sbaoPk(a []byte) (i []int32) {
	// 可下注的选项数量31
	// 0   小(1:1)，点数4-10
	// 1   大(1:1)，点数11-17
	// 2   单(1:1)，相加总数为单(5,7,9,11,13,15,17)
	// 3   双(1:1)，相加总数为双(4,6,8,10,12,14,16)
	// 4   单点1(1:1-3)
	// 5   单点2(1:1-3)
	// 6   单点3(1:1-3)
	// 7   单点4(1:1-3)
	// 8   单点5(1:1-3)
	// 9   单点6(1:1-3)
	// 10  豹子全围(1:24),大小单双通吃
	// 11  3个1(1:150)
	// 12  3个2(1:150)
	// 13  3个3(1:150)
	// 14  3个4(1:150)
	// 15  3个5(1:150)
	// 16  3个6(1:150)
	// 17  和4(1:50)
	// 18  和5(1:18)
	// 19  和6(1:14)
	// 20  和7(1:12)
	// 21  和8(1:8)
	// 22  和9(1:6)
	// 23  和10(1:6)
	// 24  和11(1:6)
	// 25  和12(1:6)
	// 26  和13(1:8)
	// 27  和14(1:12)
	// 28  和15(1:14)
	// 29  和16(1:18)
	// 30  和17(1:50)

	i = make([]int32, 31)

	// 三个点的和
	sum := a[0] + a[1] + a[2]

	//小(1:1)，点数4-10
	//大(1:1)，点数11-17
	if sum >= 4 && sum <= 10 {
		i[0] = 1*radix + radix
	} else if sum >= 11 && sum <= 17 {
		i[1] = 1*radix + radix
	}

	// 单(1:1)，相加总数为单(5,7,9,11,13,15,17)
	// 双(1:1)，相加总数为双(4,6,8,10,12,14,16)
	if sum%2 == 1 {
		i[2] = 1*radix + radix
	} else {
		i[3] = 1*radix + radix
	}

	// 单点1-6,赔(1:1-3)
	if c := findPoint(a, 1); c > 0 {
		i[4] = c*radix + radix
	}
	if c := findPoint(a, 2); c > 0 {
		i[5] = c*radix + radix
	}
	if c := findPoint(a, 3); c > 0 {
		i[6] = c*radix + radix
	}
	if c := findPoint(a, 4); c > 0 {
		i[7] = c*radix + radix
	}
	if c := findPoint(a, 5); c > 0 {
		i[8] = c*radix + radix
	}
	if c := findPoint(a, 6); c > 0 {
		i[9] = c*radix+ radix
	}

	//豹子全围(1:24),大小单双通吃
	if a[0] == a[1] && a[1] == a[2] {
		// 大小单双通吃
		i[0] = lostRadix
		i[1] = lostRadix
		i[2] = lostRadix
		i[3] = lostRadix
		i[10] = 24*radix + radix
		//3个一样(1:150)
		switch a[0] {
		case 1:
			i[11] = 150*radix + radix
		case 2:
			i[12] = 150*radix + radix
		case 3:
			i[13] = 150*radix + radix
		case 4:
			i[14] = 150*radix + radix
		case 5:
			i[15] = 150*radix + radix
		case 6:
			i[16] = 150*radix + radix
		}
		log.Debug("豹子全围")
	}

	//和
	switch sum {
	case 4:  //和4(1:50)
		i[17] = 50*radix + radix
	case 5:  //和5(1:18)
		i[18] = 18*radix + radix
	case 6:  //和6(1:14)
		i[19] = 14*radix + radix
	case 7:  //和7(1:12)
		i[20] = 12*radix + radix
	case 8:  //和8(1:8)
		i[21] = 8*radix + radix
	case 9:  //和9(1:6)
		i[22] = 6*radix + radix
	case 10:  //和10(1:6)
		i[23] = 6*radix + radix
	case 11:  //和11(1:6)
		i[24] = 6*radix + radix
	case 12:  //和12(1:6)
		i[25] = 6*radix + radix
	case 13:  //和13(1:8)
		i[26] = 8*radix + radix
	case 14:  //和14(1:12)
		i[27] = 12*radix + radix
	case 15:  //和15(1:14)
		i[28] = 14*radix + radix
	case 16:  //和16(1:18)
		i[29] = 18*radix + radix
	case 17:  //和17(1:50)
		i[30] = 50*radix + radix
	}
	return
}
