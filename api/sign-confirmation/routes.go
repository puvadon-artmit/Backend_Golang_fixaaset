package SignatureRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	SignatureController "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation/controllers"
)

func SetupSignatureRoutes(router fiber.Router) {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Static("/uploads", "./uploads")

	asset_checkGroup := router.Group("/signature")

	asset_checkGroup.Get("/get-signature", SignatureController.GetAll)
	asset_checkGroup.Get("/get-by-id/:signature_id", SignatureController.GetByIdHandlers)
	asset_checkGroup.Get("/get-asset-count-id/:asset_count_id", SignatureController.GetByAsset_countHandler)
	// asset_checkGroup.Post("/upload-signature", SignatureController.UploadFile)
	asset_checkGroup.Post("/upload-signature", SignatureController.Create)
	asset_checkGroup.Get("/get-image/:filename", SignatureController.GetImage)
	asset_checkGroup.Patch("/update-signature_id/:signature_id", SignatureController.UpdateSignatureHandler)

}
