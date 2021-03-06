package models

import (
	"time"
)

// WeatherRequest is an incoming request for weather
type WeatherRequest struct {
	ID           int64
	Lat          float64
	Lon          float64
	Created      time.Time
	IsComplete   bool
	IsInProgress bool
}
