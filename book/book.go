package book

import (
	"homework4/database"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func Getbooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func Getbook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	c.JSON(book)

}

func Newbook(c *fiber.Ctx) {
	listofbooks := []Book{
		{
			Title:      "The Lord of the Rings",
			PageNumber: 1234,
			Stock:      10,
			Price:      12.99,
			Author:     "J.R.R. Tolkien",
		},
		{
			Title:      "The Hobbit",
			PageNumber: 1234,
			Stock:      10,
			Price:      12.99,
			Author:     "J.R.R. Tolkien",
		},
		{
			Title:      "The Cat in the Hat",
			PageNumber: 1234,
			Stock:      10,
			Price:      12.99,
			Author:     "Dr. Seuss",
		},
	}

	db := database.DBConn
	db.Create(&listofbooks)
	c.Send("Book Successfully created")
	c.JSON(listofbooks)

}

func Deletebook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}

type Book struct {
	gorm.Model
	Title      string  `json:"title"`
	PageNumber int     `json:"page_number"`
	Stock      int     `json:"stock"`
	Price      float64 `json:"price"`
	Author     string  `json:"author"`
}
