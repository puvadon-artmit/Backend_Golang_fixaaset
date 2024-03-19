package Inspection_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerBranch_Story "github.com/puvadon-artmit/gofiber-template/api/branch_story/controllers"
)

func SetupBranch_StoryRoutes(router fiber.Router) {
	app := router.Group("branch_story")
	app.Get("/get-branch-story", ControllerBranch_Story.GetAllHandler)
	app.Get("/get-by-id/:branch_story_id", ControllerBranch_Story.GetById)
	app.Post("/create-branch-story", ControllerBranch_Story.Create)
	app.Patch("/update-branch-story/:branch_story_id", ControllerBranch_Story.UpdateInspection_story)
	app.Delete("/delete-branch-story/:branch_story_id", ControllerBranch_Story.DeleteInspection_story)
}
