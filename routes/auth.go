package routes

import (
	"github.com/erdembaran/go-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	authGroup := app.Group("/api/v1/auth")

	authGroup.Post("/login", controllers.Login)
	authGroup.Post("/logout", controllers.Logout)
	authGroup.Post("/register", controllers.Register)
	authGroup.Post("/forgot-password", controllers.ForgotPassword)
}
