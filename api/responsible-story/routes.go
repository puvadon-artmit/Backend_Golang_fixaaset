package Responsible_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerResponsible_Story "github.com/puvadon-artmit/gofiber-template/api/responsible-story/controllers"
)

func SetupResponsible_StoryRoutes(router fiber.Router) {
	app := router.Group("responsible-story")
	app.Get("/get-responsible-story", ControllerResponsible_Story.GetAllHandler)
	app.Get("/get-by-id/:responsible_story_id", ControllerResponsible_Story.GetById)
	app.Post("/create-responsible-story", ControllerResponsible_Story.Create)
	app.Patch("/update-responsible-story/:responsible_story_id", ControllerResponsible_Story.UpdateResponsible_Story)
	app.Delete("/delete-responsible-story/:responsible_story_id", ControllerResponsible_Story.DeleteResponsible_Story)
}
