// Package mock contains WeatherRequestRepository implementation
// with mock data
package mock

import (
	"sync"

	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/repository"
)

// WeatherRequestRepository is WeatherRequestRepository interface implementation
// containing mock data
type WeatherRequestRepository struct {
	storage []*models.WeatherRequest
	serial  int64
	sync.Mutex
}

// New constructs and returns pointer to mock repository
func New() repository.WeatherRequestRepository {
	return &WeatherRequestRepository{}
}

// Add stores request to repository
func (r *WeatherRequestRepository) Add(req *models.WeatherRequest) (int64, error) {
	r.Lock()
	defer r.Unlock()

	r.serial++
	req.ID = r.serial

	r.storage = append(r.storage, req)

	return req.ID, nil
}

// GetNotComplete returns first available not complete request
func (r *WeatherRequestRepository) GetNotComplete() (*models.WeatherRequest, error) {
	r.Lock()
	defer r.Unlock()

	for _, req := range r.storage {
		if !req.IsComplete {
			return req, nil
		}
	}

	return nil, nil
}
