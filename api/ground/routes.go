package item_modelRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	GroundControllers "github.com/puvadon-artmit/gofiber-template/api/ground/controllers"
)

func SetupGroundRoutes(router fiber.Router) {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Static("/uploads", "./uploads")

	categoryGroup := router.Group("/ground")

	categoryGroup.Get("/get-ground", GroundControllers.GetAll)
	categoryGroup.Get("/get-by-id/:ground_id", GroundControllers.GetById)
	categoryGroup.Post("/upload", GroundControllers.UploadFile)
	categoryGroup.Get("/get-image/:filename", GroundControllers.GetImage)
	categoryGroup.Patch("/update/:ground_id", GroundControllers.UpdateItemModelHandler)
}
