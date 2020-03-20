package workers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/stevenkitter/tools/database"
	"github.com/stevenkitter/tools/models/tools"

	"github.com/stevenkitter/tools/wxhttp"

	"log"
	"math/rand"
	"net/url"
	"time"
)

// SnsPath 字节跳动接口
const SnsPath = "https://tp-pay.snssdk.com/gateway-u"

var db *gorm.DB

// SnsResponse 字节跳动提供的数据
type SnsResponse struct {
	Response ResponseInfo `json:"response"`
	Sign     string       `json:"sign"`
	Version  string       `json:"version"`
}

// ResponseInfo response
type ResponseInfo struct {
	Code     string    `json:"code"`
	Msg      string    `json:"msg"`
	CardInfo *CardInfo `json:"card_info"`
}

// CardInfo card 
type CardInfo struct {
	BankCode           string `json:"bank_code"`
	BindCardID         string `json:"bind_card_id"`
	CardLevel          int    `json:"card_level"`
	CardNo             string `json:"card_no"`
	CardNoMask         string `json:"card_no_mask"`
	CardType           string `json:"card_type"`
	CardTypeName       string `json:"card_type_name"`
	CertificateNumMask string `json:"certificate_num_mask"`
	CertificateType    string `json:"certificate_type"`
	FrontBankCode      string `json:"front_bank_code"`
	FrontBankCodeName  string `json:"front_bank_code_name"`
	IconURL            string `json:"icon_url"`
	MobileMask         string `json:"mobile_mask"`
	Msg                string `json:"msg"`
	NeedPwd            string `json:"need_pwd"`
	NeedRepaire        string `json:"need_repaire"`
	NeedSendSms        string `json:"need_send_sms"`
	PerdayLimit        int    `json:"perday_limit"`
	PerpayLimit        int    `json:"perpay_limit"`
	QuickpayMark       string `json:"quickpay_mark"`
	Status             string `json:"status"`
	TrueNameMask       string `json:"true_name_mask"`
}

// Requester 请求银行的信息
type Requester struct {
	ProxyAddress string
}

func init() {
	DBWPath := "35.220.159.74:3306"
	DBPassword := os.Getenv("MYSQL_PWD")
	d, err := JwesUtilsConn(DBWPath, DBPassword)
	if err != nil {
		panic(err)
	}
	db = d
	db.AutoMigrate(&tools.Bank{})
	db.AutoMigrate(&tools.BankBin{})
	db.AutoMigrate(&tools.ErrorBankBin{})
}
// JwesUtilsConn conn
func JwesUtilsConn(dbPath, password string) (*gorm.DB, error) {
	return database.ConnectMysqlDB("tools", password, dbPath, "tools")
}

// RequestBankInfo 请求银行信息
func (r *Requester) RequestBankInfo(no uint64) {
	log.Printf("now no is %d", no)
	s := rand.Intn(2)
	time.Sleep(time.Duration(s) * time.Second)
	cli := wxhttp.Client{ProxyAddress: r.ProxyAddress}
	v := url.Values{}
	v.Add("app_id", "800026247955")
	v.Add("biz_content", fmt.Sprintf(`{"risk_info":{"risk_str":{"ip":"192.168.124.7","version_code":"8.2.2","did":"34678026815","user_agent":"Video 4.2.2 rv:4.2.2.7 (iPhone; iOS 13.3.1; zh_CN)","pay_refer":"","vid":"7E8E8C0D-FC7A-43F0-9094-9C1FB87DB8F4","app_name":"video_article","brand":"Apple","channel":"local_test","device_id":"7E8E8C0D-FC7A-43F0-9094-9C1FB87DB8F4","resolution":"667*375","aid":"32","version":"5.1.9","platform":"2","os_api":"13.3.1","ac":"wifi","os_version":"13.3.1","build_number":"4.2.2.7","device_platform":"iphone","device_type":"iPhone8,1","iid":"103818091056","idfa":"2BAA867B-CE99-4401-9E0B-EE20F831AC20"}},"method":"cashdesk.sdk.card.cardinfo","source":"bind_card","service":"only_card_bin","is_fuzzy_match":true,"merchant_id":"1200002624","card_no":"%d"}`, no))
	v.Add("method", "tp.cashdesk.card_info")
	headers := &map[string]string{
		"x-ss-cookie": "uid_tt=fd16a32bbb0fbdfb50085a01f30720c2; _ga=GA1.2.1654844027.1583304593; d_ticket=5ca9d936744e3e4dfa1563699475ea99711aa; sid_tt=85501f49102bd921b2eb7b75b3249b5c; sdk-version=1; x-Tt-Token=0085501f49102bd921b2eb7b75b3249b5cd68d644998fe43ca10ccf1393b364f963192fe57e649133e9809ead3e751be5d52; uid_tt_ss=fd16a32bbb0fbdfb50085a01f30720c2; odin_tt=8a31d300a04402d084bd5c05bbbb3dae508a4cb8c04655361c74cacdfeca4208022e69d3e43e49aa1c73628fce14d609; sessionid=85501f49102bd921b2eb7b75b3249b5c; sessionid_ss=85501f49102bd921b2eb7b75b3249b5c; ttreq=1$9adf5c442fecc6338cf11e81b8b29f34aea4c1d3; tp_tt_aid=32; install_id=103818091056; sid_guard=85501f49102bd921b2eb7b75b3249b5c%7C1583304183%7C5184000%7CSun%2C+03-May-2020+06%3A43%3A03+GMT",
	}
	d, err := cli.RequestFormURLEncode(SnsPath, v, headers)
	if err != nil {
		log.Printf("请求外在银行卡信息接口出错 %v", err)
		// 换新的代理ip
		p := NewProxyMan()
		pro := p.NewAddress()
		r.ProxyAddress = pro
		var binDest tools.ErrorBankBin
		db.Where(tools.ErrorBankBin{
			BinCode: fmt.Sprintf("%d", no),
		}).Assign(tools.BankBin{}).FirstOrCreate(&binDest)
		return
	}
	db.Unscoped().Where("bin_code = ?", no).Delete(tools.ErrorBankBin{})
	var result SnsResponse
	err = json.Unmarshal(d, &result)
	if err != nil {
		log.Printf("请求外在银行卡信息接口解析对象出错 %v", err)
		return
	}
	if result.Response.Code != "CD0000" && result.Response.CardInfo == nil {
		log.Printf("请求外在银行卡信息接口返回信息 %s", result.Response.Msg)
		return
	}
	codeArray := strings.Split(result.Response.CardInfo.FrontBankCode, "_")
	if len(codeArray) != 2 {
		log.Printf("请求外在银行卡信息返回银行编号有误")
		return
	}

	// 保存银行信息
	// 保存bin对应银行code
	var dest tools.Bank
	db.Where(tools.Bank{
		Code: codeArray[0],
	}).Assign(tools.Bank{
		Name:     result.Response.CardInfo.FrontBankCodeName,
		IconPath: result.Response.CardInfo.IconURL,
	}).FirstOrCreate(&dest)

	var binDest tools.BankBin
	db.Where(tools.BankBin{
		BinCode: fmt.Sprintf("%d", no),
	}).Assign(tools.BankBin{
		Code:     codeArray[0],
		CardType: result.Response.CardInfo.CardTypeName,
	}).FirstOrCreate(&binDest)

}
