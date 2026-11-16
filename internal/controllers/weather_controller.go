package controllers

import (
	"encoding/json"
	"lesson49/internal/models"
	"math/rand/v2"
	"net/http"
	"time"
)

type WeatherController struct {
}

func (controller *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(newWeathers())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func newWeathers() []models.Weather {
	weathers := []models.Weather{}

	for i := 0; i < 24; i++ {
		weathers = append(weathers, models.Weather{
			TemperatureC: rand.IntN(70) - 20,
			Time:         time.Now().Add(time.Hour * time.Duration(i)),
		})
	}

	return weathers
}
