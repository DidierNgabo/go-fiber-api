package main

import (
	"fmt"

	"github.com/DidierNgabo/go-fiber-api/book"
	"github.com/DidierNgabo/go-fiber-api/db"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/book", book.GetBooks)
	app.Get("api/v1/book/:id", book.GetBook)
	app.Post("api/v1/book", book.NewBook)
	app.Delete("api/v1/book/:id", book.DeleteBook)
}

func initDb() {
	var err error
	db.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect to db")
	}

	fmt.Println("database successfully connected")

	db.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}
func main() {
	app := fiber.New()
	initDb()
	defer db.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}
