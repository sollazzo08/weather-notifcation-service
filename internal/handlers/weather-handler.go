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

		lat := query.Get("lat")
		lon := query.Get("lon")

		fmt.Println(lat, lon, "test")
		weatherData, err := service.GetWeather(lat, lon string)
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weatherData)
	}
}
