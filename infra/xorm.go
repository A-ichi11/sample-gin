package infra

import (
	"log"

	"github.com/go-xorm/xorm"
)

func DBInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return engine
}
