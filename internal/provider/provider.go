package provider

import "github.com/U-Mina/weather-api/internal/models"

// WeatherProvider define behaviour of any service that can get current weather data
type WeatherProvider interface {
	FetchCurrentWeather(city string) (*models.WeatherData, err)
}
