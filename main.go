package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	database "github.com/nicolas-camacho/phrase_api/database"
)

func main() {
	_, err := database.Connect()
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			if _, ok := err.(*fiber.Error); ok {
				return errors.New("this is a fiber error")
			}

			return errors.New("this is a managed error")
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Welcome to Phrase API")
	})

	app.Listen(":3000")
}
