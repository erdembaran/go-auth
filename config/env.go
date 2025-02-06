package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("ENV") == "production" {
		log.Println("Running in production mode, skipping .env file load")
		return
	}
	
	envFile := ".env.local"
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: Error loading %s file: %v", envFile, err)
		log.Println("Using environment variables from system")
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}