package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

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
	
	if resp.StatusCode != 200 {
		fmt.Println("Error: Response code not 200")
	} else {
		fmt.Println("Success: Response Code 200")
	}

	json, err := io.ReadAll(resp.Body)

	fmt.Println(string(json))

}
