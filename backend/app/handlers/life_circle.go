package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/ydecide_blog/app/dto"
	"github.com/rs-karasal/ydecide_blog/app/models"
	"github.com/rs-karasal/ydecide_blog/config"
	"github.com/rs-karasal/ydecide_blog/database"
	"gorm.io/gorm"
)

func GetLifeCircle(c *fiber.Ctx) error {
	var lifeCircle models.LifeCircle

	if err := database.DB.Where("user_id = ?", config.SuperDeciderUUID).First(&lifeCircle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "LifeCircle not found for this user",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve life circle",
		})
	}

	response := dto.LifeCircleResponse{
		HealthAndBody:           lifeCircle.HealthAndBody,
		LoveAndRelationships:    lifeCircle.LoveAndRelationships,
		FamilyAndFriends:        lifeCircle.FamilyAndFriends,
		PersonalGrowth:          lifeCircle.PersonalGrowth,
		CareerAndFinance:        lifeCircle.CareerAndFinance,
		JoyAndRelax:             lifeCircle.JoyAndRelax,
		PhysicalEnvironment:     lifeCircle.PhysicalEnvironment,
		EmotionsAndFullfillment: lifeCircle.EmotionsAndFullfillment,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"life_circle": response,
	})
}

func CreateLifeCircle(c *fiber.Ctx) error {
	var lifeCircle models.LifeCircle
	if err := database.DB.Where("user_id = ?", config.SuperDeciderUUID).First(&lifeCircle).Error; err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "LifeCircle already exists for this user",
		})
	}

	newLifeCircle := models.LifeCircle{
		UserID:                  config.SuperDeciderUUID,
		HealthAndBody:           0,
		LoveAndRelationships:    0,
		FamilyAndFriends:        0,
		PersonalGrowth:          0,
		CareerAndFinance:        0,
		JoyAndRelax:             0,
		PhysicalEnvironment:     0,
		EmotionsAndFullfillment: 0,
	}

	if err := database.DB.Create(&newLifeCircle).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create life circle",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "LifeCircle created successfully",
		"life_circle": newLifeCircle,
	})
}

func UpdateLifeCircle(c *fiber.Ctx) error {
	log.Printf("Request body: %s", string(c.Body()))

	var updateRequest dto.UpdateLifeCircleRequest
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	log.Printf("Parsed request: %+v", updateRequest)

	var lifeCircle models.LifeCircle
	if err := database.DB.Where("user_id = ?", config.SuperDeciderUUID).First(&lifeCircle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "LifeCircle not found for this user",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve life circle",
		})
	}

	lifeCircle.HealthAndBody = updateRequest.HealthAndBody
	fmt.Println("HEALTH SUKA:", updateRequest.HealthAndBody)
	lifeCircle.LoveAndRelationships = updateRequest.LoveAndRelationships
	lifeCircle.FamilyAndFriends = updateRequest.FamilyAndFriends
	lifeCircle.PersonalGrowth = updateRequest.PersonalGrowth
	lifeCircle.CareerAndFinance = updateRequest.CareerAndFinance
	lifeCircle.JoyAndRelax = updateRequest.JoyAndRelax
	lifeCircle.PhysicalEnvironment = updateRequest.PhysicalEnvironment
	lifeCircle.EmotionsAndFullfillment = updateRequest.EmotionsAndFullfillment

	if err := database.DB.Save(&lifeCircle).Error; err != nil {
		log.Println("Error updating LifeCircle:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update life circle",
		})
	}
	log.Println("LifeCircle updated successfully:", lifeCircle)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "LifeCircle updated successfully",
		"life_circle": lifeCircle,
	})
}
