package wxhttp

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 网络请求封装

// Client client
type Client struct {
	ProxyAddress string
}

// RequestJSON json 请求
// path 资源路径
// data json请求数据
func (c *Client) RequestJSON(path string, data []byte, customHeaders *map[string]string) (result []byte, err error) {
	request, err := http.NewRequest("POST", path, bytes.NewReader(data))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	return c.requestResource(request, customHeaders)
}

// RequestFormURLEncode form 请求 application/x-www-form-urlencoded
// path 资源路径
// values 请求数据
func (c *Client) RequestFormURLEncode(path string, values url.Values, customHeaders *map[string]string) (result []byte, err error) {
	request, err := http.NewRequest("POST", path, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.requestResource(request, customHeaders)
}

// RequestGet get
func (c *Client) RequestGet(path string, customHeaders *map[string]string) (result []byte, err error) {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	return c.requestResource(request, customHeaders)
}

// RequestResource 请求资源
// request
// customHeaders
func (c *Client) requestResource(request *http.Request, customHeaders *map[string]string) (result []byte, err error) {
	httpClient := http.Client{}
	if c.ProxyAddress != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(c.ProxyAddress)
		}
		httpClient.Transport = &http.Transport{
			Proxy:           proxy,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	if customHeaders != nil {
		for k, v := range *customHeaders {
			request.Header.Set(k, v)
		}
	}
	rsp, err := httpClient.Do(request)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}
