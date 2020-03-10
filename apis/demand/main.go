package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stevenkitter/tools/apis/demand/controller"
)

// MYSQL_PWD MYSQL_HOST REDIS_HOST
const (
	Port = ":8100"
)

func main() {
	ct := controller.NewController()
	if err := ct.Run(Port); err != nil {
		panic(err)
	}
}
