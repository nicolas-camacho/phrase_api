package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-camacho/phrase_api/pkg/phrases"
)

//PhrasesRouter is use to set the routes of the phrase service
func PhrasesRouter(app fiber.Router, service phrases.PhraseService) {
	phrasesRouter := app.Group("/phrases")

	phrasesRouter.Post("/", addPhrase(service))
	phrasesRouter.Get("/", getPhrases(service))
	phrasesRouter.Get("/random", getRandomPhrase(service))
	phrasesRouter.Get("/:id", getPhrase(service))
}

func addPhrase(service phrases.PhraseService) fiber.Handler {
	return func(context *fiber.Ctx) error {
		var requestBody phrases.Phrase
		if err := context.BodyParser(&requestBody); err != nil {
			return context.Status(fiber.StatusBadRequest).SendString("Error with the request")
		}
		response, dberr := service.CreatePhrase(&requestBody)
		if dberr != nil {
			return context.Status(fiber.StatusInternalServerError).SendString(dberr.Error())
		}
		return context.Status(fiber.StatusCreated).JSON(response)
	}
}

func getPhrases(service phrases.PhraseService) fiber.Handler {
	return func(context *fiber.Ctx) error {
		response, dberr := service.ReadPhrases()
		if dberr != nil {
			return context.Status(fiber.StatusInternalServerError).SendString(dberr.Error())
		}
		return context.Status(fiber.StatusOK).JSON(response)
	}
}

func getPhrase(service phrases.PhraseService) fiber.Handler {
	return func(context *fiber.Ctx) error {
		requestedID, err := strconv.Atoi(context.Params("id"))
		if err != nil {
			return context.Status(fiber.StatusBadRequest).SendString("Invalid Id")
		}
		response, dberr := service.ReadPhrase(requestedID)
		if dberr != nil {
			return context.Status(fiber.StatusInternalServerError).SendString(dberr.Error())
		}
		return context.Status(fiber.StatusOK).JSON(response)
	}
}

func getRandomPhrase(service phrases.PhraseService) fiber.Handler {
	return func(context *fiber.Ctx) error {
		response, dberr := service.ObtainPhrase()
		if dberr != nil {
			return context.Status(fiber.StatusInternalServerError).SendString(dberr.Error())
		}
		return context.Status(fiber.StatusOK).JSON(response)
	}
}
