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
		t.Errorf("Oops....%v is not %v ", res, got)
	}
}
