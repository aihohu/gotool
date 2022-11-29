package main

import (
	"github.com/aihohu/gotool/arrayUtil"
	"github.com/aihohu/gotool/dateUtil"
	"github.com/aihohu/gotool/idUtil"
	"time"
)

func main() {
	println(dateUtil.Now())
	println(dateUtil.Today())

	strArr := []string{"v", "b", "c"}
	println(arrayUtil.ContainsString("a", strArr))

	for i := 0; i < 2; i++ {
		println(idUtil.SnowflakeId())
		println(idUtil.RandomUUID())
		println(idUtil.SimpleUUID())
	}

	println(dateUtil.FormatNow("yyyy-MM-dd HH:mm:ss"))
	println(dateUtil.FormatNow("yyyy/MM/dd HH:mm:ss"))
	println(dateUtil.FormatNow("yyyy.MM.dd HH:mm:ss"))
	println(dateUtil.FormatNow("yyyy年MM月dd日 HH时mm分ss秒"))
	println(dateUtil.FormatNow("yyyy-MM-dd"))
	println(dateUtil.FormatNow("HH:mm:ss"))

	println(dateUtil.Format(time.Now(), "yyyy-MM-dd HH:mm:ss"))
	println(dateUtil.Format(time.Now(), "yyyy/MM/dd HH:mm:ss"))
	println(dateUtil.Format(time.Now(), "yyyy.MM.dd HH:mm:ss"))
	println(dateUtil.Format(time.Now(), "yyyy年MM月dd日 HH时mm分ss秒"))
	println(dateUtil.Format(time.Now(), "yyyy-MM-dd"))
	println(dateUtil.Format(time.Now(), "HH:mm:ss"))
}
