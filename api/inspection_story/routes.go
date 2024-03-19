package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerInspection_Story "github.com/puvadon-artmit/gofiber-template/api/inspection_story/controllers"
)

func SetupInspection_StoryRoutes(router fiber.Router) {
	app := router.Group("inspection_story")
	app.Get("/get-inspection-story", ControllerInspection_Story.GetAllHandler)
	app.Get("/get-by-id/:inspection_story_id", ControllerInspection_Story.GetById)
	app.Post("/create-inspection-story", ControllerInspection_Story.Create)
	app.Patch("/update-inspection-story/:inspection_story_id", ControllerInspection_Story.UpdateInspection_story)
	app.Delete("/delete-inspection-story/:inspection_story_id", ControllerInspection_Story.DeleteInspection_story)
}
