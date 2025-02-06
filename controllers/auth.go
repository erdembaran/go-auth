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

	return c.JSON(fiber.Map{"user": user, "message": "Logged in successfully!"})
}


func Logout(c *fiber.Ctx) error {
	utils.ClearTokenCookie(c)
	return c.JSON(fiber.Map{"message": "Logged out successfully!"})
}

func ForgotPassword(c *fiber.Ctx) error {
	// parse the incoming request
	var req models.ForgotPasswordRequest
	if err := c.BodyParser(&req); err != nil || req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// to find a user with the specified email 
	var user models.User
	filter := bson.M{"email": req.Email}
	err := database.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// generate a secure random token
	resetToken := utils.GenerateRandomToken()
	resetTokenExpiresAt := primitive.NewDateTimeFromTime(time.Now().Add(10 * time.Minute))


	// update the user with the reset token and expiry date
	update := bson.M{
		"$set": bson.M{
			"resetPasswordToken":   resetToken,
			"resetPasswordExpiresAt": resetTokenExpiresAt,
		},
	}

	// update the user in the database
	_, err = database.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set reset token"})
	}

	// send the reset link to the user's email
	resetLink := "http://localhost:3000/reset-password?token=" + resetToken
	emailBody := "Click the link to reset your password: " + resetLink
	err = utils.SendEmail(req.Email, "Password Reset", emailBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
	}

	return c.JSON(fiber.Map{"message": "Password reset link sent to your email."})
}

func ResetPassword(c *fiber.Ctx) error {
	// parse the incoming request to get the new password
	var req *models.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil || req.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "New password is required"})
	}

	// retrieve the token from parameters
	token := c.Params("token")
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token is missing"})
	}

	// find user by reset token and check expiration
	var user models.User
	filter := bson.M{
		"resetPasswordToken": token,
		"resetPasswordExpiresAt": bson.M{
			"$gt": primitive.NewDateTimeFromTime(time.Now()), // ensure token is not expired
		},
	}
	err := database.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// hash the new password
	hashedPassword := utils.GeneratePassword(req.NewPassword)
	if hashedPassword == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// update the user's password and clear the reset token fields
	update := bson.M{
		"$set": bson.M{"password": hashedPassword},
		"$unset": bson.M{
			"resetPasswordToken":     nil,
			"resetPasswordExpiresAt": nil,
		},
	}
	_, err = database.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to reset password"})
	}

	return c.JSON(fiber.Map{"message": "Password reset successfully!"})
}


  




