package main

import (
	"log"
	"web-api/book"
	"web-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/webapidb?charset=utf8mb4&parseTime=True&loc=Local" // perhatikan nama db & root pass db disini
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)	// 2:25:11 
	

	r := gin.Default()
	v1 := r.Group("/v1")	
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	r.Run() 
}

// ==========
//  Tutorial Golang Web API Bahasa Indonesia - Full Course
// https://www.youtube.com/watch?v=GjI0GSvmcSU&t=3261s
// ==========

//goal layer---
// main
// handler/controller
// service
// repository
// db
// mysql