package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// ConnectMysqlDB
func ConnectMysqlDB(user, password, dbPath, database string) (*gorm.DB, error) {
	sqlUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", user, password, dbPath, database)
	return gorm.Open("mysql", sqlUrl)
}
