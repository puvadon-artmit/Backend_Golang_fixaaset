package Main_CategoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerMain_Category "github.com/puvadon-artmit/gofiber-template/api/main_category/controllers"
)

func SetupMain_CategoryRoutes(router fiber.Router) {
	app := router.Group("main_category")
	app.Get("/get-main_category", ControllerMain_Category.GetAllHandler)
	app.Get("/get-main-id/:id", ControllerMain_Category.GetById)
	app.Post("/create-main_category", ControllerMain_Category.Create)
	app.Patch("/update/:main_category_id", ControllerMain_Category.UpdateMain_Category)
	app.Delete("/delete-main_category/:main_category-id", ControllerMain_Category.DeleteMain_Category)
}
