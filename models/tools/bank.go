package tools

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/stevenkitter/tools/models"
)

type Bank struct {
	models.Base

	Code     string `sql:"comment:'缩写'" json:"code"`
	Name     string `sql:"comment:'名称'" json:"name"`
	IconPath string `sql:"comment:'头像地址'" gorm:"size:1000" json:"iconPath"`
}

func (b *Bank) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}

type BankBin struct {
	models.Base
	Code     string `sql:"comment:'银行缩写'"`
	BinCode  string `sql:"comment:'bin码'"`
	CardType string `sql:"comment:'储蓄卡'"`
}

func (b *BankBin) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}

type ErrorBankBin struct {
	models.Base
	BinCode string `sql:"comment:'bin码'"`
}

func (b *ErrorBankBin) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}
