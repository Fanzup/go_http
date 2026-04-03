package auth

import (
	"demo/weather_check/configs"
	"demo/weather_check/packages/response"
	"fmt"
	"net/http"
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
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
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
