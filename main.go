package main

import (
	"book-library/api/handlers"

	"github.com/labstack/echo"
)

// @title Book Library CRUD APP
// @version 1.0.0
// @host localhost:1323
// @BasePath /books

func main() {
	e := echo.New()
    e.GET("/books", handlers.GetAllBooks)
	e.GET("/books/:id", handlers.GetBookById)
	e.POST("/books", handlers.StoreBooks)
	e.PATCH("/books/checkout", handlers.CheckoutBook)
	e.PATCH("/books/checkin", handlers.CheckinBook)
	e.DELETE("/books/:id", handlers.DeleteBook)
    e.Logger.Fatal(e.Start(":1323"))
} 