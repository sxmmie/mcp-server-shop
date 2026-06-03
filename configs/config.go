package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL    string // http://localhost:8080/api/v1
	AuthToken string // JWT for the currently logged in user
	LogLevel  string
	Transport string // stdio -> no HTTP for now
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		authToken = os.Getenv("JWT_TOKEN")
	}

	// build config
	config := &Config{
		APIURL:    getEnv("API_URL", "http://localhost:8080"),
		AuthToken: authToken,
		LogLevel:  getEnv("LOG_LEVEL", "debug"),
		Transport: getEnv("TRANSPORT", "stdio"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
