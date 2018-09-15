package services

import (
	"github.com/Darkren/weatherservice/models"
)

type Weather interface {
	Get(lat, lon float64) (*models.Weather, error)
}
