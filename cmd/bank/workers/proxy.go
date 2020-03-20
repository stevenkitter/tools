package workers

import (
	"encoding/json"
	"log"

	"github.com/stevenkitter/tools/wxhttp"
)

const (
	// ProxyGetAllURL proxy
	ProxyGetAllURL = "http://35.220.159.74:5010/get_all/"
	// ProxyGetURL get
	ProxyGetURL    = "http://35.220.159.74:5010/get"
)

// ProxyResponse ip返回值
type ProxyResponse struct {
	Proxy string `json:"proxy"`
	IP    string `json:"ip"`
	Type  string `json:"type"`
}

// ProxyMan p
// 获取代理地址
// 使用此地址请求资源
// 50秒的时间获取50个可用的代理地址
type ProxyMan struct {
	AddressList []string
}

// NewProxyMan new
func NewProxyMan() ProxyMan {
	return ProxyMan{
		AddressList: make([]string, 0),
	}
}

// RequestAddressList list
func (p *ProxyMan) RequestAddressList() {
	c := wxhttp.Client{}
	rsp, err := c.RequestGet(ProxyGetAllURL, nil)
	if err != nil {
		panic(err)
	}
	var result []*ProxyResponse
	err = json.Unmarshal(rsp, &result)
	if err != nil {
		panic(err)
	}
	for _, i := range result {
		p.AddressList = append(p.AddressList, "http://"+i.Proxy)
	}
}
// NewAddress n
func (p *ProxyMan) NewAddress() string {
	c := wxhttp.Client{}
	rsp, err := c.RequestGet(ProxyGetURL, nil)
	if err != nil {
		panic(err)
	}
	var result ProxyResponse
	err = json.Unmarshal(rsp, &result)
	if err != nil {
		log.Printf("获取新的ip代理出错 %v", err)
		return ""
	}
	return "http://" + result.Proxy
}
