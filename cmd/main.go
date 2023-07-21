package main

import (
	"log"

	"github.com/Pramessh-P/weather-api/internal/api/handlers"
	"github.com/Pramessh-P/weather-api/internal/database"
	"github.com/Pramessh-P/weather-api/internal/initializers"
	"github.com/Pramessh-P/weather-api/internal/repository"
	"github.com/Pramessh-P/weather-api/internal/service"
	"github.com/Pramessh-P/weather-api/pkg/http"
)

const (

	port = ":8080"
)

func init(){
	initializers.LoadEnvVaraibles()
	database.DBConnection()
	database.SyncDatabase()
}
func main() {
	
	weatherRepo:=repository.NewWeatherRepository()
	weatherService:=service.NewWeatherService(weatherRepo)
	weatherHandler:=handlers.NewWeatherHandler(weatherService)
	userRepo:=repository.NewUserRepository()
	userServ:=service.NewuserService(userRepo)
	userHandler:=handlers.NewUserhandler(userServ)

	server := http.NewServer(port, weatherHandler,userHandler)

	// Start the server
	log.Fatal(server.Start())
}
