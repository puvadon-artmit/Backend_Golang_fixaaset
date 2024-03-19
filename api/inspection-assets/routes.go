package TypeplanRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerInspection "github.com/puvadon-artmit/gofiber-template/api/inspection-assets/controllers"
)

func SetupAsset_countRoutes(router fiber.Router) {
	app := router.Group("inspection")
	app.Get("/get-inspection", ControllerInspection.GetAllHandler)
	app.Get("/get-by-id/:asset_count_id", ControllerInspection.GetById)
	app.Post("/create-inspection", ControllerInspection.Create)
	app.Patch("/update-inspection/:asset_count_id", ControllerInspection.UpdateAsset_count)
	app.Delete("/delete-inspection/:controller_id", ControllerInspection.DeleteAsset_count)
	app.Get("/getcount-asset-count", ControllerInspection.GetCountAsset_count)
}
