package authControllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	AuthDto "github.com/puvadon-artmit/gofiber-template/api/auth/entitys/request"
	AuthServices "github.com/puvadon-artmit/gofiber-template/api/auth/services"
	"github.com/puvadon-artmit/gofiber-template/config"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
)

func GetUserIdByToken(c *fiber.Ctx) error {
	token := config.GetUser(c)
	return c.JSON(fiber.Map{
		"status": "success",
		"result": token,
	})
}

func GetAll(c *fiber.Ctx) error {
	db := database.DB
	var role model.Role
	type User struct {
		UserID string `json:"user_id"`
	}
	type RoleResult struct {
		RoleName        string `json:"role_name"`
		RoleDisplayName string `json:"role_display_name"`
		RoleDescription string `json:"role_description"`
		Users           []User `json:"users"`
	}
	var result RoleResult
	tx := db.Preload("Users").Find(&role).Scan(&result)
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  tx.Error,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetProfile(c *fiber.Ctx) error {
	id := c.Locals("id")
	result, err := AuthServices.GetProfileById(id.(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetProfileById(c *fiber.Ctx) error {
	userID := c.Params("userID")

	profiles, err := AuthServices.GetProfileById(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"profiles": profiles,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	value, err := AuthServices.GetAll()
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

func SignUp(c *fiber.Ctx) error {
	user := new(model.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	_, err = AuthServices.SignUp(*user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"result": "Signup successfully.",
	})
}

func SignIn(c *fiber.Ctx) error {
	var dto AuthDto.LoginDto
	c.BodyParser(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	tokenString, err := AuthServices.SignIn(dto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"token":  tokenString,
	})
}

func CreateNew(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	var roles []model.Role

	createdUser, err := AuthServices.CreateNewUser(user, roles)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   createdUser,
	})
}
