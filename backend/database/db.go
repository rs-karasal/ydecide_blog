package database

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/rs-karasal/ydecide_blog/app/models"
	"github.com/rs-karasal/ydecide_blog/app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	fmt.Println("Connection Opened to Database")

	MigrateSchemas()

	CreateSuperUser()
}

func MigrateSchemas() {
	err := DB.AutoMigrate(&models.User{}, &models.Post{}, &models.LifeCircle{})
	if err != nil {
		log.Fatalf("failed to migrate schemas: %v", err)
	}
	fmt.Println("Database schemas migrated")
}

func CreateSuperUser() {
	superDeciderUUID := os.Getenv("SUPERUSER_UUID")
	if superDeciderUUID == "" {
		log.Fatalf("SUPERUSER_UUID is not set in environment variables")
	}

	nullUUID, err := uuid.Parse(superDeciderUUID)
	fmt.Println("nullUUID: ", nullUUID)
	if err != nil {
		log.Fatalf("failed to parse SUPERUSER_UUID: %v", err)
	}

	var user models.User
	err = DB.First(&user, "id = ?", nullUUID).Error
	if err == gorm.ErrRecordNotFound {
		superDeciderName := os.Getenv("SUPERUSER_USERNAME")
		superDeciderPassword := os.Getenv("SUPERUSER_PASSWORD")

		if superDeciderName == "" || superDeciderPassword == "" {
			log.Fatalf("Superuser credentials not found in environment variables")
		}

		superDecider := models.User{
			ID:           nullUUID,
			Username:     superDeciderName,
			PasswordHash: utils.GeneratePassword(superDeciderPassword),
		}

		if err := DB.Create(&superDecider).Error; err != nil {
			log.Fatalf("failed to create superuser: %v", err)
		} else {
			fmt.Println("Superuser created successfully")
		}
	} else if err != nil {
		log.Fatalf("failed to check for superuser: %v", err)
	} else {
		fmt.Println("Superuser already exists")
	}
}
