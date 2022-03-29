package main

import (
	"fmt"
	"library-api/handler"
	"library-api/model"
	"library-api/repository"
	"library-api/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	dsn := "host=localhost user=postgres password=postgres dbname=book port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(model.Book{})

	bookRepository := repository.NewRepository(db)
	bookService := service.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	books, err := bookService.FindAll()

	if err != nil {
		log.Fatal(err.Error())
		return
	}
	for _, e := range books {
		fmt.Println(e)
	}

	v1 := router.Group("/v1")

	router.GET("/", bookHandler.HelloHandler)

	v1.GET("/books", bookHandler.BookHandler)
	v1.GET("/books/:id", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.CreateBookHandler)

	router.Run(":8081")
}
