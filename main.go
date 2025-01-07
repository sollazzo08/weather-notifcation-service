package main

import (
	"fmt"
	"net/http"
)

// handleHelloWorld is a handler function that responds to requests with "Hello, World!".
// w - http.ResponseWriter: Used to send a response back to the client.
// r - *http.Request: Represents the incoming HTTP request.
func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // Write a simple response to the client
}

func main() {
	fmt.Println("Starting Project...")

	// Create a new ServeMux (multiplexer) using the short variable declaration (:=).
	// ServeMux is a request router that maps URL paths to handler functions.
	// The inferred type of mux is *http.ServeMux (a pointer to a ServeMux instance).
	mux := http.NewServeMux()

	// Register the "/hello" route with the handleHelloWorld function.
	// When a request is made to "/hello", the server will execute handleHelloWorld.
	mux.HandleFunc("/hello", handleHelloWorld)

	// Start the HTTP server on port 8080.
	// http.ListenAndServe binds the server to the specified port and uses the mux
	// as the handler to route incoming requests.
	// Passing mux as a pointer ensures any state changes (if needed) affect the server.
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
