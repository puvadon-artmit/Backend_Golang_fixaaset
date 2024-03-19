package AssetsRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAssets "github.com/puvadon-artmit/gofiber-template/api/assets/controllers"
)

func SetupAssetsRoutes(router fiber.Router) {
	app := router.Group("assets")
	app.Get("/get-assets", ControllerAssets.GetAllHandler)
	app.Get("/get-property-code/:property_code", ControllerAssets.GetByProperty_codeHandler)
	app.Delete("/get-clear", ControllerAssets.GetDeletePropertyCodes)
	app.Get("/get-category-id/:category_id", ControllerAssets.GetByCategoryHandler)
	app.Get("/get-by-id/:assets_id", ControllerAssets.GetById)
	app.Post("/create-assets", ControllerAssets.Create)
	app.Patch("/update-assets/:assets_id", ControllerAssets.UpdateAssetByID)
	app.Delete("/delete-assets/:assets_id", ControllerAssets.DeleteAsset)
	app.Get("/getcount-assets", ControllerAssets.GetCountAssets)
	app.Get("/getcount-allassets", ControllerAssets.GetCountAllAssets)

}
