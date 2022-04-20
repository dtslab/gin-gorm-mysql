package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()	
	r.GET("/", rootHandler)
	r.GET("/hello", helloHandler)
	r.GET("/books/:id/:title", booksHandler)
	r.GET("/query", queryHandler)
	r.POST("/books", postBooksHandler)

	r.Run() 
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "Dedy Tri Samudra",
		"bio": "I'm a software engineer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"content": "Hello World",
		"subtitle": "belajar gin framework",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(200, gin.H{
		"id": id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(200, gin.H{
		"title": title,
		"price": price,
	})
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context) { 
	var bookInput BookInput
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