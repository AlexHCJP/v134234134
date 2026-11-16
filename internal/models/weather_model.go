package models

import "time"

type Weather struct {
	Time         time.Time `json:"time"`
	TemperatureC int       `json:"temperature"`
}
