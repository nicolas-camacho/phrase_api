package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-camacho/phrase_api/pkg/phrases"
)

//PhrasesRouter is use to set the routes of the phrase service
func PhrasesRouter(app fiber.Router, service phrases.PhraseService) {
	app.Post("/phrases", addPhrase(service))
}

func addPhrase(service phrases.PhraseService) fiber.Handler {
	return func(context *fiber.Ctx) error {
		var requestBody phrases.Phrase
		if err := context.BodyParser(&requestBody); err != nil {
			return context.Status(fiber.StatusBadRequest).SendString("Error with the request")
		}
		response, dberr := service.CreatePhrase(&requestBody)
		if dberr != nil {
			return context.Status(fiber.StatusInternalServerError).SendString("Error while inserting the phrase")
		}
		return context.Status(fiber.StatusCreated).JSON(response)
	}
}
