package controllers

import (
	"context"
	"time"

	"github.com/erdembaran/go-auth/database"
	"github.com/erdembaran/go-auth/models"
	"github.com/erdembaran/go-auth/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Check if required fields are provided
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username, email, and password are required",
		})	
	}

	// Check if user already exists
	filter := bson.M{"email": user.Email}
	count, err := database.Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already in use."})
	}

	// Hash the password
	user.Password = utils.GeneratePassword(user.Password)
	if user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Set timestamps
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	// Insert new user
	result, err := database.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(fiber.Map{"user": user})
}

func Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Check if required fields are provided
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// Find user by email
	var user models.User
	filter := bson.M{"email": req.Email}
	err := database.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password."})
	}

	// Verify password
	if !utils.ComparePassword(user.Password, req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password."})
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	utils.SetTokenCookie(c, token)

	return c.JSON(fiber.Map{"token": token, "user": user, "message": "Logged in successfully!"})
}


func Logout(c *fiber.Ctx) error {
	utils.ClearTokenCookie(c)
	return c.JSON(fiber.Map{"message": "Logged out successfully!"})
}


