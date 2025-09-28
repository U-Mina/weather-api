package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "github.com/U-Mina/weather-api/internal/models"
	// "weater-api/internal/models"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, my Go weather api!")
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Checking weather...")
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

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
	http.HandleFunc("/", helloHandler)

	// http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Checking the weather...")
	// })
	http.HandleFunc("/weather", weatherHandler)

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
