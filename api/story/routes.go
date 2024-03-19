package storyRoutes

import (
	"github.com/gofiber/fiber/v2"
	StoryController "github.com/puvadon-artmit/gofiber-template/api/story/controllers"
)

func SetupStoryRoutes(router fiber.Router) {
	app := router.Group("story")
	app.Get("/get-story", StoryController.GetAll)
	app.Get("/get-by-id/:story_id", StoryController.GetById)
	app.Post("/create-story", StoryController.Create)
}
