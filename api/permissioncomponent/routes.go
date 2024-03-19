package Permission_componentyRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerPermission_componenty "github.com/puvadon-artmit/gofiber-template/api/permissioncomponent/controllers"
)

func SetupPermission_componentyRoutes(router fiber.Router) {
	app := router.Group("permissioncomponent")
	app.Get("/get-permissioncomponent", ControllerPermission_componenty.GetAllHandler)
	app.Get("/get-by-id/:permission_component_id", ControllerPermission_componenty.GetById)
	app.Post("/create-permissioncomponent", ControllerPermission_componenty.Create)
	app.Patch("/update-permissioncomponent/:permission_component_id", ControllerPermission_componenty.UpdatePermissionComponent)
	app.Delete("/delete-permissioncomponent/:permission_component_id", ControllerPermission_componenty.DeletePermissionComponent)
}
