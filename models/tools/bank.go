package tools

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/stevenkitter/tools/models"
)
// Bank b
type Bank struct {
	models.Base

	Code     string `sql:"comment:'缩写'" json:"code"`
	Name     string `sql:"comment:'名称'" json:"name"`
	IconPath string `sql:"comment:'头像地址'" gorm:"size:1000" json:"iconPath"`
}
// BeforeCreate hook
func (b *Bank) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}
// BankBin b
type BankBin struct {
	models.Base
	Code     string `sql:"comment:'银行缩写'"`
	BinCode  string `sql:"comment:'bin码'"`
	CardType string `sql:"comment:'储蓄卡'"`
}
// BeforeCreate hook
func (b *BankBin) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}
// ErrorBankBin e
type ErrorBankBin struct {
	models.Base
	BinCode string `sql:"comment:'bin码'"`
}
// BeforeCreate hook
func (b *ErrorBankBin) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}
