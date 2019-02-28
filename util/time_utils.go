package util

import "time"

var (
	ReviseTimeSecond int32 //跨天修正
	TimeZoneSecond   int32
)

// Constants
const (
	WeekTotalDay     = 7 //一周有7天
	SecondsPerMinute = 60
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerDay    = 24 * SecondsPerHour
)

func SetTimeZone(timeZone, reviseTime int32) {
	ReviseTimeSecond = reviseTime * 3600
	TimeZoneSecond = timeZone * 3600
}

func Now() int32 {
	return int32(time.Now().Unix())
}

func LocalNow() int32 {
	return int32(time.Now().Unix()) + TimeZoneSecond
}

func GetWeek(day int32) int32 {
	//1970年1月1日 农历 十一月廿四 星期四
	return (day + 4) % WeekTotalDay
}

func CurrentReviseTime() (day int32, hour int32, minute int32) {
	now := LocalNow()
	day = (now - ReviseTimeSecond) / (24 * 3600)
	hour = (now / 3600) % 24
	minute = (now / 60) % 60
	return
}

// BeginningOfDay method
func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// DayAfterSeconds returns the time after specified seconds.
func DayAfterSeconds(t time.Time, seconds int) time.Time {
	return time.Unix(t.Unix()+int64(seconds), int64(t.Nanosecond()))
}

// DayAfterHours returns the time after specified hours.
func DayAfterHours(t time.Time, hours int) time.Time {
	return DayAfterSeconds(t, hours*3600)
}

// TODO IsSameDay method, change to IsSameDay(t1, t2 int64, zeroHour int64)
// func IsSameDay(t1, t2 time.Time) bool {
// 	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
// }
