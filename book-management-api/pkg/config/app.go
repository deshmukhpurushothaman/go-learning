package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:deshmukh@/book-management-api?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
