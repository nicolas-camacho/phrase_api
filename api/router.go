package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-camacho/phrase_api/api/routes"
	"github.com/nicolas-camacho/phrase_api/pkg/phrases"
	"gorm.io/gorm"
)

//NewRouter is used to group every service from inside the api route
func NewRouter(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	webView := app.Group("/")

	phrasesRepo := phrases.NewRepository(db)
	phraseService := phrases.NewService(phrasesRepo)

	routes.PhrasesRouter(api, phraseService)
	routes.WebRouter(webView, phraseService)
}
