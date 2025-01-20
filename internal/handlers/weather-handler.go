package handlers

import (
	"encoding/json"
	"github.com/mike/weather-notification-service/internal/weather"
	"net/http"
)

func WeatherHandler(service *weather.WeatherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		weatherData, err := service.GetWeather()
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weatherData)
	}
}
