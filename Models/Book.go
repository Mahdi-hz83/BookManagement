package Models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ISBN  uint `gorm:"primaryKey"`
	Title string
	Price float32
}

type bookManager interface {
	RetrieveBooks() ([]Book, error)
	AddBook(book Book) error
	DeleteBook(isbn string) (int64, error)
}
