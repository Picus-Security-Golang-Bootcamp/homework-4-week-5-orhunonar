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

func Hello(c *fiber.Ctx) {
	c.Send("Hello World!")
}
func setupRoutes(app *fiber.App) {
	app.Get("/", Hello)

	app.Get("/api/v1/book", book.Getbooks)
	app.Get("/api/v1/book/:id", book.Getbook)
	app.Post("/api/v1/book", book.Newbook)
	app.Delete("/api/v1/book/:id", book.Deletebook)

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
