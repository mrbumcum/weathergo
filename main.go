package main

import (
	"encoding/json"
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error has occured", err)
	}

	var weather Weather
	json.Unmarshal(data, &weather)

	struct_location, current, _ := weather.Location, weather.Current, weather.Forecast

	message := fmt.Sprintf(
		"%s, %s  %0.2f, %s",
		struct_location.Name,
		struct_location.Country,
		current.TempF,
		current.Condition.Text,
	)

	fmt.Println(message)



}
