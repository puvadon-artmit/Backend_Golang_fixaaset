package Controllercount_main_category

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	Servicescount_main_category "github.com/puvadon-artmit/gofiber-template/api/count_main_category/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := Servicescount_main_category.GetById(c.Params("Asset_count_Main_Category_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

func GetAllHandler(c *fiber.Ctx) error {
	records, err := Servicescount_main_category.GetAllRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": records,
	})
}

func Create(c *fiber.Ctx) error {
	count_main_category := new(model.Asset_count_Main_Category)
	err := c.BodyParser(count_main_category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(count_main_category)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := Servicescount_main_category.CreateNewcount_main_category(*count_main_category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdStatus,
	})
}

func Updatecount_main_category(c *fiber.Ctx) error {
	id := c.Params("Asset_count_Main_Category_id")

	updatedType := new(model.Asset_count_Main_Category)
	err := c.BodyParser(updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedType)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	updatedStatus, err := Servicescount_main_category.Updatecount_main_category(id, *updatedType)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedStatus,
	})
}

func Deletecount_main_category(c *fiber.Ctx) error {
	count_main_categoryID := c.Params("asset_Asset_count_Main_Category_id")

	count_main_category, err := Servicescount_main_category.GetById(count_main_categoryID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = Servicescount_main_category.Deletecount_main_category(count_main_category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "count_main_category deleted successfully",
	})
}

func GetAsset_countHandler(c *fiber.Ctx) error {
	value, err := Servicescount_main_category.GetByAsset_countIDDB(c.Params("asset_count_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}
