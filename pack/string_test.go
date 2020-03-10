package pack

import "testing"

func TestPrefixSection(t *testing.T) {
	d := "12341234123"
	dest := "123412"
	dd := PrefixSection(d, 6)
	if dest != dd {
		t.Errorf("测试无效 得到了 %s 应该是 %s", dd, dest)
	}
}
