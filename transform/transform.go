package transform

import (
	"fmt"
	"weather/getweather"
)

func Transform(weatherObject getweather.Weather) string {
	weatherIcon := transformIcon(weatherObject.WeatherIcon[0])
	temperature := fmt.Sprintf("%.0f", weatherObject.Main.Temp)
	return weatherIcon + "  " + temperature + "°F"

}

func transformIcon(weaterIcon getweather.WeatherIcon) string {
	weatherMap := map[string]string{
		"01d": "",
		"01n": "",
		"02d": "",
		"02n": "",
		"03d": "",
		"03n": "",
		"04d": "",
		"04n": "",
		"09d": "",
		"09n": "",
		"10d": "",
		"10n": "",
		"11d": "",
		"11n": "",
		"13d": "",
		"13n": "",
		"50d": "",
		"50n": "",
	}

	weatherIcon, ok := weatherMap[weaterIcon.Icon]
	if !ok {
		return ""
	}
	return weatherIcon
}
