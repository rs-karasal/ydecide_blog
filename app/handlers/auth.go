package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/ydecide_blog/app/models"
	"github.com/rs-karasal/ydecide_blog/config"
	"github.com/rs-karasal/ydecide_blog/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	query := `SELECT id, password FROM users WHERE username = $1`
	var storedUser models.User
	err := config.DB.QueryRow(context.Background(), query, user.Username).Scan(&storedUser.ID, &storedUser.Password)
	if err != nil {
		log.Printf("Error fetching user from the database: %v\n", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
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

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err = config.DB.Exec(context.Background(), query, user.Username, string(hashedPassword))
	if err != nil {
		log.Printf("Error insrting user into the database: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not register user",
		})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}
