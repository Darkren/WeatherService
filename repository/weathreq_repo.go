package repository

import "github.com/Darkren/weatherservice/models"

type WeatherRequestRepository interface {
	Add(req *models.WeatherRequest) (int64, error)
	GetNotComplete() (*models.WeatherRequest, error)
}
