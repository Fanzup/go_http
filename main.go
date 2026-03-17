package main

import (
	"demo/weather_check/geo"
	"demo/weather_check/weather"
	"flag"
	"fmt"
)

func main() {
	a := 1337
	fmt.Println(a)
	city := flag.String("city", "", "User`s city ")
	format := flag.Int("format", 1, "Format for print weather")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData, _ := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
