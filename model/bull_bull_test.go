package model

import (
	"fmt"
	"os"
	"testing"
)

// go test -v -run="TestBull"

func TestBullone(t *testing.T) {
	niu := []byte{A10, AJ, AQ, BlackJoker, AA}
	v := PokerArrayString(niu)
	key := GetBullPower(niu,true)
	msg := fmt.Sprintf("%v, k:%v\n", v, key)
	fmt.Println(msg)
}

func TestBull(t *testing.T) {
	total := 0
	f, err := os.Create("bull.txt")
	if err != nil {
		t.Fatal(err)
	}
	powNum := [13]int{}
	powRate := [13]float64{}
	defer f.Close()
	const maxIndex = 54
	for i0 := 0; i0 < maxIndex-4; i0++ {
		a0 := initPoker[i0]
		for i1 := i0 + 1; i1 < maxIndex-3; i1++ {
			a1 := initPoker[i1]
			for i2 := i1 + 1; i2 < maxIndex-2; i2++ {
				a2 := initPoker[i2]
				for i3 := i2 + 1; i3 < maxIndex-1; i3++ {
					a3 := initPoker[i3]
					for i4 := i3 + 1; i4 < maxIndex; i4++ {
						a4 := initPoker[i4]
						niu := []byte{a0, a1, a2, a3, a4}
						key := GetBullPower(niu, true)
						v := PokerArrayString(niu)
						msg := fmt.Sprintf("%v, k:%v\n", v, key)
						f.WriteString(msg)
						total++
						powNum[key]++
					}
				}
			}
		}
	}
	f.WriteString(fmt.Sprintf("\n%v", powNum))
	for i := 0; i < 13; i++ {
		rate := float64(powNum[i]) / float64(total)
		powRate[i] = rate
		f.WriteString(fmt.Sprintf("\ni:%v,rate:%v", i, rate))
	}
	f.Sync()
}

func TestBull2(t *testing.T) {
	total := 0
	f, err := os.Create("bull_no.txt")
	if err != nil {
		t.Fatal(err)
	}
	powNum := [13]int{}
	powRate := [13]float64{}
	defer f.Close()
	const maxIndex = 52
	for i0 := 0; i0 < maxIndex-4; i0++ {
		a0 := initPoker[i0]
		for i1 := i0 + 1; i1 < maxIndex-3; i1++ {
			a1 := initPoker[i1]
			for i2 := i1 + 1; i2 < maxIndex-2; i2++ {
				a2 := initPoker[i2]
				for i3 := i2 + 1; i3 < maxIndex-1; i3++ {
					a3 := initPoker[i3]
					for i4 := i3 + 1; i4 < maxIndex; i4++ {
						a4 := initPoker[i4]
						niu := []byte{a0, a1, a2, a3, a4}
						key := GetBullPower(niu, false)
						v := PokerArrayString(niu)
						msg := fmt.Sprintf("%v, k:%v\n", v, key)
						f.WriteString(msg)
						total++
						powNum[key]++
					}
				}
			}
		}
	}
	f.WriteString(fmt.Sprintf("\n%v", powNum))
	for i := 0; i < 13; i++ {
		rate := float64(powNum[i]) / float64(total)
		powRate[i] = rate
		f.WriteString(fmt.Sprintf("\ni:%v,rate:%v", i, rate))
	}
	f.Sync()
}