package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// ConnectMysqlDB c
func ConnectMysqlDB(user, password, dbPath, database string) (*gorm.DB, error) {
	sqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", user, password, dbPath, database)
	return gorm.Open("mysql", sqlURL)
}
