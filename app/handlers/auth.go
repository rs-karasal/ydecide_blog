package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/ydecide_blog/app/models"
	"github.com/rs-karasal/ydecide_blog/pkg/jwt"
)

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// TODO: mocked user verification, need to replace with actual DB check
	if user.Username != "admin" || user.Password != "password" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(user.Username, time.Hour*24)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
	}

	return c.JSON(fiber.Map{"token": token})
}
