package Responds

import (
	"New_Book_Management/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondWithBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"Error": message})
}

func RespondWithNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"Error": message})
}

func RespondWithInternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"Error": message})
}

func RespondWithReturningData(c *gin.Context, data Models.Book) {
	c.JSON(http.StatusOK, gin.H{"Data": data})
}
