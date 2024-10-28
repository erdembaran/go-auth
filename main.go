package main

import (
	"fmt"
	"log"

	"github.com/erdembaran/go-auth/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("go go go auth!")

	config.LoadEnv()

	app := fiber.New()

	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "go go go auth!",
		})
	})

	port := config.GetEnv("PORT", "4000")

	log.Fatal(app.Listen(":" + port ))
}