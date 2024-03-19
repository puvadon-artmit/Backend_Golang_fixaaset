package asset_checkController

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/puvadon-artmit/gofiber-template/api/sign-confirmation/services"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetById(c *fiber.Ctx) error {
	value, err := services.GetById(c.Params("signature_id"))
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

func UpdateSignatureHandler(c *fiber.Ctx) error {
	SignatureID := c.Params("signature_id")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	var updatedItem model.Signature
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	existingSignature, err := services.GetById(SignatureID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// existingSignature.SignatureName = updatedItem.SignatureName

	if len(form.File["signature"]) > 0 {
		file := form.File["signature"][0]

		randomName := uuid.New().String()
		filename := randomName + "_" + file.Filename
		err := c.SaveFile(file, "./uploads/"+filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		existingSignature.Signature = &filename
	}

	if err := services.UpdateSignatureByID(SignatureID, *existingSignature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": existingSignature,
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

func GetByIdHandlers(c *fiber.Ctx) error {
	SignatureID := c.Params("signature_id")

	signature, err := services.GetById(SignatureID)
	if err != nil {
		return c.Status(404).SendString("")
	}

	endpoint := os.Getenv("AWS_REGION")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	expiry := 1 * time.Minute

	filename := *signature.Signature

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), os.Getenv("AWS_BUCKET_NAME"), filename, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	return c.SendString(presignedURL.String())
}

// type CreateSignatureInput struct {
// 	SignatureID   string `json:"signature_id"`
// 	Signature     string `json:"signature"`
// 	Round_countID string `json:"asset_count"`
// 	SignatureName string `form:"signature_name"`
// 	UserID        string `json:"user_id"`
// }

// type UploadForm struct {
// 	Round_countID string `form:"asset_count_id"`
// 	SignatureName string `form:"signature_name"`
// 	Signature     string `form:"signature"`
// 	UserID        string `form:"user_id"`
// }

// func UploadFile(c *fiber.Ctx) error {
// 	form := new(UploadForm)
// 	if err := c.BodyParser(form); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	// ตรวจสอบว่า Round_countID ไม่เป็นค่าว่างเปล่า
// 	if form.Round_countID == "" {
// 		return c.Status(fiber.StatusBadRequest).SendString("Round_countID is required")
// 	}

// 	file, err := c.FormFile("signature")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	randomName := uuid.New().String()
// 	filename := randomName + "_" + file.Filename

// 	err = c.SaveFile(file, "./uploads/"+filename)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	SignatureID := uuid.New().String()
// 	signatureImageName := filename

// 	input := CreateSignatureInput{
// 		SignatureName: form.SignatureName,
// 		Signature:     signatureImageName,
// 		Round_countID: form.Round_countID,
// 		UserID:        form.UserID,
// 	}

// 	item_model := model.Signature{
// 		SignatureID:   SignatureID,
// 		SignatureName: &input.SignatureName,
// 		Signature:     &signatureImageName,
// 		UserID:        input.UserID,
// 	}

// 	if err := database.DB.Create(&item_model).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	return c.SendString(SignatureID)
// }

func Create(c *fiber.Ctx) error {
	signature := new(model.Signature)
	err := c.BodyParser(signature)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(signature)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdSignature, err := services.CreateNewSignature(*signature)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": createdSignature,
	})
}

func GetImage(c *fiber.Ctx) error {
	imagePath := "./uploads/" + c.Params("filename")
	return c.SendFile(imagePath)
}

func GetByAsset_countHandler(c *fiber.Ctx) error {
	value, err := services.GetByAsset_countDB(c.Params("asset_count_id"))
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
