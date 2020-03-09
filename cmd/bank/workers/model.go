package workers

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"time"
)

type Base struct {
	ID        string     `gorm:"primary_key" sql:"comment:'主键'"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP(3);not null;type:timestamp(3)" sql:"comment:'创建时间'"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null" sql:"comment:'修改时间'"`
	DeletedAt *time.Time `sql:"index;comment:'软删除时间'"`
}

type Bank struct {
	Base

	Code     string `sql:"comment:'缩写'"`
	Name     string `sql:"comment:'名称'"`
	IconPath string `sql:"comment:'头像地址'"`
	CardType string `sql:"comment:'储蓄卡'"`
}

func (b *Bank) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}

type BankBin struct {
	Base
	Code    string `sql:"comment:'银行缩写'"`
	BinCode string `sql:"comment:'bin码'"`
}

func (b *BankBin) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}

type ErrorBankBin struct {
	Base
	BinCode string `sql:"comment:'bin码'"`
}

func (b *ErrorBankBin) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}
