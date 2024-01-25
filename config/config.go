package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config stores the application configuration
type Config struct {
	AllowedOrigins []string
	CorsMaxAge     int // Max Age of Preflight requests
	MongoDbUri     string
	// Add other configuration variables here
}

// LoadConfig loads environment variables from the .env file
// and returns a Config struct
func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	allowedOrigins := getEnvAsSlice("ALLOWED_ORIGINS", ",", []string{"*"})
	mongodbUri := getEnvAsString("MONGODB_URL", "")

	return &Config{
		AllowedOrigins: allowedOrigins,
		CorsMaxAge:     getEnvAsInt("CORS_MAXAGE", 12), // Provide a default value
		MongoDbUri:     mongodbUri,
		// initialize other fields...
	}
}

func getEnvAsSlice(name string, delimiter string, defaultValue []string) []string {
	value, exists := os.LookupEnv(name)
	if !exists {
		return defaultValue
	}
	return strings.Split(value, delimiter)
}

// Helper functions to parse environment variables
func getEnvAsString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnvAsString(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// Add more helper functions as needed for other data types like bool, float, etc.
