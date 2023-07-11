package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pramessh-P/weather-api/internal/model"
	"github.com/Pramessh-P/weather-api/internal/service"
)

type WeatherHandler struct {
	weatherService service.WeatherService
}

func NewWeatherHandler(weatherService service.WeatherService)*WeatherHandler{
	return &WeatherHandler{
		weatherService: weatherService,
	}
}
func (h *WeatherHandler)GetWeatherHandler(w http.ResponseWriter,r *http.Request){

	city:=r.FormValue("city")
	if city == ""{
		http.Error(w,"city is required field",http.StatusBadRequest)
		return
	}
	weather,err := h.weatherService.GetWeather(city)
	if err!=nil{
		log.Printf("failed to get weather data %v",err)
		http.Error(w,"failed to get weather data",http.StatusInternalServerError)
		return 
	}

	response := struct{

		City string `json:"city"`
		Temperature float64 `json:"temperature"`

	}{
		City:weather.City ,
		Temperature: weather.Temperature,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	err=json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	log.Printf("temperature of the city  %s is >> %f",response.City,response.Temperature)
}
func (h *WeatherHandler) LogWeatherHandler(w http.ResponseWriter, r *http.Request) {
	var weather model.Weather
	err := json.NewDecoder(r.Body).Decode(&weather)
	if err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = h.weatherService.LogWeather(&weather)
	if err != nil {
		log.Printf("Failed to log weather data: %v", err)
		http.Error(w, "Failed to log weather data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}