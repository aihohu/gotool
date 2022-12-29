# 概述

## 介绍

针对文件进行处理。

## 使用

```go
import (
	"github.com/aihohu/gotool/fileUtil"
)
```



# 保存文件

保存文件，如果文件路径不存在，就创建。

用法：`SaveMkdirAllFile(file *multipart.FileHeader, dst string)`

> file：文件
>
> dst：存储路径
