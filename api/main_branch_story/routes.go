package TypeplanRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMain_Branch_Story "github.com/puvadon-artmit/gofiber-template/api/main_branch_story/controllers"
)

func SetupMain_Branch_StoryRoutes(router fiber.Router) {
	app := router.Group("main-branch-story")
	app.Get("/get-main-branch-story", ControllerMain_Branch_Story.GetAllHandler)
	app.Get("/get-by-id/:main_branch_story_id", ControllerMain_Branch_Story.GetById)
	app.Post("/create-main-branch-story", ControllerMain_Branch_Story.Create)
	app.Patch("/update-main-branch-story/:main_branch_story_id", ControllerMain_Branch_Story.UpdateMain_Branch_Story)
	app.Delete("/delete-main-branch-story/:main_branch_story_id", ControllerMain_Branch_Story.DeleteMain_Branch_Story)
}
