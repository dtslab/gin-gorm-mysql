package main

import (
	"fmt"
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
	// ------CREATE (CRUD)------
	// https://gorm.io/docs/index.html

	// book := book.Book{}
	// book.Title = "Belajar Tatabog"
	// book.Price = 60000
	// book.Discount = 10
	// book.Rating = 4
	// book.Description = "buku belajar masak"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error creating book record")
	// 	fmt.Println("==========================")
	// }

	// ------READ (CRUD)------
	// https://gorm.io/docs/index.html

	// var book []book.Book

	// err = db.Debug().Where("rating = ?", 4).Find(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error finding book record")
	// 	fmt.Println("==========================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title : ", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	// ------UPDATE (CRUD)------
	// https://gorm.io/docs/index.html

	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).Find(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error Update book record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Belajar Tataboga (Revisi)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error Update book record")
	// 	fmt.Println("==========================")
	// }

	// ------DELETE (CRUD)------
	// https://gorm.io/docs/index.html

	var book book.Book

	err = db.Debug().Where("id = ?", 1).Find(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("error Deleting book record")
		fmt.Println("==========================")
	}

	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("error deleting book record")
		fmt.Println("==========================")
	}

	r := gin.Default()
	v1 := r.Group("/v1")	
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	r.Run() 
}

// ==========
//  Tutorial Golang Web API Bahasa Indonesia - Full Course
// https://www.youtube.com/watch?v=GjI0GSvmcSU&t=3261s
// ==========