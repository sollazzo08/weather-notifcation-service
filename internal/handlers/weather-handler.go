package handlers

import (
	"fmt"
	"net/http"
)


func WeatherHandler (w http.ResponseWriter , r *http.Request){

	fmt.Println("Weather Handler")
}
