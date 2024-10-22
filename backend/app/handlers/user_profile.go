package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs-karasal/ydecide_blog/app/dto"
	"github.com/rs-karasal/ydecide_blog/app/models"
	"github.com/rs-karasal/ydecide_blog/database"
	"gorm.io/gorm"
)

func GetUserProfile(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	var profile models.UserProfile
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User profile not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"userProfile": profile,
	})
}

func GetAllUserProfiles(c *fiber.Ctx) error {
	var profiles []models.UserProfile

	if err := database.DB.Find(&profiles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user profiles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"userProfiles": profiles,
	})
}

// func GetOwnUserProfile(c *fiber.Ctx) error {
// 	userID := c.Locals("user_id").(string)

// 	var profile models.UserProfile
// 	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 				"error": "User profile not found",
// 			})
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to retrieve user profile",
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"userProfile": profile,
// 	})
// }

func CreateUserProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized (no user_id in locals)",
		})
	}

	// Check if the profile already exists for this user
	var existingProfile models.UserProfile
	if err := database.DB.Where("user_id = ?", userID).First(&existingProfile).Error; err != gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserProfile already exists for this user",
		})
	}

	// Parse the request body
	var req dto.UserProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Create the UserProfile
	userProfile := models.UserProfile{
		UUIDModel:   models.UUIDModel{ID: userID},
		Name:        sql.NullString{String: req.Name, Valid: req.Name != ""},
		BirthDate:   sql.NullString{String: req.BirthDate, Valid: req.BirthDate != ""},
		Photo:       sql.NullString{String: req.Photo, Valid: req.Photo != ""},
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Profession:  sql.NullString{String: req.Profession, Valid: req.Profession != ""},
		Languages:   sql.NullString{String: req.Languages, Valid: req.Languages != ""},
		Location:    sql.NullString{String: req.Location, Valid: req.Location != ""},
	}

	// Save the new profile to the database
	if err := database.DB.Create(&userProfile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user profile",
		})
	}

	// Return success
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":      "UserProfile created successfully",
		"user_profile": userProfile,
	})
}

func UpdateUserProfile(c *fiber.Ctx) error {
	var req dto.UserProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized (invalid or missing user_id)",
		})
	}

	var profile models.UserProfile
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User profile not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user profile",
		})
	}

	profile.Name = sql.NullString{String: req.Name, Valid: req.Name != ""}
	profile.BirthDate = sql.NullString{String: req.BirthDate, Valid: req.BirthDate != ""}
	profile.Photo = sql.NullString{String: req.Photo, Valid: req.Photo != ""}
	profile.Description = sql.NullString{String: req.Description, Valid: req.Description != ""}
	profile.Profession = sql.NullString{String: req.Profession, Valid: req.Profession != ""}
	profile.Languages = sql.NullString{String: req.Languages, Valid: req.Languages != ""}
	profile.Location = sql.NullString{String: req.Location, Valid: req.Location != ""}

	if err := database.DB.Model(&profile).Where("id = ?", profile.ID).Updates(profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Profile updated successfully",
		"userProfile": profile,
	})
}

func DeleteUserProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized (invalid or missing user_id)",
		})
	}

	fmt.Println("USER ID:", userID)

	var profile models.UserProfile
	if err := database.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User profile not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user profile",
		})
	}

	fmt.Println("PROFILE:", profile)

	if err := database.DB.Unscoped().Where("user_id = ?", userID).Delete(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profile deleted successfully",
	})
}
