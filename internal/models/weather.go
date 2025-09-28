package models

// exported file, start with Capitalized
// `json:"content"`
type WeatherData struct {
	City        string  `json:"location"`
	Temperature float64 `json:"temperature"`
	TempMax		float64 `json:"temp_max"`
	TempMin		float64 `json:"temp_min"`
	Humidity	int		`json:"humidity"`
	Description string  `json:"description"`
}
