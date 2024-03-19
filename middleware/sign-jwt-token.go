package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/puvadon-artmit/gofiber-template/config"
)

var secretKey = config.GetEnvConfig("SECRET_KEY")

func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: successHandler,
		ErrorHandler:   errorHandler,
		SigningKey:     []byte(secretKey),
		SigningMethod:  "HS256",
	})
}

func successHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	ID := claims["id"].(string)
	c.Locals("id", ID)
	return c.Next()
}

func errorHandler(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}
