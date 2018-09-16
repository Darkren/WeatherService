package models

// Weather is a weather information
type Weather struct {
	Temperature int `json:"temperature"`
	Humidity    int `json:"humidity"`
	Pressure    int `json:"pressure"`
}
