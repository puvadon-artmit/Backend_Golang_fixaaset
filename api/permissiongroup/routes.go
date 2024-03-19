package PermissionGrouptyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerPermissionGroupty "github.com/puvadon-artmit/gofiber-template/api/permissiongroup/controllers"
)

func SetupPermissionGroupRoutes(router fiber.Router) {
	app := router.Group("permissiongroup")
	app.Get("/get-permissiongroup", ControllerPermissionGroupty.GetAllHandler)
	app.Get("/get-by-id/:permission_group_id", ControllerPermissionGroupty.GetById)
	app.Post("/create-permissiongroup", ControllerPermissionGroupty.Create)
	app.Patch("/update-permissiongroup/:permission_group_id", ControllerPermissionGroupty.UpdatePermissionGroup)
	app.Delete("/delete-permissiongroup/:permission_group_id", ControllerPermissionGroupty.DeletePermissionGroup)
}
