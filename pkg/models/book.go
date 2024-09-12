package models

import (
	"github.com/NaveenCJoy/book-store-go/pkg/config"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// db.NewRecord(b) //checks if the Book instance b is new. In gorm v2, NewRecord is not needed. gorm automatically checks this.
	db := config.GetDB()
	db.Create(&b) //inserts the new Book to the db ie, creates a new row in books table
	return b
}

func GetAllBooks() []Book {
	var Books []Book //create an empty slice of Books structs
	db := config.GetDB()
	db.Find(&Books) // find all records from books table and populate Books variable
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := config.GetDB()
	db = db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db := config.GetDB()
	db.Where("ID=?", ID).Delete(&book)
	return book
}
