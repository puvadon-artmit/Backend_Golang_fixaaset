package Typeplan_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan_Story "github.com/puvadon-artmit/gofiber-template/api/typeplan_story/controllers"
)

func SetupTypeplan_StoryRoutes(router fiber.Router) {
	app := router.Group("typeplan_story")
	app.Get("/get-typeplan-story", ControllerTypeplan_Story.GetAllHandler)
	app.Get("/get-by-id/:typeplan_story_id", ControllerTypeplan_Story.GetById)
	app.Post("/create-typeplan-story", ControllerTypeplan_Story.Create)
	app.Patch("/update-typeplan-story/:typeplan_story_id", ControllerTypeplan_Story.UpdateTypeplan_Story)
	app.Delete("/delete-typeplan-story/:typeplan_story_id", ControllerTypeplan_Story.DeleteTypeplan_Story)
}
