package count_categoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerTypeplan "github.com/puvadon-artmit/gofiber-template/api/count_category/controllers"
)

func Setupcount_categoryRoutes(router fiber.Router) {
	app := router.Group("count_category")
	app.Get("/get-count-category", ControllerTypeplan.GetAllHandler)
	app.Get("/get-by-id/:asset_count_category_id", ControllerTypeplan.GetById)
	app.Get("/get-asset-count/:asset_count_id", ControllerTypeplan.GetAsset_countHandler)
	app.Post("/create-count-category", ControllerTypeplan.Create)
	app.Patch("/update-count-category/:asset_count_category_id", ControllerTypeplan.Updatecount_category)
	app.Delete("/delete-count-category/:asset_count_category_id", ControllerTypeplan.Deletecount_category)
}
