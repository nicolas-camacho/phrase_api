package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-camacho/phrase_api/pkg/phrases"
)

//WebRouter is use to set the routes of the web service.
func WebRouter(app fiber.Router, service phrases.PhraseService) {
	app.Get("/", handleIndex(service))
}

func handleIndex(service phrases.PhraseService) fiber.Handler {
	return func(context *fiber.Ctx) error {
		randomPhrase, err := service.ObtainPhrase()
		if err != nil {
			return context.Render("index", fiber.Map{
				"Phrase": nil,
				"Name":   nil,
				"Id":     nil,
				"Error":  err,
			})
		}
		return context.Render("index", fiber.Map{
			"Phrase": randomPhrase.Content,
			"Name":   os.Getenv("FRIEND_NAME"),
			"Id":     randomPhrase.ID,
			"Error":  false,
		})
	}
}
