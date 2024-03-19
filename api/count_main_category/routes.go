package count_main_categoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan "github.com/puvadon-artmit/gofiber-template/api/count_main_category/controllers"
)

func Setupcount_main_categoryRoutes(router fiber.Router) {
	app := router.Group("count-main-category")
	app.Get("/get-count-main-category", ControllerTypeplan.GetAllHandler)
	app.Get("/get-by-id/:asset_count_main_category_id", ControllerTypeplan.GetById)
	app.Get("/get-main-asset-count/:asset_count_id", ControllerTypeplan.GetAsset_countHandler)
	app.Post("/create-count-main-category", ControllerTypeplan.Create)
	app.Patch("/update-count-main-category/:asset_count_main_category_id", ControllerTypeplan.Updatecount_main_category)
	app.Delete("/delete-count-main-category/:asset_count_main_category_id", ControllerTypeplan.Deletecount_main_category)
}
