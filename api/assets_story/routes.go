package Assets_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	Assets_StoryController "github.com/puvadon-artmit/gofiber-template/api/assets_story/controllers"
)

func SetupAssets_StoryRoutes(router fiber.Router) {
	app := router.Group("assets_story")
	app.Get("/get-assets_story", Assets_StoryController.GetAll)
	app.Get("/get-by-id/:assets_story_id", Assets_StoryController.GetById)
	app.Post("/create-assets_story", Assets_StoryController.Create)
}
