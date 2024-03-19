package branchRoutes

import (
	"github.com/gofiber/fiber/v2"
	BranchController "github.com/puvadon-artmit/gofiber-template/api/branch/controllers"
)

func SetupBranchRoutes(router fiber.Router) {
	app := router.Group("branch")
	app.Get("/get-all", BranchController.GetAll)
	app.Get("/get-by-id/:permission_id", BranchController.GetById)
	app.Post("/create", BranchController.Create)
}
