package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Pramessh-P/weather-api/internal/model"
	"github.com/Pramessh-P/weather-api/internal/repository"
)

type WeatherService interface {
	GetWeather(city string) (*model.Weather,error)
	LogWeather(weather *model.Weather)error
}

type weatherService struct{
	weatherRepository repository.WeatherRepository
}

func NewWeatherService(weatherRepository repository.WeatherRepository)*weatherService{

	return &weatherService{
		weatherRepository: weatherRepository,
	}
}

func (s *weatherService)GetWeather(city string)(*model.Weather,error){
	Weather,err :=query(city)
	if err !=nil{
		return nil,err
	}
	weather := &model.Weather{
		City:        city,
		Temperature: Weather.Main.Kelvin,
	}

	return weather, nil
}

func (s *weatherService)LogWeather(weather *model.Weather)error{
	err := validateWeatherData(weather)
	if err != nil {
		return fmt.Errorf("invalid weather data: %v", err)
	}

	// Store weather data using the repository
	err = s.weatherRepository.StoreWeather(weather)
	if err != nil {
		return fmt.Errorf("failed to store weather data: %v", err)
	}

	log.Printf("Weather data logged for city '%s'", weather.City)
	return nil
}
func validateWeatherData(weather *model.Weather) error {
	if weather.City==""{
		return fmt.Errorf("city is required")
	}
	return nil
}
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type apiConfig struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

func query(city string )(WeatherData,error){
	apiConfig,err:=loadApiConfig(".apiConfig")
	if err !=nil{
		fmt.Println("erfafasr")
		return WeatherData{},err
	}
	resp,err:=http.Get("http://api.openweathermap.org/data/2.5/weather?APPID="+apiConfig.OpenWeatherMapApiKey+"&q="+city)
	if err !=nil{
		return WeatherData{},err
	}
	defer resp.Body.Close()
	var d WeatherData
	if err=json.NewDecoder(resp.Body).Decode(&d);err !=nil{
		fmt.Println("err")
		return WeatherData{},err
	}
	if d.Name==""{
		return WeatherData{},fmt.Errorf("unavailable city")
	}
	d.Main.Kelvin=kelvinToCelsius(d.Main.Kelvin)
	return d,nil
}
func loadApiConfig(file string) (apiConfig, error) {
	workingDir,_:=os.Getwd()
	fmt.Println(workingDir)
	bytes, err := os.ReadFile(file)
	if err !=nil{
		fmt.Println("err")
		return apiConfig{} ,err
	}
	fmt.Println()
	var c apiConfig
	err=json.Unmarshal(bytes,&c);
	if err!=nil{
		fmt.Println("json erro")
		return apiConfig{} ,err
	}
	return c,nil
}
func kelvinToCelsius(kelvin float64) float64 {
	celsius:=kelvin - 273.15
	two:=fmt.Sprintf("%.3f",celsius)
	f, err := strconv.ParseFloat(two, 64)
	if err != nil {
		fmt.Println("Error:", err)
		
	}
	return f
}