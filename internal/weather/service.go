package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherService is a struct that encapsulates the configuration for interacting with OpenWeatherAPI.
type WeatherService struct {
	APIKey    string
	Latitude  string
	Longitude string
}

func NewWeatherService(apiKey, latitude, longitude string) *WeatherService {
	return &WeatherService{
		APIKey:    apiKey,
		Latitude:  latitude,
		Longitude: longitude,
	}
}
func (s *WeatherService) GetWeather() (*WeatherResponse, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&units=imperial&exclude=minutely&appid=%s",
		s.Latitude, s.Longitude, s.APIKey,
	)

	resp, err := http.Get(url)

	fmt.Println(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP response: %d", resp.StatusCode)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}

	return &weather, nil
}

// // Another approach to concat the string is to use fmt.SPrintf
// response, err := http.Get("https://api.openweathermap.org/data/3.0/onecall?lat=" + latitude + "&lon=" + longitude + "&units=" + units + "&exclude=" + exclude + "&appid=" + apiKey)

// if err != nil {
// 	fmt.Println(err.Error())
// 	os.Exit(1)
// }

// responseData, err := io.ReadAll(response.Body)

// if err != nil {
// 	fmt.Println(err.Error())
// 	return
// }

// //fmt.Println(string(responseData))

// // Creates an empty instance of the WeatherResponse struct
// var weatherResponseObj WeatherResponse

// // By passing the memory address of the weatheRepsonse obj we can update the weatherResponseObj, fill it with the JSON data
// json.Unmarshal([]byte(responseData), &weatherResponseObj)

// fmt.Println(weatherResponseObj.Current.Temp)
// fmt.Println(weatherResponseObj.Current.Time)
// fmt.Println(weatherResponseObj.Current.WindSpeed)
