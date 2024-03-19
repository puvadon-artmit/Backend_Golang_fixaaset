package permissionRoutes

import (
	"github.com/gofiber/fiber/v2"
	permissionController "github.com/puvadon-artmit/gofiber-template/api/permission/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupPermissionRoutes(router fiber.Router) {
	app := router.Group("permission")
	app.Get("/get-all", middleware.AuthorizationRequired(), permissionController.GetAll)
	app.Get("/get-by-id/:permission_id", middleware.AuthorizationRequired(), permissionController.GetById)
	app.Post("/create", middleware.AuthorizationRequired(), permissionController.Create)
}
