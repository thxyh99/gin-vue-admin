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

// Find 检测切片是否存在key值
func Find(slice []string, index int) (string, bool) {
	for i, item := range slice {
		if i == index {
			return item, true
		}
	}
	return "", false
}

// FindStringValue 检测切片是否存在value值并返回value
func FindStringValue(slice []string, value string) (string, bool) {
	for _, item := range slice {
		if value == item {
			return item, true
		}
	}
	return "", false
}

// FindStringValueKey 检测切片是否存在value值并返回key
func FindStringValueKey(slice []string, value string) (int, bool) {
	for key, item := range slice {
		if value == item {
			return key, true
		}
	}
	return -1, false
}

// FindIntValue 检测切片是否存在value值
func FindIntValue(slice []int64, value int) (int, bool) {
	for _, item := range slice {
		if value == int(item) {
			return value, true
		}
	}
	return 0, false
}

// HideIdNumber 身份证打马赛克
func HideIdNumber(idNumber string) string {
	// 假设身份证号的生日部分是从第7位到第14位
	if len(idNumber) < 14 {
		return "Invalid ID Number"
	}

	hiddenIdNumber := idNumber[:6] + "********" + idNumber[14:]
	return hiddenIdNumber
}
