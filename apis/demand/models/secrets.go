package models

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
	"github.com/stevenkitter/tools/models"
)

// Secrets s
type Secrets struct {
	models.Base

	AppID     string `json:"用户识别号"`
	AppSecret string `json:"用户密码"`
}

// BeforeCreate hook
func (s *Secrets) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", ksuid.New().String())
}
