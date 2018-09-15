package services

import (
	"github.com/Darkren/weatherservice/models"
)

// Weather is a weather service interface which allows to
// get weather info from external services by lat and lon
type Weather interface {
	Get(lat, lon float64) (*models.Weather, error)
}
