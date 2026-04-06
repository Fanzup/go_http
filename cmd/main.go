package main

import (
	"demo/weather_check/configs"
	"demo/weather_check/internal/auth"
	"demo/weather_check/internal/auth/link"
	"demo/weather_check/packages/db"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	//Handler
	auth.NewHelloHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081!")
	server.ListenAndServe()
}
