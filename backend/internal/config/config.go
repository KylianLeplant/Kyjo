package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration loaded from environment variables.
type Config struct {
	DatabaseURL string
	Port        string
}

// Load reads environment variables and returns a Config.
// It attempts to load a .env file when present but does not fail if it is missing.
func Load() (*Config, error) {
	// Try loading .env from a few common locations.
	_ = godotenv.Load(".env", "../.env", "../../.env", "../../../.env")

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DatabaseURL: databaseURL,
		Port:        port,
	}, nil
}
