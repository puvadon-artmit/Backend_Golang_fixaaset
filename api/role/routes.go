package roleRoutes

import (
	"github.com/gofiber/fiber/v2"
	roleController "github.com/puvadon-artmit/gofiber-template/api/role/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupRoleRoutes(router fiber.Router) {
	app := router.Group("role")
	app.Get("/get-by-id/:role_id", roleController.GetById)
	app.Get("/get-all", middleware.AuthorizationRequired(), roleController.GetAll)
	app.Post("/create", middleware.AuthorizationRequired(), roleController.CreateRole)
}
