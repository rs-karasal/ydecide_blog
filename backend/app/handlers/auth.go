package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/ydecide_blog/app/dto"
	"github.com/rs-karasal/ydecide_blog/app/models"
	"github.com/rs-karasal/ydecide_blog/app/utils"
	"github.com/rs-karasal/ydecide_blog/database"
)

// func Register(c *fiber.Ctx) error {
// 	var req dto.AuthRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"mmessage": err.Error(),
// 		})
// 	}

// 	var existingUser models.User
// 	if res := database.DB.Where("username = ?", req.Username).First(&existingUser); res.Error == nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "username already taken",
// 		})
// 	}

// 	user := models.User{
// 		Username:     req.Username,
// 		PasswordHash: utils.GeneratePassword(req.Password),
// 	}

// 	res := database.DB.Create(&user)
// 	if res.Error != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": res.Error.Error(),
// 		})
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message": "user created successfully",
// 		"user":    user,
// 	})
// }

func Login(c *fiber.Ctx) error {
	var req dto.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var user models.User
	res := database.DB.Where("username = ?", req.Username).First(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if !utils.ComparePassword(user.PasswordHash, req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}
