package LocationRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan "github.com/puvadon-artmit/gofiber-template/api/location/controllers"
)

func SetupLocationRoutes(router fiber.Router) {
	app := router.Group("location")
	app.Get("/get-location", ControllerTypeplan.GetAllHandler)
	app.Get("/get-by-id/:location_id", ControllerTypeplan.GetById)
	app.Post("/create-location", ControllerTypeplan.Create)
	app.Patch("/update-location/:location_id", ControllerTypeplan.UpdateLocation)
	app.Delete("/delete-location/:controller_id", ControllerTypeplan.DeleteLocation)
}
