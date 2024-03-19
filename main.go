package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/puvadon-artmit/gofiber-template/config"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/router"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false, //Deploy ปรับเป็น true
	})
	app.Use(cors.New(config.CorsConfigDefault))
	database.ConnectDB()
	router.SetupRoutes(app)
	app.Listen(":7000")
	Aws()

}
