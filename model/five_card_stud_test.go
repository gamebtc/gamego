package model

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run="TestFiveCard"
// https://blog.csdn.net/hjmnasdkl/article/details/81304329
// go test -v -run="none" -bench=.    不允许单元测试，运行所有的基准测试
// -benchmem 表示分配内存的次数和字节数，-benchtime="3s" 表示持续3秒

func writeFiveCard(t *testing.T, name string, group []*FiveCardGroup) {
	f, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	for i := 0; i < len(group); i++ {
		v := group[i]
		key := v.Key.ToString("")
		//var values string
		//for t := 0; t < len(v.Values)/5; t++ {
		//	values += PokerArrayString(v.Values[t*5:t*5+5], "")
		//	values += "|"
		//}
		//msg := fmt.Sprintf("Weight:%v, number:%v,key:|%v|, values:|%v\n", v.Weight, v.Number, key, values)
		msg := fmt.Sprintf("h:%v, n:%v,k:|%v|\n", v.Weight, v.Number, key)
		f.WriteString(msg)
		//fmt.Println(msg)
	}
	f.Sync()
}

func saveFiveCard(t *testing.T, name string, d *FiveCardDealer) {
	writeFiveCard(t, name+"FiveCard_All.txt", d.All)
	writeFiveCard(t, name+"FiveCard_StraightFlush.txt", d.StraightFlush)
	writeFiveCard(t, name+"FiveCard_FourKind.txt", d.FourKind)
	writeFiveCard(t, name+"FiveCard_FullHouse.txt", d.FullHouse)
	writeFiveCard(t, name+"FiveCard_Flush.txt", d.Flush)
	writeFiveCard(t, name+"FiveCard_Straight.txt", d.Straight)
	writeFiveCard(t, name+"FiveCard_ThreeKind.txt", d.ThreeKind)
	writeFiveCard(t, name+"FiveCard_TwoPair.txt", d.TwoPair)
	writeFiveCard(t, name+"FiveCard_OnePair.txt", d.OnePair)
	writeFiveCard(t, name+"FiveCard_Zilch.txt", d.Zilch)
}

func TestFiveCard(t *testing.T) {
	d := NewFiveCardDealer()
	saveFiveCard(t, "", d)
}

func TestKind(t *testing.T) {
	d := NewFiveCardDealer()
	fmt.Printf("\r\nThreeKind count:%v", len(d.ThreeKind))
	fmt.Printf("\r\nThreeKind:%v", d.ThreeKind)
	writeFiveCard(t, "FiveCard_ThreeKind.txt", d.ThreeKind)

	t.Error("")
	//for i := offset; i >= 0; i-- {
	//	if d.All[i].Key.StraightFlush() == 0 {
	//		d.StraightFlush = d.All[i+1 : offset+1]
	//		offset = i
	//		break
	//	}
	//}
	//for i := offset; i >= 0; i-- {
	//	if d.All[i].Key.Flush() == 0 {
	//		d.Flush = d.All[i+1 : offset+1]
	//		offset = i
	//		break
	//	}
	//}
	//for i := offset; i >= 0; i-- {
	//	if d.All[i].Key.Pair() == 0 {
	//		d.Straight = d.All[i+1 : offset+1]
	//		offset = i
	//		break
	//	}
	//}
	//
	//for i := offset; i >= 0; i-- {
	//	if d.All[i].Key.Pair() == 0 {
	//		d.Straight = d.All[i+1 : offset+1]
	//		d.IsZilch = d.All[0:i]
	//		break
	//	}
	//}
	//
	//for i := 0; i < len(d.All); i++ {
	//	if d.All[i].Key.IsSpecial() == false {
	//		d.IsSpecial = d.All[0:i]
	//		break
	//	}
	//}
}

func TestPower(t *testing.T) {
	InitGlobalFiveCard()
	a := []byte{CK, AK}
	b := []byte{DA}

	d1 := CreateFiveCardDealer(a, b)
	saveFiveCard(t, "A_", d1)

	d2 := CreateFiveCardDealer(b, a)
	saveFiveCard(t, "B_", d2)

	f1, f2 := FiveCardPk(a, b)

	fmt.Printf("\r\n实际胜率:%v, 表面胜率:%v\r\n", f1, f2)

	sum1 := d1.SumPower()
	arv1 := sum1 / len(d1.All)

	sum2 := d2.SumPower()
	arv2 := sum2 / len(d2.All)
	fmt.Printf("\r\nsum1:%v,arv1:%v --- sum2:%v,arv2:%v\r\n", sum1, arv1, sum2, arv2)
}
