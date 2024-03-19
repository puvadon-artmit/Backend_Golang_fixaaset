package Count_AutoclikRoutes

import (
	"github.com/gofiber/fiber/v2"
	ControllerCount_Autoclik "github.com/puvadon-artmit/gofiber-template/api/count_autoclik/controllers"
	"github.com/puvadon-artmit/gofiber-template/middleware"
)

func SetupCount_AutoclikRoutes(router fiber.Router) {
	app := router.Group("count_autoclik")
	app.Get("/get-count_autoclik", middleware.AuthorizationRequired(), ControllerCount_Autoclik.GetAllHandler)
	app.Get("/get-by-id/:count_autoclik_id", ControllerCount_Autoclik.GetById)
	app.Post("/create-autoclik-round", ControllerCount_Autoclik.Create)
	app.Patch("/update-count_autoclik/:count_autoclik_id", ControllerCount_Autoclik.UpdateCount_Autoclik)
	app.Delete("/delete-count_autoclik/:count_autoclik_id", ControllerCount_Autoclik.DeleteCount_Autoclik)
}
