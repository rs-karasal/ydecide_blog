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

	post.Use(middleware.JWTProtected())
	post.Post("/", handlers.CreatePost)
	post.Patch("/:id", handlers.UpdatePost)
	post.Delete("/:id", handlers.DeletePost)

	// LifeWheel
	lifeCircle := api.Group("/life-circle")
	lifeCircle.Use(middleware.JWTProtected())
	lifeCircle.Post("/", handlers.CreateLifeCircle)
	lifeCircle.Patch("/", handlers.UpdateLifeCircle)
	lifeCircle.Get("/", handlers.GetLifeCircle)

	// UserProfile
	userProfile := api.Group("/user-profiles")
	userProfile.Get("/:user_id", handlers.GetUserProfile)
	userProfile.Get("/", handlers.GetAllUserProfiles)

	userProfile.Use(middleware.JWTProtected())
	userProfile.Post("/", handlers.CreateUserProfile)
	userProfile.Patch("/", handlers.UpdateUserProfile)
	userProfile.Delete("/", handlers.DeleteUserProfile)
}
