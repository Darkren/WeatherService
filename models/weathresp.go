package models

// WeatherResponse is a response containing
// needed weather info
type WeatherResponse struct {
	ID          int64
	RequestID   int64
	Temperature int
	Humidity    int
	Pressure    int
}
