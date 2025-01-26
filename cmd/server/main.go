package main

import (
	"fmt"
	"github.com/mike/weather-notification-service/internal/config"
	"github.com/mike/weather-notification-service/internal/handlers"
	"github.com/mike/weather-notification-service/internal/weather"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Weather Notification Service...")
	// Create a new instance of http.ServeMux using the NewServeMux constructor function
	mux := http.NewServeMux()
	// Load our config file on start up
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	weatherService := weather.NewWeatherService(cfg.APIKey)

	mux.HandleFunc("/api/v1/weather", handlers.WeatherHandler(weatherService))

	// Start the server
	fmt.Println("Server is running on http://localhost:8090")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
