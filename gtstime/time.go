package gtstime

import (
	"time"
)

// return 2006-01-02 15:04:05
func GetNow() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}

// return 2006-01-02
func GetNowDay() string {
	tm := time.Now()
	return tm.Format("2006-01-02")
}

// get three month ago
func GetXMonthsAgo(x int) string {
	nowTime := time.Now()
	getTime := nowTime.AddDate(0, -x, 0)    //年，月，日   获取x个月前的时间
	resTime := getTime.Format("2006-01-02") //获取的时间的格式
	return resTime
}

// 获取当前天0点的时间戳-毫秒  2020-10-19 00:00:00 -> 1603076116000
func StrtodayToUnixNanoTime(day string) int64 {
	day = day + " 00:00:00"
	formatTime, _ := time.ParseInLocation("2006-01-02 15:04:05", day, time.Local)
	return formatTime.UnixNano() / 1e6
}

// 获取当前天23:59:59点的时间戳-毫秒  2020-10-19 23:59:59 -> 1603076116000
func Strtoday24ToUnixNanoTime(day string) int64 {
	day = day + " 23:59:59"
	formatTime, _ := time.ParseInLocation("2006-01-02 15:04:05", day, time.Local)
	return formatTime.UnixNano() / 1e6
}
