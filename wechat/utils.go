package wechat

import (
	"fmt"
	"strings"
)

// MapToPathQuery 拼接map为请求路径的参数
func MapToPathQuery(m map[string]string) string {
	if len(m) == 0 {
		return ""
	}
	ps := make([]string, 0)
	for k, v := range m {
		ps = append(ps, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(ps, "&")
}
