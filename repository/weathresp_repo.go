package repository

import "github.com/Darkren/weatherservice/models"

// WeatherResponseRepository is a storage for WeatherResponse
type WeatherResponseRepository interface {
	Add(resp *models.WeatherResponse) (int64, error)
	GetByRequestID(requestID int64) (*models.WeatherResponse, error)
}
