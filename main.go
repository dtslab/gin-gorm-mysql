package main

import (
	"fmt"
	"log"
	"web-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/webapidb?charset=utf8mb4&parseTime=True&loc=Local" // perhatikan nama db & root pass db disini
  _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}

	fmt.Println("Connected to database successfully")

	r := gin.Default()
	v1 := r.Group("/v1")	
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	r.Run() 
}

// 1:23:04 https://www.youtube.com/watch?v=GjI0GSvmcSU&t=3261s
