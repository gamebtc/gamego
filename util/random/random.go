package random

import (
	"math/rand"
	"time"
)

// Constants definition
const (
	randomLimit = 10000
)

var (
//random rand.Rand
)

func init() {
	rand.Seed(time.Now().UnixNano())
	//random.Seed(time.Now().UnixNano())
}

// Hit returns if probability > a random number in [0,10000)
func Hit(probability int) bool {
	return Intn(randomLimit) < probability
}

// Intn returns a random number in [0,limit)
func Intn(limit int) int {
	return rand.Intn(limit)
}

// Perm method: same as rand.Perm
func Perm(limit int) []int {
	return rand.Perm(limit)
}

// ChooseWeightItem choose one index from giving array 'candidates'
func ChooseWeightItem(candidates []int, luckyValue int32) int {
	lastIndex := len(candidates) - 1
	if lastIndex < 0 {
		return -1
	}
	maxValue := candidates[lastIndex]
	rValue := Intn(maxValue) + (int(luckyValue) * maxValue / 10000)

	//因为数组已排序,改为二分法查找
	left, right := 0, lastIndex
	for left <= right {
		middle := (left + right) >> 1
		value := candidates[middle]
		if value > rValue {
			right = middle - 1
		} else if value <= rValue {
			left = middle + 1
		}
	}
	if left <= lastIndex {
		return left
	}

	return lastIndex
}
