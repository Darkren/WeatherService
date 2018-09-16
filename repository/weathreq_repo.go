package repository

import "github.com/Darkren/weatherservice/models"

// WeatherRequestRepository is a storage of WeatherRequest
type WeatherRequestRepository interface {
	Add(req *models.WeatherRequest) (int64, error)
	GetForProcessing() (*models.WeatherRequest, error)
	ProcessingFinished(id int64) error
}
