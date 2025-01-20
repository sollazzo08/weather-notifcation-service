package main

import (
	"fmt"
	"github.com/mike/weather-notification-service/internal/config"
	"github.com/mike/weather-notification-service/internal/handlers"
	"github.com/mike/weather-notification-service/internal/weather"
	"log"
	"net/http"
)

// handleHelloWorld is a handler function that responds to requests with "Hello, World!".
// w - http.ResponseWriter: Used to send a response back to the client.
// r - *http.Request: Represents the incoming HTTP request.
// func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!") // Write a simple response to the client
// }

func main() {
	fmt.Println("Starting Weather Notification Service...")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	mux := http.NewServeMux()

	weatherService := weather.NewWeatherService(cfg.APIKey, cfg.Longitude, cfg.Latitude)

	mux.HandleFunc("/api/v1/weather", handlers.WeatherHandler(weatherService))

	// Start the server
	fmt.Println("Server is running on http://localhost:8090")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
