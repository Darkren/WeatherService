package mock

import "github.com/Darkren/weatherservice/repository"

type MockWeatherRequestRepository struct {
}

func New() repository.WeatherRequestRepository {
	return &MockWeatherRequestRepository{}
}
