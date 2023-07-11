package main

import (
	"log"

	"github.com/Pramessh-P/weather-api/internal/api/handlers"
	"github.com/Pramessh-P/weather-api/internal/repository"
	"github.com/Pramessh-P/weather-api/internal/service"
	"github.com/Pramessh-P/weather-api/pkg/http"
)

const (

	port = ":8080"
)
func main() {
	weatherRepo:=repository.NewWeatherRepository()
	weatherService:=service.NewWeatherService(weatherRepo)
	weatherHandler:=handlers.NewWeatherHandler(weatherService)
	server := http.NewServer(port, weatherHandler)

	// Start the server
	log.Fatal(server.Start())
}
