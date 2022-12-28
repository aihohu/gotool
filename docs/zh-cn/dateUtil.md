# 概述

## 介绍

针对时间日期进行处理。

## 使用

```go
import (
	"github.com/aihohu/gotool/dateUtil"
)
```



# 当前时间

获取当前时间

```go
// 输出结果：2022-12-28 22:08:30
dateUtil.Now()
```



获取当前日期

```go
// 输出结果：2022-12-28
dateUtil.Today()
```



没有-的当前时间

```go
// 输出结果：20221228220830
dateUtil.NowTimeNoSeparate()
```



# 格式化日期输出字符串

```go
// 格式化当前时间
dateUtil.FormatNow("yyyy-MM-dd HH:mm:ss")
dateUtil.FormatNow("yyyy/MM/dd HH:mm:ss")
dateUtil.FormatNow("yyyy.MM.dd HH:mm:ss")
dateUtil.FormatNow("yyyy年MM月dd日 HH时mm分ss秒")
dateUtil.FormatNow("yyyy-MM-dd")
dateUtil.FormatNow("HH:mm:ss")

// 格式化日期
dateUtil.Format(time.Now(), "yyyy-MM-dd HH:mm:ss")
dateUtil.Format(time.Now(), "yyyy/MM/dd HH:mm:ss")
dateUtil.Format(time.Now(), "yyyy.MM.dd HH:mm:ss")
dateUtil.Format(time.Now(), "yyyy年MM月dd日 HH时mm分ss秒")
dateUtil.Format(time.Now(), "yyyy-MM-dd")
dateUtil.Format(time.Now(), "HH:mm:ss")
```



# 字符串转日期

```go
dateUtil.Parse("2022-12-28 22:08:30")
```



# 日期时间差

计算两个日期之间的时间差（相差天数、相差小时数等等）：

```go
// 字符串转日期
parse := dateUtil.Parse("2022-12-28 22:08:30")
now := time.Now()

// 判断两个日期相差的时长，只保留绝对值
dateUtil.Between(parse, now, dateUtil.Ms)
dateUtil.Between(parse, now, dateUtil.Seconds)
dateUtil.Between(parse, now, dateUtil.Minutes)
dateUtil.Between(parse, now, dateUtil.Hours)
dateUtil.Between(parse, now, dateUtil.Days)
```