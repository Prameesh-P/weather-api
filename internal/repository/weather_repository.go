package repository

import (
	"log"

	"github.com/Pramessh-P/weather-api/internal/model"
)

type WeatherRepository interface {
	StoreWeather(weather *model.Weather)error
}

type WeatherInMemoryRepository struct{
	weatherData map[string]*model.Weather
}

func NewWeatherRepository()*WeatherInMemoryRepository{
	return &WeatherInMemoryRepository{
		weatherData: make(map[string]*model.Weather),
	}
}

func (r *WeatherInMemoryRepository)StoreWeather(weather *model.Weather)error{
	r.weatherData[weather.City]=weather
	log.Printf("Stored weather for the city %s is : %f degree C",weather.City,weather.Temperature)
	return nil
}