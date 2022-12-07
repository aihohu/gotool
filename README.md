# gotool
GoTool是一个小而全的Go工具类库



# 安装

```shell
go get github.com/aihohu/gotool
```



# 工具组件库

| 模块             | 介绍           |
| ---------------- | -------------- |
| arrayUtil        | 数组工具       |
| dateUtil         | 时间工具       |
| desensitizedUtil | 数据脱敏工具   |
| fileUtil         | 文件工具       |
| idUtil           | 唯一ID工具     |
| randomUtil       | 随机数工具     |
| stringUtil       | 字符串处理工具 |



# 使用

```go
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

```

