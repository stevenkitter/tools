package pack

import (
	"fmt"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {
	nonce := RandomStr(8)
	timestampStr := fmt.Sprintf("%d", GetCurrTs())
	id := "1Yv4vFyYK27tOdh1CTtL25ObC19"
	secret := "be010be1127d3547ec72d57561e61b8a81a03209"
	s := Signature(nonce, timestampStr, id, secret)
	log.Printf("ok %s", s)
}
