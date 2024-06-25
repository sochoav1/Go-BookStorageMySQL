package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sochoav1/Go-BookStorageMySQL/pkg/config"
	"github.com/sochoav1/bookstore/pkg/config"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Autor       string `json:"autor"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	DB = config.DB
	DB.AutoMigrate(&Book{})
}
