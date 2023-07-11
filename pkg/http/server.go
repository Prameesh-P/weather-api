package http

import (
	"fmt"
	"net/http"

	"github.com/Pramessh-P/weather-api/internal/api/handlers"
)

type Server struct {
	addr           string
	weatherHandler *handlers.WeatherHandler
}

func NewServer(addr string, weatherHandler *handlers.WeatherHandler) *Server {
	return &Server{
		addr:           addr,
		weatherHandler: weatherHandler,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Define the API routes
	mux.HandleFunc("/weather", s.weatherHandler.GetWeatherHandler)
	mux.HandleFunc("/log", s.weatherHandler.LogWeatherHandler)

	server := &http.Server{
		Addr:    s.addr,
		Handler: mux,
	}

	fmt.Printf("Server listening on %s\n", s.addr)
	return server.ListenAndServe()
}
