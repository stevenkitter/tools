package models

import "time"

// Base b
type Base struct {
	ID        string     `gorm:"primary_key" sql:"comment:'主键'" json:"id"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP(3);not null;type:timestamp(3)" sql:"comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null" sql:"comment:'修改时间'" json:"updated_at"`
	DeletedAt *time.Time `sql:"index;comment:'软删除时间'" json:"-"`
}
