package wechat

import (
	"encoding/xml"
	"fmt"

	"github.com/silenceper/wechat/util"
	"github.com/stevenkitter/tools/wxhttp"
)

// 微信转账

// TransferParams t
type TransferParams struct {
	PartnerTradeNo string
	OpenID         string
	ReUserName     string
	Amount         uint64
	Desc           string
	IP             string
}

// transferRequest 请求参数
type transferRequest struct {
	MchAppID       string `xml:"mch_appid"`
	MchID          string `xml:"mchid"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	OpenID         string `xml:"openid"`
	CheckName      string `xml:"check_name"`
	ReUserName     string `xml:"re_user_name"`
	Amount         uint64 `xml:"amount"`
	Desc           string `xml:"desc"`
	SpBillCreateIP string `xml:"spbill_create_ip"`
}

// TransferResponse rsp
type TransferResponse struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	MchAppID       string `xml:"mch_appid"`
	MchID          string `xml:"mchid"`
	NonceStr       string `xml:"nonce_str"`
	ResultCode     string `xml:"result_code"`
	ErrCode        string `xml:"err_code"`
	ErrCodeDes     string `xml:"err_code_des"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentNo      string `xml:"payment_no"`
	PaymentTime    string `xml:"payment_time"`
}

// WeChatTransfer 微信转账
func (a *Agent) WeChatTransfer(p *TransferParams) (rsp TransferResponse, err error) {
	nonceStr := RandomStr(32)
	param := make(map[string]interface{})
	param["mch_appid"] = a.appID
	param["mchid"] = a.mchID
	param["nonce_str"] = nonceStr
	param["partner_trade_no"] = p.PartnerTradeNo
	param["openid"] = p.OpenID
	param["check_name"] = "FORCE_CHECK"
	param["re_user_name"] = p.ReUserName
	param["amount"] = p.Amount
	param["desc"] = p.Desc
	param["spbill_create_ip"] = p.IP

	bizKey := "&key=" + a.payKey
	str := OrderParam(param, bizKey)
	sign := util.MD5Sum(str)
	request := transferRequest{
		MchAppID:       a.appID,
		MchID:          a.mchID,
		NonceStr:       nonceStr,
		Sign:           sign,
		PartnerTradeNo: p.PartnerTradeNo,
		OpenID:         p.OpenID,
		CheckName:      "FORCE_CHECK",
		ReUserName:     p.ReUserName,
		Amount:         p.Amount,
		Desc:           p.Desc,
		SpBillCreateIP: p.IP,
	}
	rawRet, err := wxhttp.PostXMLWithTLS(transferGateway, request, a.capath, a.mchID)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &rsp)
	if err != nil {
		return
	}
	if rsp.ReturnCode == "SUCCESS" {
		if rsp.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = fmt.Errorf("refund error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [params : %s] [sign : %s]",
		string(rawRet), str, sign)
	return
}
