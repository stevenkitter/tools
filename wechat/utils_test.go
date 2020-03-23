package wechat

import "testing"

func TestMapToPathQuery(t *testing.T) {
	m := map[string]string{
		"ok":   "fine",
		"book": "my",
	}
	should := "ok=fine&book=my"
	result := MapToPathQuery(m)
	if should != result {
		t.Fatal("func MapToPathQuery test failed")
	}
}