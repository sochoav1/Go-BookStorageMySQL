package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Connect() {
	dsn := "sochoav8a:PrintPassword321_@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	DB = database
}

func GetDb() *gorm.DB {
	return DB
}
