package pack

import (
	"strconv"
	"strings"
)

// LuHn 银行卡校验方法
func LuHn(s string) bool {
	ss := strings.Split(s, "")
	if len(ss) < 16 {
		return false
	}
	var sum int64
	sum = 0
	for i := 0; i < len(ss)-1; i++ {
		t, _ := strconv.ParseInt(ss[i], 10, 64)
		if i%2 == 0 {
			if t*2 >= 10 {
				sum += (t * 2) - 9
			} else {
				sum += t * 2
			}
		} else {
			sum += t
		}
	}
	lastNum, _ := strconv.ParseInt(ss[len(ss)-1], 10, 64)
	var audit bool
	if (sum+lastNum)%10 == 0 {
		audit = true
	}
	return audit
}
