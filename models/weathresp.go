package models

// WeatherResponse is a response containing
// needed weather info
type WeatherResponse struct {
	ID          int64 `json:"id"`
	RequestID   int64 `json:"requestId"`
	IsSucceeded bool  `json:"isSucceeded"`
	Weather
}
