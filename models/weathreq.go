package models

import (
	"time"
)

// WeatherRequest is an incoming request for weather
type WeatherRequest struct {
	ID         int64
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Created    time.Time
	IsComplete bool
}
