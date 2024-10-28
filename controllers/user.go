package controllers

import (
	"context"

	"github.com/erdembaran/go-auth/database"
	"github.com/erdembaran/go-auth/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	cursor, err := database.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return err
		}

		users = append(users, user)
	}
		return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "Invalid user ID format"})
	}

	filter := bson.M{"_id": userId}

	user := database.Collection.FindOne(c.Context(), filter)

	if user.Err() != nil {
		return c.Status(404).JSON(fiber.Map{"success": false, "message": "User not found"})
	}

	var result bson.M
	err = user.Decode(&result)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "Error decoding user data", "error": err.Error()})
	}

	return c.Status(200).JSON(result)
}