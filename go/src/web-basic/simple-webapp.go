package main

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)

type weatherData struct{
	Name string `json:"name"`
	Main struct {
		kelvin float64 `json:"temp"`
	}  `json:main`
}


func query(city string) (weatherData, error){
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=YOUR_API_KEY&q=" + city)
	if err != nil{
		return weatherData{}, err
	}
	defer resp.Body.Close()
	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil{
		return weatherData{}, err
	}
	return d, nil
}

func handle(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "successfully set")
}

func weatherInfo(w http.ResponseWriter, r *http.Request){
	city := "tokyo"
	data, err := query(city)
	fmt.Print(data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}


func main(){
	http.HandleFunc("/", handle)
	http.HandleFunc("/weather/", weatherInfo)
	http.ListenAndServe(":8000", nil)
}