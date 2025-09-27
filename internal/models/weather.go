package models

// exported file, start with Capitalized
// `json:"content"`
type WeatherData struct {
	City        string  `json:"location"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}
