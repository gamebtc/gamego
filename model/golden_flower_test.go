package model

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run="TestGod"
// https://blog.csdn.net/hjmnasdkl/article/details/81304329
// go test -v -run="none" -bench=.    不允许单元测试，运行所有的基准测试
// -benchmem 表示分配内存的次数和字节数，-benchtime="3s" 表示持续3秒

func writeGoldenFlower(t *testing.T, name string, group []*GoldenFlowerGroup) {
	f, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	for i := 0; i < len(group); i++ {
		v := group[i]
		key := v.Key.ToString("")
		var values string
		for t := 0; t < 6; t++ {
			values += PokerArrayString(v.Values[t*3:t*3+3], "")
			values += "|"
		}
		msg := fmt.Sprintf("Power:%v, number:%v,key:|%v|, values:|%v\n", v.Power, v.Number, key, values)
		f.WriteString(msg)
		fmt.Println(msg)
	}
	f.Sync()
}

func TestGod(t *testing.T) {
	d := NewGoldenFlowerDealer(true)

	writeGoldenFlower(t, "GoldenFlower_All.txt", d.All)
	writeGoldenFlower(t, "GoldenFlower_ThreeKind.txt", d.ThreeKind)
	writeGoldenFlower(t, "GoldenFlower_StraightFlush.txt", d.StraightFlush)
	writeGoldenFlower(t, "GoldenFlower_Flush.txt", d.Flush)
	writeGoldenFlower(t, "GoldenFlower_Straight.txt", d.Straight)
	writeGoldenFlower(t, "GoldenFlower_Pair.txt", d.Pair)
	writeGoldenFlower(t, "GoldenFlower_Zilch.txt", d.Zilch)
	writeGoldenFlower(t, "GoldenFlower_Special.txt", d.Special)
}
