package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/ydecide_blog/app/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)
}
