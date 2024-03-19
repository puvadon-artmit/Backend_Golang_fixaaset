package TypeplanRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMain_branch "github.com/puvadon-artmit/gofiber-template/api/main-branch/controllers"
)

func SetupMain_branchRoutes(router fiber.Router) {
	app := router.Group("main-branch")
	app.Get("/get-main-branch", ControllerMain_branch.GetAllHandler)
	app.Get("/get-by-id/:main-branch-id", ControllerMain_branch.GetById)
	app.Post("/create-main-branch", ControllerMain_branch.Create)
	app.Patch("/update-main-branch/:main-branch-id", ControllerMain_branch.UpdateMain_branch)
	app.Delete("/delete-main-branch/:main-branch-id", ControllerMain_branch.DeleteMain_branch)
}
