package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	APIKey string
}

// LoadConfig starts with a captial letter here so we are able to import it from other files
// Returning a pointer which contains the address of the Config "object"
// In Go, we loosely refer to structs with behavior (methods) as objects

// A pointer is used for the following reasons:
// 1. Efficiency: Avoids copying the entire Config struct, which can be costly if it contains many fields.
// 2. Mutability: Allows the caller to modify the Config object if needed, and those changes persist outside the function.
// 3. Consistency: Returning a pointer is a common practice for shared or reusable resources like configuration data.
func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, fmt.Errorf("error loading .env file")
	}

	// Read environment variables
	apiKey := os.Getenv("API_KEY")

	// Validate required variables
	if apiKey == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}
	// Use the & operator to get the memory address of Config and return it as a pointer.
	// The & sign is required because:
	// 1. The return type of the function is *Config, which expects a pointer.
	// 2. Config is a struct value, and we need to explicitly convert it to a pointer.
	// 3. Returning a pointer avoids copying the entire struct, improving efficiency,
	//    especially for larger structs or shared resources like configurations.
	return &Config{
		APIKey: apiKey,
	}, nil
}
