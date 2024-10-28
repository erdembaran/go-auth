package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("go go go auth!")

	app := fiber.New()

	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "go go go auth!",
		})
	})

	log.Fatal(app.Listen(":4000"))
}