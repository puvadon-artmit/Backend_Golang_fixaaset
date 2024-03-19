package Category_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerStoryRoutes "github.com/puvadon-artmit/gofiber-template/api/category_story/controllers"
)

func SetupCategory_StoryRoutes(router fiber.Router) {
	app := router.Group("category-story")
	app.Get("/get-category-story", ControllerStoryRoutes.GetAllHandler)
	app.Get("/get-by-id/:category_story_id", ControllerStoryRoutes.GetById)
	app.Post("/create-category-story", ControllerStoryRoutes.Create)
	app.Patch("/update-category-story/:category_story_id", ControllerStoryRoutes.UpdateCategory_Story)
	app.Delete("/delete-category-story/:controller_id", ControllerStoryRoutes.DeleteCategory_Story)
}
