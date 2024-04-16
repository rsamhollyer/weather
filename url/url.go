package url

import (
	"net/url"
	"weather/config"
)

func CreateUrl() (string, error) {
	weatherApi, err := url.Parse("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		return "", err
	}
	config := config.NewConfig()
	weatherApiQuery := weatherApi.Query()
	weatherApiQuery.Set("appid", config.GetApiKey())
	weatherApiQuery.Set("id", config.GetCityId())
	weatherApiQuery.Set("units", config.GetUnits())
	weatherApi.RawQuery = weatherApiQuery.Encode()

	return weatherApi.String(), nil
}
