package dateUtil

import "time"

// Now 当前时间，格式 yyyy-MM-dd HH:mm:ss
// @return 当前时间的标准形式字符串
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
