package typeRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerType "github.com/puvadon-artmit/gofiber-template/api/type_things/controllers"
)

func SetupTypeRoutes(router fiber.Router) {
	app := router.Group("type")
	app.Get("/get-type", ControllerType.GetAllHandler)
	app.Get("/get-by-id/:type_id", ControllerType.GetById)
	app.Post("/create-type", ControllerType.Create)
}
