package workers

import (
	"encoding/json"
	"github.com/stevenkitter/tools/pack"
	"github.com/stevenkitter/tools/wxHttp"
)

const ProxyGetURL = "http://localhost:8080/v2/ip"

type ProxyResponse struct {
	Ip   string `json:"ip"`
	Type string `json:"type"`
}

// ProxyMan
// 获取代理地址
// 使用此地址请求资源
// 50秒的时间获取50个可用的代理地址
type ProxyMan struct {
	AddressList []string
}

func NewProxyMan() ProxyMan {
	return ProxyMan{
		AddressList: make([]string, 0),
	}
}

func (p *ProxyMan) RequestAddressList() {
	c := wxHttp.Client{}
	rsp, err := c.RequestGet(ProxyGetURL, nil)
	if err != nil {
		panic(err)
	}
	var result ProxyResponse
	err = json.Unmarshal(rsp, &result)
	if err != nil {
		panic(err)
	}
	if !pack.Contain(p.AddressList, result.Ip) {
		p.AddressList = append(p.AddressList, result.Ip)
	}
}
