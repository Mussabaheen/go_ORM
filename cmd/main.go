package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mussabaheen/orm/pkg/bookstore"
	"github.com/mussabaheen/orm/pkg/db"
)

func main() {
	fmt.Println("GO ORM")
	r := gin.Default()
	defer r.Run()

	mongoClient := db.NewMongoClient()
	bookRepository := bookstore.NewRepository(mongoClient)
	bookService := bookstore.NewService(bookRepository)
	controller := bookstore.NewController(bookService)
	controller.RegisterRoutes(r)
}
