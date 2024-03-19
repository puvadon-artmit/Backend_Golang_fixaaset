package groupRoutes

import (
	"github.com/gofiber/fiber/v2"
	GroupController "github.com/puvadon-artmit/gofiber-template/api/group/controllers"
)

func SetupGroupRoutes(router fiber.Router) {
	app := router.Group("group")
	app.Get("/get-group", GroupController.GetAll)
	app.Get("/get-by-id/:group_id", GroupController.GetById)
	app.Post("/create-group", GroupController.Create)
}
