package request

import (
	"demo/weather_check/packages/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, reader *http.Request) (*T, error) {
	body, err := Decode[T](reader.Body)
	if err != nil {
		response.JsonResponse(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		response.JsonResponse(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
