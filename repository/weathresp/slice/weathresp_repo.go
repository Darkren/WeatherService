// Package slice contains WeatherResponseRepository implementation
// with slice as a storage
package slice

import (
	"sync"

	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/repository"
)

// WeatherResponseRepository is WeatherResponseRepository interface implementation
// containing data in slice
type WeatherResponseRepository struct {
	storage []*models.WeatherResponse
	serial  int64
	sync.Mutex
}

// New constructs and returns pointer to slice repository
func New() repository.WeatherResponseRepository {
	return &WeatherResponseRepository{}
}

// Add stores response to repository
func (r *WeatherResponseRepository) Add(resp *models.WeatherResponse) (int64, error) {
	r.Lock()
	defer r.Unlock()

	r.serial++
	resp.ID = r.serial

	r.storage = append(r.storage, resp)

	return resp.ID, nil
}

// GetByRequestID returns response with the specified request ID
func (r *WeatherResponseRepository) GetByRequestID(requestID int64) (*models.WeatherResponse, error) {
	r.Lock()
	defer r.Unlock()

	for _, resp := range r.storage {
		if resp.RequestID == requestID {
			return resp, nil
		}
	}

	return nil, nil
}
