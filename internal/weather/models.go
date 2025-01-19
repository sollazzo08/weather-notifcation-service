package weather

// WeatherResponse represents the JSON response from the OpenWeatherAPI.
type WeatherResponse struct {
	Current struct {
		Time      int64   `json:"dt"`
		Temp      float64 `json:"temp"`
		WindSpeed float64 `json:"wind_speed"`
	} `json:"current"`
}
