package statusRoutes

import (
	"github.com/gofiber/fiber/v2"
	StatusController "github.com/puvadon-artmit/gofiber-template/api/status/controllers"
)

func SetupStatusRoutes(router fiber.Router) {
	app := router.Group("status")
	app.Get("/get-status", StatusController.GetAll)
	app.Get("/get-by-id/:status_id", StatusController.GetById)
	app.Post("/create-status", StatusController.Create)
}
