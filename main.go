package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rs-karasal/ydecide_blog/app/routes"
	"github.com/rs-karasal/ydecide_blog/config"
)

func main() {
	if os.Getenv("DOCKER_ENV") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Error loading .env file, using environment variables from Docker instead")
		}
	}

	config.ConnectDatabase()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
