package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/stevenkitter/tools/wxhttp"
)

// 正常授权过程 手机端第一次进入本地没有openID 调用微信sdk授权获取code换取基本信息
// 手机端有openID直接进系统 请求相关数据
// 如果返回需要授权 则调用sdk获取code换取基本信息
// 正常情况下直接返回微信基本信息
// 微信登陆 授权 存储用户基本信息 打回jwt用户本地存储
// 携带jwt请求业务接口
//

// ErrorResponse 微信错误信息
type ErrorResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// OauthResponse 微信accessToken返回值
type OauthResponse struct {
	ErrorResponse
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

// UserInfoResponse 微信用户基本信息
type UserInfoResponse struct {
	ErrorResponse
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// OauthError 授权错误类型
type OauthError error

var (
	// ErrNeedAuth 需要重新授权
	ErrNeedAuth        OauthError = errors.New("need auth again")
	errAcceeTokenValid OauthError = errors.New("access_token is invalid")
)

// CodeAccessToken 调起微信app通过code换取凭证
func (a *Agent) CodeAccessToken(code string) (rsp *OauthResponse, err error) {
	m := map[string]string{
		"appid":      appID,
		"secret":     appSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	data, err := a.getWeChatData(accessTokenURL, m)
	if err != nil {
		return
	}
	rsp = &OauthResponse{}
	err = json.Unmarshal(data, rsp)
	if err != nil {
		return
	}
	if rsp.Errcode != 0 {
		err = errors.New(rsp.Errmsg)
		return
	}
	return
}

// SaveOauthToken 保存token到redis 如果保存不了 则也没有问题
// accessToken存入redis加过期时间 refresh_token 也存入数据库
// 分别是2小时 30天
// 使用的时候直接去取 没有值则使用refresh_token来刷新
// key wechat:oauth:openId:accessToken -> accessToken 2hour
// key wechat:oauth:openId:refreshToken -> refreshToken 30day
func (a *Agent) SaveOauthToken(oauth *OauthResponse) error {
	err := a.saveAccessToken(oauth.Openid, oauth.AccessToken)
	if err != nil {
		return err
	}
	return a.saveRefreshToken(oauth.Openid, oauth.RefreshToken)
}

// RefreshToken 刷新token
// 从库里获取refresh_token如果不存在 则提示用户重新授权
// 如果存在则刷新并返回数据 并存下access_token
func (a *Agent) RefreshToken(openID string) (oauth *OauthResponse, err error) {
	k1 := fmt.Sprintf(refreshTokenCacheKey, openID)
	refreshToken := a.cache.Get(k1).Val()
	if refreshToken == "" {
		err = ErrNeedAuth
		return
	}
	data, err := a.getWeChatData(refreshTokenURL, map[string]string{
		"appid":         appID,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	})
	if err != nil {
		return
	}
	oauth = &OauthResponse{}
	err = json.Unmarshal(data, oauth)
	if err != nil {
		return
	}
	if oauth.Errcode != 0 {
		err = errors.New(oauth.Errmsg)
		return
	}
	// 只存accessToken
	a.saveAccessToken(oauth.AccessToken, openID)
	return
}

// GetAccessToken 获取accessToken
// redis 有则打回
// 没有 则请求刷新token 刷新token有则刷新accessToken并打回 刷新token没有则请求重新授权
func (a *Agent) GetAccessToken(openID string) (accessToken string, err error) {
	accessTokenFromCache := a.getAccessTokenFromCache(openID)
	if accessTokenFromCache != "" {
		accessToken = accessTokenFromCache
		return
	}
	rsp, err := a.RefreshToken(openID)
	if err != nil {
		err = ErrNeedAuth
		return
	}
	accessToken = rsp.AccessToken
	return
}

// UserInfo 微信用户基本信息
func (a *Agent) UserInfo(accessToken, openID string) (rsp *UserInfoResponse, err error) {
	data, err := a.getWeChatData(userinfoURL, map[string]string{
		"access_token": accessToken,
		"openid":       openID,
	})
	if err != nil {
		return
	}
	rsp = &UserInfoResponse{}
	err = json.Unmarshal(data, rsp)
	if err != nil {
		return
	}
	if rsp.Errcode != 0 {
		err = errors.New(rsp.Errmsg)
		return
	}
	return
}

// UserInfoFromOpenID openID 获取用户信息
func (a *Agent) UserInfoFromOpenID(openID string) (rsp *UserInfoResponse, err error) {
	accessToken, err := a.GetAccessToken(openID)
	if err != nil {
		return
	}
	return a.UserInfo(accessToken, openID)
}

// GetWeChatData 封装 query请求
func (a *Agent) getWeChatData(path string, param map[string]string) ([]byte, error) {
	htp := wxhttp.Client{}
	querys := MapToPathQuery(param)
	p := ""
	if querys == "" {
		p = path
	} else {
		p = path + "?" + querys
	}
	return htp.RequestGet(p, nil)
}

func (a *Agent) accessTokenValid(accessToken, openID string) (rsp *ErrorResponse, err error) {
	data, err := a.getWeChatData(accessTokenValidURL, map[string]string{
		"access_token": accessToken,
		"openid":       openID,
	})
	if err != nil {
		return
	}
	rsp = &ErrorResponse{}
	err = json.Unmarshal(data, rsp)
	if err != nil {
		return
	}
	if rsp.Errcode != 0 {
		err = errAcceeTokenValid
		return
	}
	return
}

var (
	accessTokenCacheKey  = "wechat:oauth:%s:accessToken"
	refreshTokenCacheKey = "wechat:oauth:%s:refreshToken"
)

func (a *Agent) saveAccessToken(openID, accessToken string) error {
	k0 := fmt.Sprintf(accessTokenCacheKey, openID)
	return a.cache.Set(k0, accessToken, 2*time.Hour).Err()
}
func (a *Agent) getAccessTokenFromCache(openID string) string {
	k0 := fmt.Sprintf(accessTokenCacheKey, openID)
	return a.cache.Get(k0).Val()
}
func (a *Agent) saveRefreshToken(openID, refreshToken string) error {
	k0 := fmt.Sprintf(refreshTokenCacheKey, openID)
	return a.cache.Set(k0, refreshToken, 30*24*time.Hour).Err()
}
func (a *Agent) getRefreshTokenFromCache(openID string) string {
	k0 := fmt.Sprintf(refreshTokenCacheKey, openID)
	return a.cache.Get(k0).Val()
}
