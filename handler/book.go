package handler

import (
	"fmt"
	"net/http"
	"web-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "Dedy Tri Samudra",
		"bio": "I'm a software engineer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"content": "Hello World",
		"subtitle": "belajar gin framework",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(200, gin.H{
		"id": id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBooksHandler(c *gin.Context) { 
	var bookInput book.BookInput
	err := c.ShouldBindJSON(&bookInput)	
	if err != nil {		
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return					
	}
	c.JSON(200, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}