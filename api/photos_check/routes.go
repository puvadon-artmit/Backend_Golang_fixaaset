package photos_checkRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	photos_checkController "github.com/puvadon-artmit/gofiber-template/api/photos_check/controllers"
)

func SetupPhotos_checkRoutes(router fiber.Router) {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Static("/uploads", "./uploads")

	photos_checkGroup := router.Group("/photos_check")
	photos_checkGroup.Get("/get-all", photos_checkController.GetAll)
	photos_checkGroup.Get("/get-by-id/:photos_check_id", photos_checkController.GetByIdHandler)
	// photos_checkGroup.Get("/get-photo/:photo", photos_checkController.GetPhotoById)
	// photos_checkGroup.Post("/upload", photos_checkController.UploadFileToS3AndSaveToDB)
	photos_checkGroup.Post("/create", photos_checkController.Create)
	photos_checkGroup.Get("/get-signed-url/:photos_check_id", photos_checkController.GetByIdHandler)
	// photos_checkGroup.Get("/get-image/:filename", photos_checkController.GetImageCheck)
	photos_checkGroup.Delete("/delete-image/:photos_check_id", photos_checkController.DeletePhotos)

}
