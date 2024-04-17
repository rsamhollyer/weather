This is my first little Go program. It's a rewrite of [this](https://github.com/polybar/polybar-scripts/blob/master/polybar-scripts/openweathermap-simple/openweathermap-simple.sh) polybar module in bash.

![screenshot](assets/Bar.png)

It's a simple weather module for polybar that uses the OpenWeatherMap API. It displays the current weather and temperature for a given city. It also displays the current weather icon. I use the city ID to get the weather data, but you can also use the city name. Just make adjustments to the *url.go* and *config.go* files. 

Update the *config.go* file with your OpenWeatherMap API key and city ID. You can get your API key [here](https://openweathermap.org/api). You can get your city ID [here](http://bulk.openweathermap.org/sample/).