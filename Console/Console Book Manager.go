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

func (bm *BookManagementConsole) RetrieveAllBooks() ([]Models.Book, error) {
	var books []Models.Book
	result := bm.DB.Find(&books)
	return books, result.Error
}

func (bm *BookManagementConsole) RetrieveBookByISBN(isbn string) (Models.Book, error) {
	var book Models.Book
	result := bm.DB.First(&book, "isbn = ?", isbn)
	return book, result.Error
}

func (bm *BookManagementConsole) AddBook(book Models.Book) error {
	result := bm.DB.Create(&book)
	return result.Error
}

func (bm *BookManagementConsole) DeleteBook(isbn string) (int64, error) {
	result := bm.DB.Delete(&Models.Book{}, "isbn = ?", isbn)
	return result.RowsAffected, result.Error
}

func (bm *BookManagementConsole) UpdateBook(isbn string, updatedBook Models.Book) (int64, error) {
	var book Models.Book
	result := bm.DB.First(&book, "isbn = ?", isbn)
	if result.Error != nil {
		return 0, result.Error
	}

	book.Title = updatedBook.Title
	book.Price = updatedBook.Price
	book.Status = updatedBook.Status

	result = bm.DB.Save(&book)
	return result.RowsAffected, result.Error
}

func (bm *BookManagementConsole) ShowBooks() {
	books, err := bm.RetrieveAllBooks()
	if err != nil {
		fmt.Println("Error retrieving books:", err)
		return
	}

	if len(books) == 0 {
		log.Println("No books found.")
		return
	}

	for _, book := range books {
		log.Printf("ISBN: %d, Title: %s, Price: %f , Status: %s\n", book.ISBN, book.Title, book.Price, book.Status)
	}
}

func (bm *BookManagementConsole) ShowBookByISBN() {
	fmt.Print("Enter ISBN of the book to retrieve: ")
	var isbn string
	fmt.Scan(&isbn)

	book, err := bm.RetrieveBookByISBN(isbn)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Book not found.")
		} else {
			fmt.Println("Error retrieving book:", err)
		}
		return
	}

	fmt.Printf("ISBN: %d, Title: %s, Price: %f , Status: %s\n", book.ISBN, book.Title, book.Price, book.Status)
}

func (bm *BookManagementConsole) AddNewBook() {
	var book Models.Book

	fmt.Print("Enter ISBN: ")
	fmt.Scan(&book.ISBN)
	fmt.Print("Enter Title: ")
	fmt.Scan(&book.Title)
	fmt.Print("Enter Price: ")
	fmt.Scan(&book.Price)
	fmt.Print("Enter Status: ")
	fmt.Scan(&book.Status)

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

func (bm *BookManagementConsole) UpdateBookDetails() {
	fmt.Print("Enter ISBN of the book to update: ")
	var isbn string
	fmt.Scan(&isbn)

	var updatedBook Models.Book
	fmt.Print("Enter new Title: ")
	fmt.Scan(&updatedBook.Title)
	fmt.Print("Enter new Price: ")
	fmt.Scan(&updatedBook.Price)
	fmt.Print("Enter New Status: ")
	fmt.Scan(&updatedBook.Status)

	rowsAffected, err := bm.UpdateBook(isbn, updatedBook)
	if err != nil {
		fmt.Println("Error updating book:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("Book not found.")
	} else {
		fmt.Println("Book updated successfully.")
	}
}
