package arrayUtil

// ContainsString strArray[]中是否包含target元素
// @param target 字符串元素
// @param strArray 字符串数组
// @return 是否包含target元素
func ContainsString(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}
