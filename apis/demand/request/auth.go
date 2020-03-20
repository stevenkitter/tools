package request

// AuthParam nonce appId timestamp sign
type AuthParam struct {
	Nonce     string `form:"nonce"`
	AppID     string `form:"appId"`
	Timestamp string `form:"timestamp"`
	Sign      string `form:"sign"`
}
