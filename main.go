package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type apiConfig struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(file string) (apiConfig, error) {
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
func main() {
	http.HandleFunc("/weather/",func(w http.ResponseWriter, r *http.Request) {
		city :=strings.SplitN(r.URL.Path,"/",3)[2]
		data,err := query(city)
		
		if err !=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":8080",nil)
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