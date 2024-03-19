package ControllerMain_branch

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ServicesMain_branch "github.com/puvadon-artmit/gofiber-template/api/main-branch/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := ServicesMain_branch.GetById(c.Params("main_branch_id"))
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
	records, err := ServicesMain_branch.GetAllRecords()
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
	asset_count := new(model.Main_branch)
	err := c.BodyParser(asset_count)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(asset_count)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := ServicesMain_branch.CreateNewMain_branch(*asset_count)
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
// 	var inspections map[string][]model.Main_branch
// 	err := c.BodyParser(&inspections)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	for _, branchList := range inspections {
// 		for _, inspection := range branchList {
// 			err = validator.New().Struct(inspection)
// 			if err != nil {
// 				return c.Status(400).JSON(fiber.Map{
// 					"status": "error",
// 					"error":  err.Error(),
// 				})
// 			}

// 			_, err = ServicesMain_branch.CreateNewMain_branch(inspection)
// 			if err != nil {
// 				return c.Status(500).JSON(fiber.Map{
// 					"status": "error",
// 					"error":  err.Error(),
// 				})
// 			}
// 		}
// 	}

// 	return c.JSON(fiber.Map{
// 		"status": "success",
// 		"result": "Main branches created successfully",
// 	})
// }

func UpdateMain_branch(c *fiber.Ctx) error {
	id := c.Params("main_branch_id")

	updatedType := new(model.Main_branch)
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

	updatedStatus, err := ServicesMain_branch.UpdateMain_branch(id, *updatedType)
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

func DeleteMain_branch(c *fiber.Ctx) error {
	MainBranchID := c.Params("main_branch_id")

	// สร้างตัวแปรเพื่อเก็บข้อมูลที่ต้องการลบ
	asset_count, err := ServicesMain_branch.GetById(MainBranchID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = ServicesMain_branch.DeleteMain_branch(asset_count)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Main_branch deleted successfully",
	})
}
