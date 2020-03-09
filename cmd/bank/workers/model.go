package workers

import "time"

type Bank struct {
	ID        string     `gorm:"primary_key" sql:"comment:'主键'"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP(3);not null;type:timestamp(3)" sql:"comment:'创建时间'"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null" sql:"comment:'修改时间'"`
	DeletedAt *time.Time `sql:"index;comment:'软删除时间'"`

	Code     string `sql:"comment:'缩写'"`
	Name     string `sql:"comment:'名称'"`
	IconPath string `sql:"comment:'头像地址'"`
	CardType string `sql:"comment:'储蓄卡'"`
}
