package repository

import "github.com/Darkren/weatherservice/models"

type WeatherResponseRepository interface {
	Add(resp *models.WeatherResponse) (int64, error)
	GetById(id int64) (*models.WeatherResponse, error)
}
