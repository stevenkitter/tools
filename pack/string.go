package pack

import "strings"


// Contain 包含
func Contain(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}
// PrefixSection 字符中截取前num段
func PrefixSection(dest string, num int) string {
	arr := strings.Split(dest, "")
	if len(arr) < num {
		return ""
	}
	d := arr[:num]
	return strings.Join(d, "")
}
