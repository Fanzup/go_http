package weather_test

import (
	"demo/weather_check/geo"
	"demo/weather_check/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	city := "London"
	geoData := geo.GeoData{
		City: city,
	}
	format := 3
	res, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(res, city) {
		t.Error("City is dont find!")
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 123},
	{name: "Null", format: 0},
	{name: "Minus", format: -1},
	{name: "Minus", format: 1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			city := "Moscow"
			geoData := geo.GeoData{
				City: city,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrorWrongFormat {
				t.Error("Format is good!")
			}
		})
	}

}
