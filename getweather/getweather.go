package getweather

import (
	"encoding/json"
	"io"
	"net/http"
)

type WeatherIcon struct {
	Icon string `json:"icon"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

type Weather struct {
	WeatherIcon []WeatherIcon `json:"weather"`
	Main        Main          `json:"main"`
}

func GetWeather(weatherUrl string) (Weather, error) {
	response, err := http.Get(weatherUrl)
	var weatherObject Weather
	if err != nil {
		return weatherObject, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return weatherObject, err
	}

	err = json.Unmarshal(responseData, &weatherObject)

	if err != nil {
		return weatherObject, err
	}

	return weatherObject, nil
}
