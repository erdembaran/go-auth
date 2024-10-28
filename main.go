package main

import (
	"fmt"
	"log"

	"github.com/erdembaran/go-auth/config"
	"github.com/erdembaran/go-auth/database"
	"github.com/erdembaran/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("go go go auth!")

	config.LoadEnv()

	database.ConnectDB()

	app := fiber.New()

	routes.Setup(app)

	port := config.GetEnv("PORT", "4000")

	log.Fatal(app.Listen(":" + port ))
}