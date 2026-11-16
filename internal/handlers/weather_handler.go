package handlers

import (
	"lesson49/internal/controllers"
	"net/http"
)

type WeatherHandler struct {
	controller controllers.WeatherController
}

func NewWeatherHandler() *WeatherHandler {
	return &WeatherHandler{}
}

func (handler *WeatherHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /weather", handler.controller.GetWeather)
}
