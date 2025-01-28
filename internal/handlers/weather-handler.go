package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mike/weather-notification-service/internal/weather"
	"net/http"
)

func WeatherHandler(service *weather.WeatherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()

		zip := query.Get("zip")

		weatherData, err := service.GetWeatherByZip(zip)
		if err != nil {
			// Check if the error is an HTTPError
			var httpErr *weather.HTTPError
			if errors.As(err, &httpErr) {
				// Write custom error details in the response body
				w.WriteHeader(httpErr.StatusCode)
				fmt.Fprintf(w, "%s\n", httpErr.Error())
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")

		fmt.Println(weatherData)
		json.NewEncoder(w).Encode(weatherData)
	}
}
