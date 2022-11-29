package dateUtil

import "time"

// Now 当前时间，格式 yyyy-MM-dd HH:mm:ss
// @return 当前时间的标准形式字符串
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Today 当前日期字符串，格式：yyyy-MM-dd
// @return 当前日期的标准形式字符串
func Today() string {
	return time.Now().Format("2006-01-02")
}
