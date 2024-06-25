package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sochoav1/Go-BookStorageMySQL/pkg/config"
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
func (b *Book) CreateBook() *Book {
	DB.NewRecord(b)
	DB.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	DB.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := DB.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	DB.Where("ID = ?", id).Delete(book)
	return book
}
