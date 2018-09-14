package repository

import "github.com/Darkren/weatherservice/models"

// WeatherRequestRepository is a storage of WeatherRequest
type WeatherRequestRepository interface {
	Add(req *models.WeatherRequest) (int64, error)
	GetNotComplete() (*models.WeatherRequest, error)
}
