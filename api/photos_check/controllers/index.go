package photos_checkController

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
	"github.com/puvadon-artmit/gofiber-template/api/photos_check/services"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
)

// func GetPhotoById(c *fiber.Ctx) error {
// 	value, err := services.GetById(c.Params("photos_check_id"))
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err,
// 		})
// 	}

// 	responseData := fiber.Map{
// 		"status": "success",
// 		"result": value,
// 	}
// 	if value.Photos != "" {
// 		imagePath := "./uploads/" + value.Photos
// 		return c.SendFile(imagePath)
// 	}

// 	return c.JSON(responseData)
// }

func GetAll(c *fiber.Ctx) error {
	value, err := services.GetAll()
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

// func UploadFile(c *fiber.Ctx) error {
// 	file, err := c.FormFile("photos")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	randomName := uuid.New().String()
// 	filename := randomName + filepath.Ext(file.Filename)

// 	err = c.SaveFile(file, "./uploads/"+filename)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	photos_checkID := uuid.New().String()
// 	photos_check := model.Photos_check{
// 		Photos_checkID: photos_checkID,
// 		Photos:         filename,
// 	}

// 	if err := database.DB.Create(&photos_check).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	responseData := struct {
// 		Photos_checkID string `json:"photos_check_id"`
// 		Photos         string `json:"photos"`
// 	}{
// 		Photos_checkID: photos_checkID,
// 		Photos:         filename,
// 	}

// 	return c.JSON(responseData)
// }

// func UploadFile(c *fiber.Ctx) error {
// 	file, err := c.FormFile("photos")
// 	photos_checkID := uuid.New().String()
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	randomName := uuid.New().String()
// 	filename := randomName + filepath.Ext(file.Filename)

// 	err = c.SaveFile(file, "./uploads/"+filename)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	photos_check := model.Photos_check{
// 		Photos_checkID: photos_checkID,
// 		Photos:         filename,
// 	}

// 	if err := database.DB.Create(&photos_check).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	if err := UploadToS3(filename, "./uploads/"+filename); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	if err := os.Remove("./uploads/" + filename); err != nil {
// 		log.Println("Failed to delete temporary file:", err)
// 	}

// 	responseData := struct {
// 		Photos_checkID string `json:"photos_check_id"`
// 		Photos         string `json:"photos"`
// 	}{
// 		Photos_checkID: photos_checkID,
// 		Photos:         filename,
// 	}

// 	return c.JSON(responseData)
// }

// func UploadToS3(filename, filePath string) error {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion("ap-southeast-1"),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	client := s3.NewFromConfig(cfg)

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
// 		Bucket: aws.String("acgbucket-all"),
// 		Key:    aws.String(filename),
// 		Body:   file,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("File uploaded to S3: %s\n", filename)
// 	return nil
// }

// func UploadFileToS3AndSaveToDB(c *fiber.Ctx) error {
// 	// รับไฟล์จาก multipart/form-data request
// 	file, err := c.FormFile("photos")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	// สร้างชื่อสุ่มสำหรับไฟล์
// 	randomName := uuid.New().String()
// 	filename := randomName + filepath.Ext(file.Filename)

// 	// เปิดไฟล์เพื่ออ่านข้อมูล
// 	fileContent, err := file.Open()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}
// 	defer fileContent.Close()

// 	fileSize := file.Size
// 	fileBuffer := make([]byte, fileSize)
// 	_, err = fileContent.Read(fileBuffer)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	err = UploadBufferToS3(filename, fileBuffer)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	// บันทึกข้อมูลเกี่ยวกับไฟล์ลงในฐานข้อมูล
// 	err = SaveFileDataToDB(filename)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	responseData := struct {
// 		Filename string `json:"filename"`
// 	}{
// 		Filename: filename,
// 	}

// 	return c.JSON(responseData)
// }

// func UploadBufferToS3(filename string, fileBuffer []byte) error {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion("ap-southeast-1"),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	client := s3.NewFromConfig(cfg)

// 	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
// 		Bucket: aws.String("acgbucket-all"),
// 		Key:    aws.String(filename),
// 		Body:   bytes.NewReader(fileBuffer),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("File uploaded to S3: %s\n", filename)
// 	return nil
// }

func SaveFileDataToDB(filename string) error {

	photos_checkID := uuid.New().String()
	photos_check := model.Photos_check{
		Photos_checkID: photos_checkID,
		Photos:         filename,
	}

	if err := database.DB.Create(&photos_check).Error; err != nil {
		return err
	}

	return nil
}

// func GetImage(c *fiber.Ctx) error {
// 	imagePath := "./uploads/" + c.Params("filename")
// 	return c.SendFile(imagePath)
// }

// func GetImageCheck(c *fiber.Ctx) error {
// 	filename := c.Params("filename") + ".jpg"

// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion("ap-southeast-1"),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	client := s3.NewFromConfig(cfg)

// 	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
// 		Bucket: aws.String("acgbucket-all"),
// 		Key:    aws.String(filename),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	imageData, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	imageBase64 := base64.StdEncoding.EncodeToString(imageData)

// 	tmpl := `<img src="data:image/png;base64,{{ .Image }}" alt="Uploaded Image">`
// 	t := template.Must(template.New("image").Parse(tmpl))
// 	data := struct {
// 		Image string
// 	}{
// 		Image: imageBase64,
// 	}

// 	var tpl bytes.Buffer
// 	if err := t.Execute(&tpl, data); err != nil {
// 		return err
// 	}

// 	return c.SendString(tpl.String())
// }

// func DeletePhotos(c *fiber.Ctx) error {
// 	photosCheckID := c.Params("photos_check_id")

// 	photosCheck, err := services.GetById(photosCheckID)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{
// 			"status":  "error",
// 			"message": "Not found!",
// 			"error":   err,
// 		})
// 	}

// 	err = services.DeletePhotos_check(photosCheck)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"status": "error",
// 			"error":  err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"status":  "success",
// 		"message": "Photos deleted successfully",
// 	})
// }

// func deleteFileIfExists(filename string) error {
// 	filePath := filepath.Join("./uploads/", filename)
// 	if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 		return nil
// 	}
// 	if err := os.Remove(filePath); err != nil {
// 		return err
// 	}
// 	return nil
// }

func Create(c *fiber.Ctx) error {
	photos_check := new(model.Photos_check)
	err := c.BodyParser(photos_check)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	err = validator.New().Struct(photos_check)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	createdStatus, err := services.CreateNewPhotos_check(*photos_check)
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

func GetAllHandler(c *fiber.Ctx) error {
	records, err := services.GetAll()
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

// func GetByIdHandler(c *fiber.Ctx) error {
// 	photosURL, err := services.GetById(c.Params("photos_check_id"))
// 	if err != nil {
// 		return c.Status(404).SendString("")
// 	}
// 	return c.SendString(photosURL)
// }

// func GetByIdHandler(c *fiber.Ctx) error {

// 	photosCheckID := c.Params("photos_check_id")

// 	photosURL, err := services.GetById(photosCheckID)
// 	if err != nil {
// 		return c.Status(404).SendString("")
// 	}

// 	presignedURL := fmt.Sprintf("https://acgbucket-all.s3.ap-southeast-1.amazonaws.com/%s?AWSAccessKeyId=AKIAIYXHK62WHKGF7W4Q&Expires=%d", photosURL, time.Now().Add(5*time.Minute).Unix())

// 	return c.SendString(presignedURL)
// }

func GetByIdHandler(c *fiber.Ctx) error {
	photosCheckID := c.Params("photos_check_id")

	photosURL, err := services.GetById(photosCheckID)
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

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), os.Getenv("AWS_BUCKET_NAME"), photosURL, expiry, nil)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate pre-signed URL")
	}

	return c.SendString(presignedURL.String())
}

func DeletePhotos(c *fiber.Ctx) error {
	photosCheckID := c.Params("photos_check_id")

	photosCheck, err := services.GetByPhotoId(photosCheckID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Not found!",
			"error":   err,
		})
	}

	err = services.DeletePhotos_check(photosCheck)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Photos record deleted successfully",
	})
}
