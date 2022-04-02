package main

import (
	"fmt"
	"homework4/book"
	"homework4/database"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func main() {

	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) {
	// 	c.Send("Hello World!")
	// })

	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

}

func Greeting(c *fiber.Ctx) {
	c.Send("Welcome to the Book Store")
}
func setupRoutes(app *fiber.App) {
	app.Get("/", Greeting)

	app.Get("/book", book.Getbooks)
	app.Get("/book/:id", book.Getbook)
	app.Post("/book", book.Newbooks)
	app.Post("/books", book.InsertBook)
	app.Get("/books", book.Getbooks)
	app.Delete("/book/:id", book.Deletebook)

}
func initDatabase() {

	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}
