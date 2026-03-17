package geo_test

import (
	"demo/weather_check/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange - pre-Test. data
	city := "London"
	res := geo.GeoData{
		City: "London",
	}
	//Act - run Test
	got, err := geo.GetMyLocation(city)
	//Assert - track result from Test
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != res.City {
		t.Errorf("Oops....WAIT! %v is not %v ", res, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "AbobaLOL"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Oops....WAIT! %v is not %v ", geo.ErrorNoCity, err)
	}
}
