package typeRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerResponsible "github.com/puvadon-artmit/gofiber-template/api/responsible/controllers"
)

func SetupResponsibleRoutes(router fiber.Router) {
	app := router.Group("responsible")
	app.Get("/get-responsible", ControllerResponsible.GetAllHandler)
	app.Get("/get-by-id/:responsible_id", ControllerResponsible.GetById)
	app.Post("/create-responsible", ControllerResponsible.Create)
}
