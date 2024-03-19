package TypeplanRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan "github.com/puvadon-artmit/gofiber-template/api/typeplan/controllers"
)

func SetupTypeplanRoutes(router fiber.Router) {
	app := router.Group("typeplan")
	app.Get("/get-typeplan", ControllerTypeplan.GetAllHandler)
	app.Get("/get-by-id/:typeplan_id", ControllerTypeplan.GetById)
	app.Post("/create-typeplan", ControllerTypeplan.Create)
	app.Patch("/update-typeplan/:typeplan_id", ControllerTypeplan.UpdateTypeplan)
	app.Delete("/delete-typeplan/:controller_id", ControllerTypeplan.DeleteTypeplan)
}
