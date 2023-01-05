package dateUtil

import (
	"strings"
	"time"
)

var yyyy, month, dd, hh, mm, ss = "2006", "01", "02", "15", "04", "05"
var (
	Ms      = "MS"
	Seconds = "SECOND"
	Minutes = "MINUTE"
	Hours   = "HOUR"
	Days    = "DAY"
)

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
	format = strings.ReplaceAll(format, "MM", month)
	format = strings.ReplaceAll(format, "dd", dd)
	format = strings.ReplaceAll(format, "HH", hh)
	format = strings.ReplaceAll(format, "mm", mm)
	format = strings.ReplaceAll(format, "ss", ss)

	return dateTime.Format(format)
}

// FormatNow 格式化当前时间
// @param format 日期格式：支持yyyy-MM-dd HH:mm:ss
// @return 格式化后的当前时间
func FormatNow(format string) string {

	format = strings.ReplaceAll(format, "yyyy", yyyy)
	format = strings.ReplaceAll(format, "MM", month)
	format = strings.ReplaceAll(format, "dd", dd)
	format = strings.ReplaceAll(format, "HH", hh)
	format = strings.ReplaceAll(format, "mm", mm)
	format = strings.ReplaceAll(format, "ss", ss)

	return time.Now().Format(format)
}

// FormatInLocation 格式化时区日期
// @param dateTime 日期
// @param format 日期格式：支持yyyy-MM-dd HH:mm:ss
// @param name   时区：如 Asia/Shanghai
// @return 格式化后的日期
func FormatInLocation(dateTime time.Time, format string, name string) string {
	location, err := time.LoadLocation(name)
	if err != nil {
		return Format(dateTime, format)
	}

	return Format(dateTime.In(location), format)
}

// Parse 字符串转日期
// @param format 日期字符串
// @return 日期
func Parse(format string) time.Time {
	location, err := time.ParseInLocation("2006-01-02 15:04:05", format, time.Local)
	if err != nil {
		return time.Time{}
	}

	return location
}

// Between 判断两个日期相差的时长，只保留绝对值
// @param beginDate 起始日期
// @param endDate   结束日期
// @param unit      相差的单位
// @return 日期差
func Between(beginDate time.Time, endDate time.Time, unit string) int {
	sub := endDate.Sub(beginDate)

	if unit == Ms {
		return int(sub.Abs().Milliseconds())
	} else if unit == Seconds {
		return int(sub.Abs().Seconds())
	} else if unit == Minutes {
		return int(sub.Abs().Minutes())
	} else if unit == Hours {
		return int(sub.Abs().Hours())
	} else if unit == Days {
		return int(sub.Abs().Hours() / 24)
	}

	return 0
}
