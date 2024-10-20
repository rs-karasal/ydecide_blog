package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/ydecide_blog/app/handlers"
	"github.com/rs-karasal/ydecide_blog/pkg/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	// auth.Post("/register", handlers.Register)

	// Post
	post := api.Group("/posts")
	post.Get("/", handlers.GetPosts)
	post.Get("/:id", handlers.GetPost)

	post.Use(middleware.JWTProtected)
	post.Post("/", handlers.CreatePost)
	post.Patch("/:id", handlers.UpdatePost)
	post.Delete("/:id", handlers.DeletePost)

	lifeCircle := api.Group("/life-circle")
	lifeCircle.Use(middleware.JWTProtected)
	lifeCircle.Post("/create", handlers.CreateLifeCircle)
	lifeCircle.Patch("/update", handlers.UpdateLifeCircle)
	lifeCircle.Get("/", handlers.GetLifeCircle)
}
