package auth

import (
	"demo/weather_check/configs"
	"demo/weather_check/packages/response"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewHelloHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// read body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			response.JsonResponse(w, err.Error(), 402)
		}
		if payload.Email == "" || payload.Password == "" {
			response.JsonResponse(w, "Email or Password incorrect!", 403)
			return
		}
		_, err = mail.ParseAddress(payload.Email)
		if err != nil {
			response.JsonResponse(w, "Email or Password incorrect!", 403)
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		response.JsonResponse(w, data, 200)
	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
