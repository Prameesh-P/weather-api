package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Pramessh-P/weather-api/internal/database"
	"github.com/Pramessh-P/weather-api/internal/model"
	"github.com/Pramessh-P/weather-api/internal/service"
	"gorm.io/gorm"
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
	city:=r.URL.Path[len("/weather/"):]
	if city == ""{
		http.Error(w,"city is required field",http.StatusBadRequest)
		return
	}
	weather,err := h.weatherService.GetWeather(city)
	if err!=nil{
		log.Printf("failed to get weather data %v",err)
		http.Error(w,"currently this city is not available",http.StatusInternalServerError)
		return 
	}
	UserData:=""
	usrEmail:=r.FormValue("email")
	response := struct{

		City string `json:"city"`
		Temperature float64 `json:"temperature"`
		UserData string `json:"user_data"`


	}{
		City:weather.City ,
		Temperature: weather.Temperature,
		UserData: UserData,
	}
    var user model.User
    result := database.DB.Where("email = ?", usrEmail).First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
		UserData="please login for store data!!"
		response := struct{

			City string `json:"city"`
			Temperature float64 `json:"temperature"`
			UserData string `json:"user_data"`
		}{
			City:weather.City ,
			Temperature: weather.Temperature,
			UserData: UserData,
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		err=json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("Failed to encode response: %v", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
        }
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
    }
	var userAct model.UserActivity
	userAct.City=weather.City
	userAct.Temperature=weather.Temperature
	userAct.Email=usrEmail
	currentTime:=time.Now()
	layout := "2006-01-02 15:04:05"
	formattedTime := currentTime.Format(layout)
	userAct.Time=formattedTime
	results :=database.DB.Create(&userAct)
	if results.Error != nil {
        http.Error(w,"error from update db "+results.Error.Error(),http.StatusInternalServerError)
		return
    }
	UserData="userdata stored succefully"
	response = struct{
		City string `json:"city"`
		Temperature float64 `json:"temperature"`
		UserData string `json:"user_data"`
	}{
		City:weather.City ,
		Temperature: weather.Temperature,
		UserData: UserData,
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
func (h *WeatherHandler)LogWeatherHandler(w http.ResponseWriter, r *http.Request) {
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
