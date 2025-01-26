package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherService is a struct that encapsulates the configuration for interacting with OpenWeatherAPI.
type WeatherService struct {
	APIKey string
}

// NewWeatherService creates and initializes a new WeatherService instance.
// This constructor function takes the API key, latitude, and longitude as inputs
// and returns a pointer to the initialized WeatherService struct.
//
// Parameters:
// - apiKey: The API key for authenticating with the weather API.
// - latitude: The latitude of the location to fetch weather data for.
// - longitude: The longitude of the location to fetch weather data for.
//
// Returns:
// - *WeatherService: A pointer to the initialized WeatherService instance.
func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		APIKey: apiKey,
	}
}

func (s *WeatherService) GetWeatherByZip(zip string) (*WeatherResponse, error) {
	// This API is depreciated, need to look into a new way for retrieving weather by city name
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + zip + ",US" +
		"&units=imperial&appid=" + s.APIKey

	fmt.Println(url)
	resp, err := http.Get(url)

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
