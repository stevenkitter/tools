package wechat

import (
	"github.com/go-redis/redis"
)

// Agent 使用redis来管理
type Agent struct {
	cache     *redis.Client
	appID     string
	appSecret string
	mchID     string // 商户号
	payKey    string // 支付密
	capath    string // 证书路径
}

// NewAgent 初始化微信引擎
func NewAgent(cache *redis.Client, appID, appSecret, mchID, payKey, caPath string) *Agent {
	return &Agent{
		cache:     cache,
		appID:     appID,
		appSecret: appSecret,
		mchID:     mchID,
		payKey:    payKey,
		capath:    caPath,
	}
}
