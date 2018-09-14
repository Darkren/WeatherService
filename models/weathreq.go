package models

import (
	"time"
)

// WeatherRequest is an incoming request for weather
type WeatherRequest struct {
	ID      int64
	Lat     float32
	Long    float32
	Created time.Time
	IsDone  bool
}
