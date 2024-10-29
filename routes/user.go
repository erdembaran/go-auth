package routes

import (
	"github.com/erdembaran/go-auth/controllers"
	"github.com/erdembaran/go-auth/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/api/v1/users", middleware.JWTMiddleware)

	userGroup.Get("/", controllers.GetUsers)
	userGroup.Get("/:id", controllers.GetUser)
}