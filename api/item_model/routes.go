package item_modelRoutes

import (
	//"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	Item_model "github.com/puvadon-artmit/gofiber-template/api/item_model/controllers"
	// "github.com/valyala/fasthttp"
)

func SetupItem_modelRoutes(router fiber.Router) {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Static("/uploads", "./uploads")

	Item_modelGroup := router.Group("/item-model")

	Item_modelGroup.Get("/get-item_model", Item_model.GetAll)
	Item_modelGroup.Get("/get-by-id/:item_model_id", Item_model.GetById)
	Item_modelGroup.Post("/upload", Item_model.UploadFile)
	Item_modelGroup.Post("/create", Item_model.Create)
	Item_modelGroup.Post("/get-frontpicture-url", Item_model.GetByIdHandler)
	Item_modelGroup.Post("/delete-frontpicture/:item_model_id", Item_model.DeletePhotos)
	Item_modelGroup.Post("/delete/:item_model_id", Item_model.DeletePhotos)
	Item_modelGroup.Get("/get-image/:filename", Item_model.GetImage)
	Item_modelGroup.Get("/delete-unused-images", Item_model.DeleteUnusedImagesHandler)
	Item_modelGroup.Patch("/update/:item_model_id", Item_model.UpdateItemModelHandler)

	// go func() {
	// 	ticker := time.NewTicker(3000 * time.Second)
	// 	defer ticker.Stop()
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			ctx := app.AcquireCtx(new(fasthttp.RequestCtx))
	// 			defer app.ReleaseCtx(ctx)
	// 			Item_model.DeleteUnusedImagesHandler(ctx)
	// 		}
	// 	}
	// }()
}
