package roleController

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	RoleServices "github.com/puvadon-artmit/gofiber-template/api/role/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

type Result struct {
	RoleName        string    `json:"role_name"`
	RoleDisplayName string    `json:"role_display_name"`
	RoleDescription string    `json:"role_description"`
	Status          bool      `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

func GetById(c *fiber.Ctx) error {
	result, err := RoleServices.GetById(c.Params("role_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetAll(c *fiber.Ctx) error {
	data, err := RoleServices.GetAll()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": data,
	})
}

func CreateRole(c *fiber.Ctx) error {
	role := new(model.Role)
	c.BodyParser(&role)
	err := validator.New().Struct(role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	result, err := RoleServices.CreateNewRole(*role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}
