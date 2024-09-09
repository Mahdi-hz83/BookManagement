package APIpackage

import (
	"New_Book_Management/Models"
	"New_Book_Management/Responds"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BookManagementAPI struct {
	DB *gorm.DB
}

func (bm *BookManagementAPI) RetrieveAllBooks() ([]Models.Book, error) {
	var books []Models.Book
	result := bm.DB.Find(&books)
	return books, result.Error
}

func (bm *BookManagementAPI) AddBook(book Models.Book) error {
	result := bm.DB.Create(&book)
	return result.Error
}

func (bm *BookManagementAPI) DeleteBook(isbn string) (int64, error) {
	result := bm.DB.Delete(&Models.Book{}, "isbn = ?", isbn)
	return result.RowsAffected, result.Error
}

func (bm *BookManagementAPI) UpdateBook(isbn string, UpdatedBook Models.Book) (int64, error) {
	var book Models.Book
	result := bm.DB.First(&book, "isbn = ?", isbn)
	if result.Error != nil {
		return 0, result.Error
	}

	book.Title = UpdatedBook.Title
	book.Price = UpdatedBook.Price

	result = bm.DB.Save(&book)
	return result.RowsAffected, result.Error
}

func (bm *BookManagementAPI) RetrieveBookByISBN(isbn string) (Models.Book, error) {
	var book Models.Book
	result := bm.DB.First(&book, "isbn = ?", isbn)
	return book, result.Error
}

func SetupRouter(apiBookManager *BookManagementAPI) *gin.Engine {
	router := gin.Default()
	router.GET("/books", func(c *gin.Context) {
		books, err := apiBookManager.RetrieveAllBooks()
		if err != nil {
			Responds.RespondWithInternalServerError(c, err.Error())
		}
		c.JSON(http.StatusOK, books)
	})

	router.GET("/books/:isbn", func(c *gin.Context) {
		isbn := gettingISBN(c)
		book, err := apiBookManager.RetrieveBookByISBN(isbn)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				Responds.RespondWithNotFound(c, "Book not found")
			} else {
				Responds.RespondWithInternalServerError(c, err.Error())
			}
			return
		}
		Responds.RespondWithReturningData(c, book)
	})

	router.POST("/books", func(c *gin.Context) {
		var book Models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			Responds.RespondWithBadRequest(c, err.Error())
		}
		if err := apiBookManager.AddBook(book); err != nil {
			Responds.RespondWithInternalServerError(c, err.Error())
		}
		Responds.RespondWithReturningData(c, book)
	})

	router.DELETE("/books/:isbn", func(c *gin.Context) {
		isbn := gettingISBN(c)
		rowsAffected, err := apiBookManager.DeleteBook(isbn)
		if err != nil {
			Responds.RespondWithInternalServerError(c, err.Error())
			return
		}
		if rowsAffected == 0 {
			Responds.RespondWithNotFound(c, "Book not found")
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	})

	router.PUT("/books/:isbn", func(c *gin.Context) {
		isbn := gettingISBN(c)
		var updatedBook Models.Book
		err := c.ShouldBindJSON(&updatedBook)
		if err != nil {
			Responds.RespondWithBadRequest(c, err.Error())
			return
		}

		rowsAffected, err := apiBookManager.UpdateBook(isbn, updatedBook)
		if err != nil {
			Responds.RespondWithInternalServerError(c, err.Error())
		}

		if rowsAffected == 0 {
			Responds.RespondWithNotFound(c, "Book not found")
			return
		}
		Responds.RespondWithReturningData(c, updatedBook)
	})

	return router
}

func gettingISBN(c *gin.Context) string {
	isbn := c.Param("isbn")
	return isbn
}
