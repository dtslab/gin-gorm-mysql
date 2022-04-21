package handler

import (
	"fmt"
	"net/http"
	"web-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "Dedy Tri Samudra",
		"bio": "I'm a software engineer",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"content": "Hello World",
		"subtitle": "belajar gin framework",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(200, gin.H{
		"id": id,
		"title": title,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"title": title,
		"price": price,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) { 
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)	
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

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"data": book,
	})
}