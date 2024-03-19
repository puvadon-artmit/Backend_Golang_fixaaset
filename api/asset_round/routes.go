package Asset_RoundRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerAsset_Round "github.com/puvadon-artmit/gofiber-template/api/asset_round/controllers"
)

func SetupAsset_RoundRoutes(router fiber.Router) {
	app := router.Group("asset_round")
	app.Get("/get-asset_round", ControllerAsset_Round.GetAllHandler)
	app.Get("/get-by-id/:asset_round_id", ControllerAsset_Round.GetById)
	app.Post("/create-asset-round", ControllerAsset_Round.Create)
	app.Patch("/update-asset_round/:asset_round_id", ControllerAsset_Round.UpdateAsset_Round)
	app.Delete("/delete-asset_round/:asset_round_id", ControllerAsset_Round.DeleteAsset_Round)
}
