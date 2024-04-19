package main

import (
	"fmt"
	"os"
	"weather/getweather"
	"weather/transform"
	"weather/url"
)

func main() {
	weatherurl, error := url.CreateUrl()

	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}

	weatherObject, error := getweather.GetWeather(weatherurl)
	if error != nil {
		fmt.Print(error.Error())
		os.Exit(1)
	}
	weatherString := transform.Transform(&weatherObject)
	fmt.Println(weatherString)
}
