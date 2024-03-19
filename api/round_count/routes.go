package Round_CountRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerRound_Count "github.com/puvadon-artmit/gofiber-template/api/round_count/controllers"
)

func SetupRound_CountRoutes(router fiber.Router) {
	app := router.Group("round-count")
	app.Get("/get-round-count", ControllerRound_Count.GetAllHandler)
	app.Get("/get-delete", ControllerRound_Count.GetDeleteHandler)
	app.Get("/get-by-id/:round_count_id", ControllerRound_Count.GetById)
	app.Get("/get-count-id/:Asset_countID", ControllerRound_Count.GetByNoHandler)
	app.Post("/create-round-count", ControllerRound_Count.Create)
	app.Patch("/update-round-count/:round_count_id", ControllerRound_Count.UpdateRound_Count)
	app.Delete("/delete-round-count/:round_count_id", ControllerRound_Count.DeleteRound_Count)
}
