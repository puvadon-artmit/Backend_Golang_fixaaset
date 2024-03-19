package manufacturerRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerManufacturer "github.com/puvadon-artmit/gofiber-template/api/manufacturer/controllers"
)

func SetupManufacturerRoutes(router fiber.Router) {
	app := router.Group("manufacturer")
	app.Get("/get-manufacturer", ControllerManufacturer.GetAllHandler)
	app.Get("/get-by-id/:manufacturer_id", ControllerManufacturer.GetById)
	app.Post("/create-manufacturer", ControllerManufacturer.Create)
}
