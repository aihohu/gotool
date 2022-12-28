package main

import (
	"github.com/aihohu/gotool/arrayUtil"
	"github.com/aihohu/gotool/dateUtil"
	"github.com/aihohu/gotool/idUtil"
	"github.com/aihohu/gotool/randomUtil"
	"github.com/aihohu/gotool/stringUtil"
	"time"
)

func main() {
	// 当前时间
	println(dateUtil.Now())
	// 当前日期
	println(dateUtil.Today())

	// 格式化当前时间
	println(dateUtil.FormatNow("yyyy-MM-dd HH:mm:ss"))
	println(dateUtil.FormatNow("yyyy/MM/dd HH:mm:ss"))
	println(dateUtil.FormatNow("yyyy.MM.dd HH:mm:ss"))
	println(dateUtil.FormatNow("yyyy年MM月dd日 HH时mm分ss秒"))
	println(dateUtil.FormatNow("yyyy-MM-dd"))
	println(dateUtil.FormatNow("HH:mm:ss"))

	// 格式化日期
	println(dateUtil.Format(time.Now(), "yyyy-MM-dd HH:mm:ss"))
	println(dateUtil.Format(time.Now(), "yyyy/MM/dd HH:mm:ss"))
	println(dateUtil.Format(time.Now(), "yyyy.MM.dd HH:mm:ss"))
	println(dateUtil.Format(time.Now(), "yyyy年MM月dd日 HH时mm分ss秒"))
	println(dateUtil.Format(time.Now(), "yyyy-MM-dd"))
	println(dateUtil.Format(time.Now(), "HH:mm:ss"))

	// 字符串转日期
	parse := dateUtil.Parse("2022-12-27 13:54:02")
	now := dateUtil.Parse("2022-12-28 14:55:03")

	// 判断两个日期相差的时长，只保留绝对值
	println(dateUtil.Between(parse, now, dateUtil.Ms))
	println(dateUtil.Between(parse, now, dateUtil.Seconds))
	println(dateUtil.Between(parse, now, dateUtil.Minutes))
	println(dateUtil.Between(parse, now, dateUtil.Hours))
	println(dateUtil.Between(parse, now, dateUtil.Days))

	strArr := []string{"g", "o", "t", "o", "o", "l"}
	// strArr 中是否存在 o
	println(arrayUtil.ContainsString("o", strArr))

	// 判断字符串是否为空
	println(stringUtil.IsBlank("gotool"))
	// 判断字符串是否不为空
	println(stringUtil.IsNotBlank("gotool"))

	// 生成唯一id
	println(idUtil.SnowflakeId())
	println(idUtil.RandomUUID())
	println(idUtil.SimpleUUID())

	// 生成随机数
	println(randomUtil.RandomString(6))
	println(randomUtil.RandomNumbers(6))
}
