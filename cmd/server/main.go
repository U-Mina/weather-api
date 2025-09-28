package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/U-Mina/weather-api/internal/models"
	"github.com/joho/godotenv"
	// "weather-api/internal/models"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, my Go weather api!")
}

func weatherHandler(w http.ResponseWriter, r *http.Request, apiKey string) {
	// fmt.Fprintf(w, "Checking weather...")
	// apiKey := os.Getenv("API_KEY")
	// if apiKey == "" {
	// 	http.Error(w, "API key not set", http.StatusInternalServerError)
	// 	return
	// }
	city := "Heilbronn"
	// 9f4ba2e32be02a0368d1afa4efefbf69
	requestURL := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city,
		apiKey,
	)

	/** get request to open-weather
	if response, err := http.Get(requestURL); err != nil {
		http.Error(w, "Fail to get from openWeather!", http.StatusInternalServerError)
		return
	}
	defer response.Body.close()
	=> wrong syntax, 'response' is within scope! so defer not working
	*/

	response, err := http.Get(requestURL)
	if err != nil {
		http.Error(w, "Fail to fetch data from OpenWeather!", http.StatusInternalServerError)
		return
	}

	// IMPORTANT: close connection when func finished!
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Fail to read response body", http.StatusInternalServerError)
		return
	}

	var apiResponse struct {
		Name string `json:"name"`
		Main struct {
			Temp float64 `json:"temp"`
			TempMax float64 `json:"temp_max"`
			TempMin float64 `json:"temp_min"`
			Humidity int `json:"humidity"`
			// Pressure int `json:"pressure`

		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		/*
		array in json maps to slice[] in GO
		Weather []struct ==> declared a slice of structs
		JSON OBJ -> Go 'struct'
		JSON array -> Go []struct (or '[]T' for any type)
		*/
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		http.Error(w, "Fail to parse weather data from JSON", http.StatusInternalServerError)
		return
	}

	data := models.WeatherData {
		City: apiResponse.Name,
		Temperature: apiResponse.Main.Temp,
		TempMax: apiResponse.Main.TempMax,
		TempMin: apiResponse.Main.TempMin,
		Humidity: apiResponse.Main.Humidity,
		Description: apiResponse.Weather[0].Description,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Fail to encode response", http.StatusInternalServerError)
	}
	// w.Write(body)

	// fmt.Fprintf(w, "Request URL: %s", requestURL)
	// create sample data using struct
	// data := models.WeatherData{
	// 	City:        "Helbronn",
	// 	Temperature: 13,
	// 	Description: "Rainy",
	// }

	//set content-type header to application/json
	// w.Header().Set("Content-Type", "application/json")

	// convert struct data to json
	// err := json.NewEncoder(w).Encode(data)
	// if err != nil {
	// 	http.Error(w, "Fail to convert to json", http.StatusInternalServerError)
	// }
	// if err := json.NewEncoder(w).Encode(data); err != nil {
	// 	http.Error(w, "Fail to convert to json", http.StatusInternalServerError)
	// }
}

func main() {
	// HandleFunc maps a URL path to a function.
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, Weather API!")
	// 	// Q: why use Fprintf()?? (compare with println/printf etc)
	// })

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env not found")
	}

	if apiKey := os.Getenv("API_KEY"); apiKey == "" {
		log.Fatal("API_KEY not set in .env")
	}

	http.HandleFunc("/", helloHandler)

	// http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Checking the weather...")
	// })
	// http.HandleFunc("/weather", weatherHandler)
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		// Call your original handler, passing the key as an argument
		weatherHandler(w, r, apiKey)
	})

	/* systax meaning of func(w, ...) {}
	   : when a request from '/', run this func (namely: func(w ,...))
	   : w http.ResponseWrite: tool to write response BACK to user
	   : 'w' is just a var name by convention, shortenhand for 'ResponseWriter', can change to whatever
	   : r *http.Request: holds all info about Incoming request (headers, form data...)
	*/

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
	/* http,ListenAndServer
	   : ":8080": This tells the server to listen on port 8080.
	   : nil: This argument is for a router. By passing nil,
	   : we're telling it to use Go's default router, ie, the http.HandleFunc automatically uses.
	*/

}
