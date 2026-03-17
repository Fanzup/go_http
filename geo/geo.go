package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulation struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, errors.New("City incorrect!")
		}
		return &GeoData{
			City: city,
		}, nil
	}
	res, err := http.Get("http://ip-api.com/json")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("Not 200!")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	res, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Print(err.Error())
		return false

	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false
	}
	var population CityPopulation
	json.Unmarshal(body, &population)
	return !population.Error

}
