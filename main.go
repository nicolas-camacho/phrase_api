package main

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	api "github.com/nicolas-camacho/phrase_api/api"
	database "github.com/nicolas-camacho/phrase_api/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			if _, ok := err.(*fiber.Error); ok {
				return errors.New("this is a fiber error")
			}

			return errors.New("this is a managed error")
		},
		Views: engine,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Static("/static", "./static")

	api.NewRouter(app, db)

	app.Listen(":" + os.Getenv("PORT"))
}
