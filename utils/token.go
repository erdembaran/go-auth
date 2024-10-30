package utils

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})
	secret := []byte(os.Getenv("JWT_SECRET"))
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func SetTokenCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,     
		SameSite: "Lax",    
		Expires:  time.Now().Add(2 * time.Hour), 
	})
}

func ClearTokenCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0), // Sets the expiration to the past
		HTTPOnly: true,
		Secure:   true,      // Same as when the cookie was set
		SameSite: "Lax",     // Same as when the cookie was set
	})
}

func VerifyToken(tokenString string) (bool, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}


func GenerateRandomToken() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}