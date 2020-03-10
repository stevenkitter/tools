package controller

import (
	"errors"
	"github.com/stevenkitter/tools/pack"
)

type CardInfo struct {
	BankCode string `json:"bank_code"`
	BankName string `json:"bank_name"`
	CardType string `json:"card_type"`
	IconPath string `json:"icon_path"`
}

// CardInfoBusiness 银行卡bin信息
func (ct *Controller) CardInfoBusiness(cardNo string) (result *CardInfo, err error) {
	dest := pack.PrefixSection(cardNo, 6)
	if dest == "" {
		err = errors.New("请确认输入正确的卡号")
		return
	}
	bankInfoSQL := `SELECT bb.card_type, b.code bank_code, b.name bank_name, b.icon_path
					FROM tools.bank_bins bb
					LEFT JOIN tools.banks b ON b.code = bb.code
					WHERE bin_code = ? LIMIT 1`
	var r CardInfo
	ct.d.Raw(bankInfoSQL, dest).Scan(&r)
	result = &r
	return
}

//
func (ct *Controller) CardNoValidBusiness(cardNo string) bool {
	return pack.LuHn(cardNo)
}
