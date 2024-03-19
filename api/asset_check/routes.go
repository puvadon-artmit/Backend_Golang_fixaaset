package Asset_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	asset_checkController "github.com/puvadon-artmit/gofiber-template/api/asset_check/controllers"
)

func SetupAsset_checkRoutes(router fiber.Router) {
	asset_checkGroup := router.Group("/asset_check")
	asset_checkGroup.Get("/get-asset_check", asset_checkController.GetAll)
	asset_checkGroup.Get("/get-by-id/:asset_check_id", asset_checkController.GetById)
	asset_checkGroup.Get("/get-round-count/:round_count_id", asset_checkController.GetByRound_CountHandler)
	asset_checkGroup.Get("/get-round-latest/:round_count_id", asset_checkController.GetByProperty_code_Id)
	// asset_checkGroup.Get("/get-round-latest", asset_checkController.GetAllHandler)
	asset_checkGroup.Post("/upload-asset", asset_checkController.CreateNew)
	asset_checkGroup.Patch("/update/:asset_check_id", asset_checkController.UpdateAsset_checkHandler)
}
