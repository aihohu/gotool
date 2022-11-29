package dateUtil

import (
	"strings"
	"time"
)

var yyyy, MM, dd, HH, mm, ss = "2006", "01", "02", "15", "04", "05"

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

// NowTimeNoSeparate 没有-的当前时间 yyyyMMddHHmmss
// @return 没有-的当前时间
func NowTimeNoSeparate() string {
	return time.Now().Format("20060102150405")
}

// Format 格式化日期
// @param dateTime 日期
// @param format 日期格式：支持yyyy-MM-dd HH:mm:ss
// @return 格式化后的日期
func Format(dateTime time.Time, format string) string {

	format = strings.ReplaceAll(format, "yyyy", yyyy)
	format = strings.ReplaceAll(format, "MM", MM)
	format = strings.ReplaceAll(format, "dd", dd)
	format = strings.ReplaceAll(format, "HH", HH)
	format = strings.ReplaceAll(format, "mm", mm)
	format = strings.ReplaceAll(format, "ss", ss)

	return dateTime.Format(format)
}

// FormatNow 格式化当前时间
// @param format 日期格式：支持yyyy-MM-dd HH:mm:ss
// @return 格式化后的当前时间
func FormatNow(format string) string {

	format = strings.ReplaceAll(format, "yyyy", yyyy)
	format = strings.ReplaceAll(format, "MM", MM)
	format = strings.ReplaceAll(format, "dd", dd)
	format = strings.ReplaceAll(format, "HH", HH)
	format = strings.ReplaceAll(format, "mm", mm)
	format = strings.ReplaceAll(format, "ss", ss)

	return time.Now().Format(format)
}
