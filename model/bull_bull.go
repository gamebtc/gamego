package model

var (
	// 百家乐点数映射表
	BjlPoint = [64]byte{}
)

func init() {
	BjlPoint[AA], BjlPoint[BA], BjlPoint[CA], BjlPoint[DA] = 1, 1, 1, 1
	BjlPoint[A2], BjlPoint[B2], BjlPoint[C2], BjlPoint[D2] = 2, 2, 2, 2
	BjlPoint[A3], BjlPoint[B3], BjlPoint[C3], BjlPoint[D3] = 3, 3, 3, 3
	BjlPoint[A4], BjlPoint[B4], BjlPoint[C4], BjlPoint[D4] = 4, 4, 4, 4
	BjlPoint[A5], BjlPoint[B5], BjlPoint[C5], BjlPoint[D5] = 5, 5, 5, 5
	BjlPoint[A6], BjlPoint[B6], BjlPoint[C6], BjlPoint[D6] = 6, 6, 6, 6
	BjlPoint[A7], BjlPoint[B7], BjlPoint[C7], BjlPoint[D7] = 7, 7, 7, 7
	BjlPoint[A8], BjlPoint[B8], BjlPoint[C8], BjlPoint[D8] = 8, 8, 8, 8
	BjlPoint[A9], BjlPoint[B9], BjlPoint[C9], BjlPoint[D9] = 9, 9, 9, 9
	BjlPoint[A10], BjlPoint[B10], BjlPoint[C10], BjlPoint[D10] = 10, 10, 10, 10
	BjlPoint[AJ], BjlPoint[BJ], BjlPoint[CJ], BjlPoint[DJ] = 20, 20, 20, 20
	BjlPoint[AQ], BjlPoint[BQ], BjlPoint[CQ], BjlPoint[DQ] = 30, 30, 30, 30
	BjlPoint[AK], BjlPoint[BK], BjlPoint[CK], BjlPoint[DK] = 40, 40, 40, 40
}

// 百人牛牛
const(
	BullHard = 12   //硬牛
	BullBull = 11   //牛牛
	BullSoft = 10   //软牛
	Bull9    = 9
	Bull8    = 8
	Bull7    = 7
	Bull6    = 6
	Bull5    = 5
	Bull4    = 4
	Bull3    = 3
	Bull2    = 2
	Bull1    = 1
	Bull0    = 0   //无牛
)

const(
	BullCount = 5
)

func GetBullPower(a []byte, wild bool) byte {
	points := [BullCount]byte{}
	blackIndex := -1
	redIndex := -1
	for i := 0; i < BullCount; i++ {
		p := a[i]
		if p == BlackJoker {
			blackIndex = i
			points[i] = 0
		} else if p == RedJoker {
			redIndex = i
			points[i] = 0
		} else {
			points[i] = BjlPoint[p]
		}
	}

	// 王是否做百搭
	if wild {
		// 有2个王，肯定是牛牛，
		if (blackIndex != -1) && (redIndex != -1) {
			sum := byte(0)
			for i := 0; i < BullCount; i++ {
				sum += points[i]
				if i != blackIndex && i != redIndex {
					if points[i]%10 == 0 {
						// 存在一张10点的牌，硬牛
						return BullHard
					}
				}
			}
			if sum%10 == 0 {
				// 硬牛
				return BullHard
			} else {
				// 软牛
				return BullSoft
			}
		}
		jokerIndex := redIndex
		if blackIndex != -1 {
			jokerIndex = blackIndex
		}
		// 只有1个王，找出最优的解
		if jokerIndex != -1 {
			joker := byte(0)
			maxPower := byte(Bull0)
			for i := byte(0); i < 10; i++ {
				points[jokerIndex] = i
				power := GetBullPoint(points[:])
				if power > maxPower {
					joker = i
					maxPower = power
				}
				// 牛牛
				if maxPower == BullBull {
					if joker == 0 {
						return BullHard
					}
					return BullSoft
				}
			}
			return maxPower
		}
	}
	return GetBullPoint(points[:])
}

func GetBullPoint(points []byte) byte {
	sum := byte(0)
	for i := 0; i < BullCount; i++ {
		sum += points[i]
	}
	pow := sum % 10
	//5选2
	for i := 0; i < BullCount-1; i++ {
		for j := i + 1; j < BullCount; j++ {
			if (points[i]+points[j])%10 == pow {
				if pow == 0 {
					return BullBull
				}
				return pow
			}
		}
	}
	return 0
}