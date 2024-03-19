package categoryRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	categoryController "github.com/puvadon-artmit/gofiber-template/api/category/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupCategoryRoutes(router fiber.Router) {
	// Create a new Fiber app instance
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Serve static files from the "uploads" folder
	app.Static("/uploads", "./uploads")

	// Group the routes under "/category"
	categoryGroup := router.Group("/category")

	// Define your category routes
	categoryGroup.Get("/get-all", middleware.RoleAuthorizationRequired("4nG8pLx9Wc7q3Sd5Fb2Vh6Kt1Ry0mJzOuPw"), middleware.AuthorizationRequired(), categoryController.GetAll)
	categoryGroup.Get("/get-by-id/:category_id", middleware.RoleAuthorizationRequired("4nG8pLx9Wc7q3Sd5Fb2Vh6Kt1Ry0mJzOuPw", "2dFp4sHc6qRt8xYv9MwZ5nTbGjX1lV7yU3K"), middleware.AuthorizationRequired(), categoryController.GetById)
	categoryGroup.Post("/upload", categoryController.UploadFile)
	categoryGroup.Post("/create", categoryController.CreateCategory)
	categoryGroup.Get("/get-image/:filename", categoryController.GetImage)
	categoryGroup.Get("/get-main-id/:main_category_id", categoryController.GetByMain_CategoryHandler)
	categoryGroup.Get("/getcount-category", categoryController.GetCountCategory)
	categoryGroup.Get("/getcount-categorys", middleware.RoleAuthorizationRequired("4nG8pLx9Wc7q3Sd5Fb2Vh6Kt1Ry0mJzOuPw", "2dFp4sHc6qRt8xYv9MwZ5nTbGjX1lV7yU3K"), middleware.AuthorizationRequired(), categoryController.GetCountCategory)
}
