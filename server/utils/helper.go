package utils

// GetKeyByValue 通过value获取key
func GetKeyByValue(m map[int]string, value string) int {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return -1 // 如果未找到匹配的值，则返回-1或者其他自定义值
}

// GetValueByKey 通过key获取value
func GetValueByKey(m map[int]string, key int) (string, bool) {
	value, ok := m[key]
	return value, ok
}

// InArray 校验目标值是否在数组中
func InArray(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
