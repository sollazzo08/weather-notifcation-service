package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mike/weather-notification-service/internal/weather"
)

func WeatherHandler(service *weather.WeatherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()

		zip := query.Get("zip")

		weatherData, err := service.GetWeatherByZip(zip)
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		fmt.Println(weatherData)
		json.NewEncoder(w).Encode(weatherData)
	}
}
