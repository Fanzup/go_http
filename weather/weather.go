package weather

import (
	"demo/weather_check/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrorWrongFormat = errors.New("Format is wrong!")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("ERROR URL")
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	res, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Print(err.Error())
		return "", errors.New("ERROR HTTP")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
		return "", errors.New("ERROR READ BODY")
	}

	return string(body), nil
}
