package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	APIKey    string
	Latitude  string
	Longitude string
}

// LoadConfig starts with a captial letter here so we are able to import it from other files
func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, fmt.Errorf("error loading .env file")
	}

	// Read environment variables
	apiKey := os.Getenv("API_KEY")
	latitude := os.Getenv("LAT")
	longitude := os.Getenv("LON")

	// Validate required variables
	if apiKey == "" || latitude == "" || longitude == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return &Config{
		APIKey:    apiKey,
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
