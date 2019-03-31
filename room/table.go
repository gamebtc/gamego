package room

import (
	"time"
)

// 去掉数组结尾的0
func TrimEndZero(a []int64) []int64 {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > 0 {
			return a[:i+1]
		}
	}
	return a
}

// 计划的机器人数量
func PlanRobotCount(roleCount int) int {
	robotConf := Config.Robot
	end := len(robotConf) / 6
	now := time.Now()
	minute := int32(now.Hour()*60 + now.Minute())
	for count := 0; count < end; count++ {
		i := 6 * count
		if minute >= robotConf[i] && minute < robotConf[i+1] {
			min := int(robotConf[i+2])  //最小人数
			max := int(robotConf[i+3])  //最大人数
			base := int(robotConf[i+4]) //基础人数
			rate := int(robotConf[i+5]) //真实玩家的百分比人数
			//真人数量
			count := base + roleCount*rate/100
			if count < min {
				count = min
			} else if count > max {
				count = max
			}
			return count
		}
	}
	return -1
}
