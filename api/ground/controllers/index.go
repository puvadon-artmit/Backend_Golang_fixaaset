package groundController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/api/ground/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("ground_id"))
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

func UpdateItemModelHandler(c *fiber.Ctx) error {
	itemModelID := c.Params("ground_id")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	var updatedItem model.Ground
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	existingItemModel, err := services.GetById(itemModelID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	existingItemModel.GroundName = updatedItem.GroundName
	existingItemModel.Comment = updatedItem.Comment
	existingItemModel.Location_code = updatedItem.Location_code
	existingItemModel.Building = updatedItem.Building
	existingItemModel.Floor = updatedItem.Floor

	if len(form.File["GroundImage"]) > 0 {
		file := form.File["GroundImage"][0]

		// Save the new image
		randomName := uuid.New().String()
		filename := randomName + "_" + file.Filename
		err := c.SaveFile(file, "./uploads/"+filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		existingItemModel.GroundImage = &filename
	}

	if err := services.UpdateItemModelByID(itemModelID, *existingItemModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": existingItemModel,
	})

}

func GetAll(c *fiber.Ctx) error {
	value, err := services.GetAllRecords()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"Message": "Not found!",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": value,
	})
}

type GroundInput struct {
	GroundName    string `json:"ground_name"`
	Comment       string `json:"comment"`
	Location_code string `json:"location_code"`
	Building      string `json:"building"`
	Floor         string `json:"floor"`
	Room          string `json:"room"`
	GroundImage   string `json:"ground_image"`
	UserID        string `json:"user_id"`
}

type UploadForm struct {
	UserID        string `form:"user_id"`
	GroundName    string `form:"ground_name"`
	GroundImage   string `form:"ground_image"`
	Comment       string `form:"comment"`
	Location_code string `form:"product_number"`
	Building      string `form:"building"`
	Floor         string `form:"floor"`
	Room          string `form:"room"`
}

// ---------------------------------------------------------------------------------------------

func UploadFile(c *fiber.Ctx) error {
	form := new(UploadForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	file, err := c.FormFile("ground_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	randomName := uuid.New().String()
	filename := randomName + "_" + file.Filename

	// บันทึกไฟล์ลงใน "./uploads/"
	err = c.SaveFile(file, "./uploads/"+filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	GroundID := uuid.New().String()
	GroundImageImageName := filename

	input := GroundInput{
		GroundName:    form.GroundName,
		GroundImage:   GroundImageImageName,
		UserID:        form.UserID,
		Comment:       form.Comment,
		Location_code: form.Location_code,
		Building:      form.Building,
		Room:          form.Room,
		Floor:         form.Floor,
	}

	ground := model.Ground{
		GroundID:      GroundID,
		GroundName:    &input.GroundName,
		GroundImage:   &GroundImageImageName,
		UserID:        input.UserID,
		Comment:       &input.Comment,
		Location_code: &input.Location_code,
		Room:          &input.Room,
		Building:      &input.Building,
		Floor:         &input.Floor,
	}

	if err := database.DB.Create(&ground).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString(GroundID)
}

func GetImage(c *fiber.Ctx) error {
	imagePath := "./uploads/" + c.Params("filename")
	return c.SendFile(imagePath)
}

func DeleteGroundById(c *fiber.Ctx) error {
	GroundID := c.Params("ground_id")

	ground, err := services.GetById(GroundID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = services.DeleteGroundPhoto(ground)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Ground record deleted successfully",
	})
}
