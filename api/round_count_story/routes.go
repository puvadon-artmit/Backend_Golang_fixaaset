package Round_Count_StoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerRound_Count_Story "github.com/puvadon-artmit/gofiber-template/api/round_count_story/controllers"
)

func SetupRound_Count_StoryRoutes(router fiber.Router) {
	app := router.Group("round-count-story")
	app.Get("/get-count-story", ControllerRound_Count_Story.GetAllHandler)
	app.Get("/get-by-id/:round_count_id", ControllerRound_Count_Story.GetById)
	app.Post("/create-count-story", ControllerRound_Count_Story.Create)
	app.Patch("/update-count-story/:round_count_id", ControllerRound_Count_Story.UpdateRound_Count_Story)
	app.Delete("/delete-count-story/:round_count_id", ControllerRound_Count_Story.DeleteRound_Count_Story)
}
