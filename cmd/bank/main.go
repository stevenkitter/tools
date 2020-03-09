package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stevenkitter/tools/cmd/bank/workers"
)

func main() {
	workers.Run()
}
