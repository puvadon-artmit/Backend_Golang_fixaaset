package ControllerAssets

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesAssets "github.com/puvadon-artmit/gofiber-template/api/assets/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesAssets.GetById(c.Params("assets_id"))
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

func GetByCategoryHandler(c *fiber.Ctx) error {
	value, err := ServicesAssets.GetByCategoryIDDB(c.Params("category_id"))
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

func GetCountAssets(c *fiber.Ctx) error {
	count, err := ServicesAssets.CountLatestRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(count)

}

func GetCountAllAssets(c *fiber.Ctx) error {
	allcount, err := ServicesAssets.CountAssetsByCategory()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(allcount)

}

func GetByProperty_codeHandler(c *fiber.Ctx) error {
	value, err := ServicesAssets.GetByProperty_codeDB(c.Params("property_code"))
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

func GetAllHandler(c *fiber.Ctx) error {
	records, err := ServicesAssets.GetLatestRecords()
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

func GetDeletePropertyCodes(c *fiber.Ctx) error {
	err := ServicesAssets.DeleteUnusedPropertyCodes()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Unused property codes deleted successfully",
	})
}

func Create(c *fiber.Ctx) error {
	assets := new(model.Assets)
	err := c.BodyParser(assets)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(assets)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesAssets.CreateNewAssets(*assets)
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

// func Create(c *fiber.Ctx) error {
// 	assets := make([]model.Assets, 0)
// 	err := c.BodyParser(&assets)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	for _, asset := range assets {
// 		err = validator.New().Struct(asset)
// 		if err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"status": "error",
// 				"error":  err.Error(),
// 			})
// 		}
// 	}

// 	createdAssets, err := ServicesAssets.CreateNewAssets(assets)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": createdAssets,
// 	})
// }

func UpdateAssetByID(c *fiber.Ctx) error {
	id := c.Params("assets_id")

	updatedLocation := new(model.Assets)
	err := c.BodyParser(updatedLocation)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(updatedLocation)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = ServicesAssets.UpdateAssetByID(id, *updatedLocation)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": updatedLocation,
	})
}

func DeleteAsset(c *fiber.Ctx) error {
	AssetsID := c.Params("assets_id")

	assets, err := ServicesAssets.GetById(AssetsID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAssets.DeleteAsset(assets)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "DeleteAsset deleted successfully",
	})
}

func DeletePropertyCodes(c *fiber.Ctx) error {
	AssetsID := c.Params("assets_id")

	assets, err := ServicesAssets.GetById(AssetsID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesAssets.DeleteAsset(assets)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "DeleteAsset deleted successfully",
	})
}
