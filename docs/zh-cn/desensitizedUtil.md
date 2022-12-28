# 概述

## 介绍

针对敏感数据进行脱敏处理。

## 使用

```go
import (
	"github.com/aihohu/gotool/desensitizedUtil"
)
```



# 身份证号码脱敏

```go
// 1101***********777
desensitizedUtil.IdCard("110101201801010777")
```

# 手机号码脱敏

```go
// 152****8888
desensitizedUtil.MobilePhone("15266668888")
```

# 固定电话脱敏

```go
// 010-******23
desensitizedUtil.Telephone("010-88993223")
```

# 电子邮箱脱敏

```go
// g*********ool.io
desensitizedUtil.Email("gotool@gotool.io")
```

# 地址脱敏

地址脱敏只支持中文地址。

用法：`desensitizedUtil.Address(中文地址, 敏感信息长度)`，其中**敏感信息长度**小于等于中文地址长度。

```go
// 中国北京海******
desensitizedUtil.Address("中国北京海淀区上地十街", 6)
```

# 自定义数据脱敏

用法：`desensitizedUtil.Hide(数据, 前面的front位数-从1开始, 后面的end位数-从1开始)`

```go
// 01*****789
desensitizedUtil.Hide("0123456789", 2, 3)
```

