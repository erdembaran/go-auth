package controllers

import (
	"context"

	"github.com/erdembaran/go-auth/database"
	"github.com/erdembaran/go-auth/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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