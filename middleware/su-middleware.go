package middleware

import "github.com/gofiber/fiber/v2"

func GetUserRoles(c *fiber.Ctx) []string {
	roles := c.Locals("roles").([]string)
	return roles
}

func RoleAuthorizationRequired(role_code ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestedRole := c.Query("role")

		authorized := false
		for _, role_code := range role_code {
			if role_code == requestedRole {
				authorized = true
				break
			}
		}

		if !authorized {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
				"msg":   "You don't have permission to access this resource.",
			})
		}
		return c.Next()
	}
}
