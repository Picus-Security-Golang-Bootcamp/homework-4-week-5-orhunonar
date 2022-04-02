package main

import ( //Import packages
	"fmt"                //Println
	"homework4/book"     //Book package
	"homework4/database" //Database package

	"github.com/gofiber/fiber" //Fiber package
	"gorm.io/driver/sqlite"    //SQLite package
	"gorm.io/gorm"             //GORM package
)

func main() { //Main function

	app := fiber.New() //Create new app

	initDatabase() //Initialize database

	setupRoutes(app) //Setup routes
	app.Listen(3000) //Listen on port 3000

}

func Greeting(c *fiber.Ctx) { //Greeting route
	c.Send("Welcome to the Book Store") //Send welcome message
}
func setupRoutes(app *fiber.App) { //Setup routes
	app.Get("/", Greeting) //Greeting route

	app.Get("/book", book.Getbooks)    // Get all books
	app.Get("/book/:id", book.Getbook) // Get book with id
	app.Post("/book", book.Newbooks)   // Add spesific books

	app.Post("/insert", book.InsertBook) //Insert Book using JSON

	app.Post("/buy/:id", book.BuyBook) //Buy Book using JSON

	app.Delete("/book/:id", book.Deletebook) //Delete book which id is given

}
func initDatabase() { //Initialize database

	var err error                                                            //Error variable
	database.DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) //Open database
	if err != nil {                                                          //Check for error
		panic("failed to connect database") //If error panic
	}
	fmt.Println("Connection Opened to Database") //Print connection opened
	database.DBConn.AutoMigrate(&book.Book{})    //Migrate database
	fmt.Println("Database Migrated")             //Print database migrated
}
