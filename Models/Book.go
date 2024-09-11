package Models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ISBN   uint `gorm:"primaryKey"`
	Title  string
	Price  float32
	Status string
}

type bookManager interface {
	RetrieveAllBooks() ([]Book, error)
	RetrieveBookByISBN(isbn string) (Book, error)
	AddBook(book Book) error
	DeleteBook(isbn string) (int64, error)
	UpdateBook(isbn string, book Book) (int64, error)
}
