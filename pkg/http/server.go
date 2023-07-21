package http

import (
	"fmt"
	"net/http"

	"github.com/Pramessh-P/weather-api/internal/api/handlers"
)

type Server struct {
	addr           string
	weatherHandler *handlers.WeatherHandler
	userHandler *handlers.UserHandler
}

func NewServer(addr string, weatherHandler *handlers.WeatherHandler,user *handlers.UserHandler) *Server {
	return &Server{
		addr:           addr,
		weatherHandler: weatherHandler,
		userHandler: user,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Define the API routes
	mux.HandleFunc("/weather/", s.weatherHandler.GetWeatherHandler)
	mux.HandleFunc("/log", s.weatherHandler.LogWeatherHandler)
	mux.HandleFunc("/signup",s.userHandler.Signup)
	mux.HandleFunc("/recent",s.userHandler.GetRecentAction)
	mux.HandleFunc("/login",s.userHandler.LoginUser)

	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	fmt.Printf("Server listening on %s\n", s.addr)
	return server.ListenAndServe()
}
