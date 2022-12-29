# 概述

## 介绍

针对http进行封装。

## 使用

```go
import (
	"github.com/aihohu/gotool/httpUtil"
)
```



# GET请求

```go
body := httpUtil.HttpGet(url)

println(body)
```

# POST请求

```go
body := httpUtil.HttpPost(url, data)

println(body)
```

# PUT请求

```go
body := httpUtil.HttpPut(url, data)

println(body)
```

