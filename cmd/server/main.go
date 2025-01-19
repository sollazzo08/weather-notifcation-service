package main

import (
	"encoding/json"
	"fmt"
	"github.com/mike/weather-notification-service/internal/config"
	"github.com/mike/weather-notification-service/internal/handlers"
	"io"
	"net/http"
	"os"
	"log"
)

// handleHelloWorld is a handler function that responds to requests with "Hello, World!".
// w - http.ResponseWriter: Used to send a response back to the client.
// r - *http.Request: Represents the incoming HTTP request.
// func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!") // Write a simple response to the client
// }

type WeatherResponse struct {
	Current struct {
		Time      int64   `json:"dt"`
		Temp      float64 `json:"temp"`
		WindSpeed float64 `json:"wind_speed"`
	} `json:"current"`
}

var units = "imperial"
var exclude = "minutely"

func main() {
	fmt.Println("Starting Project...")

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/weather", handlers.WeatherHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}




	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	apiKey := config.APIKey
	latitude := config.Latitude
	longitude := config.Longitude

	// Create a new ServeMux (multiplexer) using the short variable declaration (:=).
	// ServeMux is a request router that maps URL paths to handler functions.
	// The inferred type of mux is *http.ServeMux (a pointer to a ServeMux instance).
	// mux := http.NewServeMux()

	// Register the "/hello" route with the handleHelloWorld function.
	// When a request is made to "/hello", the server will execute handleHelloWorld.
	//mux.HandleFunc("/hello", handleHelloWorld)

	// Another approach to concat the string is to use fmt.SPrintf
	response, err := http.Get("https://api.openweathermap.org/data/3.0/onecall?lat=" + latitude + "&lon=" + longitude + "&units=" + units + "&exclude=" + exclude + "&appid=" + apiKey)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(string(responseData))

	// Creates an empty instance of the WeatherResponse struct
	var weatherResponseObj WeatherResponse

	// By passing the memory address of the weatheRepsonse obj we can update the weatherResponseObj, fill it with the JSON data
	json.Unmarshal([]byte(responseData), &weatherResponseObj)

	fmt.Println(weatherResponseObj.Current.Temp)
	fmt.Println(weatherResponseObj.Current.Time)
	fmt.Println(weatherResponseObj.Current.WindSpeed)

	// Start the HTTP server on port 8080.
	// http.ListenAndServe binds the server to the specified port and uses the mux
	// as the handler to route incoming requests.
	// Passing mux as a pointer ensures any state changes (if needed) affect the server.
	// fmt.Println("Listening on http://localhost:8080")
	// http.ListenAndServe(":8080", nil)
}
