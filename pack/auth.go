package pack

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Auth(nonce, timestamp, appId, appSecret, sign string) bool {
	ss := strings.Split(nonce, "")
	if len(ss) != 8 {
		return false
	}
	i, err := strconv.Atoi(timestamp)
	if err != nil {
		return false
	}
	d := time.Unix(int64(i), 0)
	t := 30 * time.Second
	if os.Getenv("ENV") == "develop" {
		return true
	}
	if time.Now().Sub(d) > t {
		return false
	}
	s := Signature(nonce, timestamp, appId, appSecret)
	if s == sign {
		return true
	}
	return false
}

//Signature sha1签名
func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetCurrTs() int64 {
	return time.Now().Unix()
}
