package middleware

import (
	"github.com/erdembaran/go-auth/utils"
	"github.com/gofiber/fiber/v2"
)


func JWTMiddleware(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}

	isValid, err := utils.VerifyToken(tokenString)
	if err != nil || !isValid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid JWT",
		})
	}

	return c.Next()
}
