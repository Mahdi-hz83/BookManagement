package API

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

func (bm *BookManagementAPI) RetrieveBooks() ([]Models.Book, error) {
	var books []Models.Book
	result := bm.DB.Find(&books)
	return books, result.Error
}

func (bm *BookManagementAPI) AddBook(book Models.Book) error {
	result := bm.DB.Create(&book)
	return result.Error
}

func (bm *BookManagementAPI) DeleteBook(isbn string) error {
	result := bm.DB.Delete(&Models.Book{}, "isbn = ?", isbn)
	return result.Error
}

func SetupRouter(apiBookManager *BookManagementAPI) *gin.Engine {
	router := gin.Default()
	router.GET("/books", func(c *gin.Context) {
		books, err := apiBookManager.RetrieveBooks()
		if err != nil {
			Responds.RespondWithInternalServerError(c, err.Error())
		}
		c.JSON(http.StatusOK, books)
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
		id := c.Param("isbn")
		if err := apiBookManager.DeleteBook(id); err != nil {
			Responds.RespondWithInternalServerError(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	})

	return router
}
