package models

type Weather struct {
	Temperature int `json:"temperature"`
	Humidity    int `json:"humidity"`
	Pressure    int `json:"pressure"`
}
