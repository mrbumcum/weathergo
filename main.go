package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempF float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch    int64   `json:"time_epoch"`
				TempF        float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var location string 

	if len(os.Args) >= 1 {
		location = os.Args[1]
	} else {
		log.Fatal("Please provide a location")
	}

	WEATHER_API := os.Getenv("WEATHER_API")

	resp, err := http.Get(
		"http://api.weatherapi.com/v1/current.json?key=" +
		WEATHER_API + 
		"&q=" + 
		location + 
		"&aqi=no",
	)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		defer resp.Body.Close()
	}
	
	// if resp.StatusCode != 200 {
	// 	fmt.Println("Error: Response code not 200")
	// } else {
	// 	fmt.Println("Success: Response Code 200")
	// }

	json, err := io.ReadAll(resp.Body)

	fmt.Println(string(json))


}
