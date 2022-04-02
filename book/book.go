package book // Package book

import ( //Import packages
	"homework4/database" //Database package

	"github.com/gofiber/fiber" //Fiber package
	"gorm.io/gorm"             //GORM package
)

func Getbooks(c *fiber.Ctx) { // Get all books
	db := database.DBConn //Get database connection
	var books []Book      //Create new list of books
	db.Find(&books)       //Find all books
	c.JSON(books)         //Return list of books
}

func Getbook(c *fiber.Ctx) { // Get book with id
	id := c.Params("id")  //Get id
	db := database.DBConn //Get database connection
	var book Book         //Create new book
	db.First(&book, id)   //Find book with id
	c.JSON(book)          //Return book

}

func Newbooks(c *fiber.Ctx) { // Add spesific books
	listofbooks := []Book{ //Create new list of books
		{ //Book 1
			Title:      "The Lord of the Rings", //Title
			PageNumber: 1234,                    //Page number
			Stock:      10,                      //Stock
			Price:      12.99,                   //Price
			Author:     "J.R.R. Tolkien",        //Author
		},
		{ //Book 2
			Title:      "The Hobbit",     //Title
			PageNumber: 1234,             //Page number
			Stock:      10,               //Stock
			Price:      12.99,            //Price
			Author:     "J.R.R. Tolkien", //Author
		},
		{ //Book 3
			Title:      "The Cat in the Hat", //Title
			PageNumber: 1234,                 //Page number
			Stock:      10,                   //Stock
			Price:      12.99,                //Price
			Author:     "Dr. Seuss",          //Author
		},
	}

	db := database.DBConn               //Get database connection
	db.Create(&listofbooks)             //Create books
	c.Send("Book Successfully created") //Return success
	c.JSON(listofbooks)                 //Return list of books

}

func Deletebook(c *fiber.Ctx) { //Delete book whic id is given
	id := c.Params("id")  //Get id
	db := database.DBConn //Get database connection

	var book Book         //Create new book
	db.First(&book, id)   //Find book with id
	if book.Title == "" { //Check if book is found
		c.Status(500).Send("No Book Found with ID") //Return error
		return                                      //Return
	}
	db.Delete(&book)                    //Delete book
	c.Send("Book Successfully deleted") //Return success
}
func BuyBook(c *fiber.Ctx) { // Buy book with id
	id := c.Params("id")  //Get id
	db := database.DBConn //Get database connection
	var book Book         //Create new book
	db.First(&book, id)   //Find book with id
	if book.Title == "" { //Check if book is found
		c.Status(500).Send("No Book Found with ID") //Return error
		return                                      //Return
	}
	if book.Stock == 0 { //Check if book is in stock
		c.Status(500).Send("No Stock") //Return error
		return                         //Return
	}
	book.Stock = book.Stock - 1        //Decrease stock
	db.Save(&book)                     //Save book
	c.Send("Book Successfully bought") //Return success
}
func InsertBook(c *fiber.Ctx) { // Add custom books
	var book Book              //Create new book
	err := c.BodyParser(&book) //Parse body
	if err != nil {            //Check if error
		c.Status(500).Send("Error in parsing the body") //Return error
		return                                          //Return
	}

	db := database.DBConn               //Get database connection
	db.Create(&book)                    //Create book
	c.Send("Book Successfully created") //Return success
	c.JSON(book)                        //Return book
}

type Book struct { //Book struct
	gorm.Model         //Model
	Title      string  `json:"title"`       //Title
	PageNumber int     `json:"page_number"` //Page number
	Stock      int     `json:"stock"`       //Stock
	Price      float64 `json:"price"`       //Price
	Author     string  `json:"author"`      //Author
}
