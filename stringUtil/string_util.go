package stringUtil

// IsBlank 检查字符串是否为空
func IsBlank(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	}
	return false
}

// IsNotBlank 检查字符串是否为非空
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
