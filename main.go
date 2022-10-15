package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {

	db := initializeDb()
	app := fiber.New()
	configureRoutes(app, db)
	serve(app)

}

func initializeDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:O_-rlzxrxP!qjQv2@tcp(localhost:3306)/zocman"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Post{})

	return db
}

func configureRoutes(app *fiber.App, db *gorm.DB) {

	app.Get("/api/posts/1", func(c *fiber.Ctx) error {
		var post Post

		db.First(&post, 1)

		return c.JSON(post)

		//return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		var posts []Post

		db.Find(&posts)

		return c.JSON(posts)

	})

	app.Post("/api/posts", func(c *fiber.Ctx) error {

		var post Post

		if err := c.BodyParser(&post); err != nil {
			return err
		}

		db.Create(&post)
		//db.Select("Title", "Description").Create(&post)
		// url := fmt.Sprintf("/posts/%d", post.Id)
		// return c.Redirect(url)

		//return c.SendString("Hello, World 👋!")

		return c.JSON("OK")
	})
}

func serve(app *fiber.App) {
	app.Listen("localhost:3000")
}
