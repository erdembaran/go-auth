package routes

import (
	"github.com/erdembaran/go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/users", controllers.GetUsers)
	app.Get("/api/v1/users/:id", controllers.GetUser)
}