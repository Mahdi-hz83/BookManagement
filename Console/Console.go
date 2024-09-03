package Console

import (
	"New_Book_Management/Models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type BookManagementConsole struct {
	DB *gorm.DB
}

func (bm *BookManagementConsole) RetrieveBooks() ([]Models.Book, error) {
	var books []Models.Book
	result := bm.DB.Find(&books)
	return books, result.Error
}

func (bm *BookManagementConsole) AddBook(book Models.Book) error {
	result := bm.DB.Create(&book)
	return result.Error
}

func (bm *BookManagementConsole) DeleteBook(isbn string) (int64, error) {
	result := bm.DB.Delete(&Models.Book{}, "isbn = ?", isbn)
	return result.RowsAffected, result.Error
}

func (bm *BookManagementConsole) ShowBooks() {
	books, err := bm.RetrieveBooks()
	if err != nil {
		fmt.Println("Error retrieving books:", err)
		return
	}

	if len(books) == 0 {
		log.Println("No books found.")
		return
	}

	for _, book := range books {
		log.Printf("ISBN: %d, Title: %s, Price: %f\n", book.ISBN, book.Title, book.Price)
	}
}

func (bm *BookManagementConsole) AddNewBook() {
	var book Models.Book

	fmt.Print("Enter ISBN: ")
	fmt.Scan(&book.ISBN)
	fmt.Print("Enter Title: ")
	fmt.Scan(&book.Title)
	fmt.Print("Enter Price: ")
	fmt.Scan(&book.Price)

	if err := bm.AddBook(book); err != nil {
		fmt.Println("Error adding book:", err)
		return
	}

	fmt.Println("Book added successfully!")
}

func (bm *BookManagementConsole) RemoveBook() {
	fmt.Print("Enter ISBN of the book to delete: ")
	var isbn string
	fmt.Scan(&isbn)

	rowsAffected, err := bm.DeleteBook(isbn)
	if err != nil {
		fmt.Println("Error deleting book:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("Book not found.")
	} else {
		fmt.Println("Book deleted successfully.")
	}
}
