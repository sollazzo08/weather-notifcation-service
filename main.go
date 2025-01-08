package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// handleHelloWorld is a handler function that responds to requests with "Hello, World!".
// w - http.ResponseWriter: Used to send a response back to the client.
// r - *http.Request: Represents the incoming HTTP request.
// func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!") // Write a simple response to the client
// }

type WeatherResponse struct {
    Current struct {
        Time       int64   `json:"dt"`
        Temp       float64  `json:"temp"`
        WindSpeed  float64  `json:"wind_speed"`
    } `json:"current"`
}

var units = "imperial"
var exclude = "minutely"

func main() {
	fmt.Println("Starting Project...")

    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    apiKey := os.Getenv("API_KEY")
    latitude := os.Getenv("LAT")
    longitude := os.Getenv("LON")

	// Create a new ServeMux (multiplexer) using the short variable declaration (:=).
	// ServeMux is a request router that maps URL paths to handler functions.
	// The inferred type of mux is *http.ServeMux (a pointer to a ServeMux instance).
	// mux := http.NewServeMux()

	// Register the "/hello" route with the handleHelloWorld function.
	// When a request is made to "/hello", the server will execute handleHelloWorld.
	//mux.HandleFunc("/hello", handleHelloWorld)

    // Another approach to concat the string is to use fmt.SPrintf
    response, err := http.Get("https://api.openweathermap.org/data/3.0/onecall?lat=" + latitude + "&lon=" + longitude + "&units=" + units + "&exclude=" + exclude + "&appid=" + apiKey);



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

    var weatherResponseObj WeatherResponse

    json.Unmarshal([]byte(responseData), &weatherResponseObj)

    fmt.Println(weatherResponseObj.Current.Temp)
    fmt.Println(weatherResponseObj.Current.Time)
    fmt.Println(weatherResponseObj.Current.WindSpeed)

    // TODO unmarshall data into structs

	// Start the HTTP server on port 8080.
	// http.ListenAndServe binds the server to the specified port and uses the mux
	// as the handler to route incoming requests.
	// Passing mux as a pointer ensures any state changes (if needed) affect the server.
	// fmt.Println("Listening on http://localhost:8080")
	// http.ListenAndServe(":8080", mux)
}
