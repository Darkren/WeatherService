// Package slice contains WeatherRequestRepository implementation
// with slice as a storage
package slice

import (
	"sync"
	"time"

	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/repository"
)

// WeatherRequestRepository is WeatherRequestRepository interface implementation
// containing data in slice
type WeatherRequestRepository struct {
	storage []*models.WeatherRequest
	serial  int64
	sync.Mutex
}

// New constructs and returns pointer to slice repository
func New() repository.WeatherRequestRepository {
	return &WeatherRequestRepository{}
}

// Add stores request to repository
func (r *WeatherRequestRepository) Add(req *models.WeatherRequest) (int64, error) {
	r.Lock()
	defer r.Unlock()

	r.serial++
	req.ID = r.serial
	req.Created = time.Now()
	req.IsComplete = false

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

// SetComplete marks the request as complete
func (r *WeatherRequestRepository) SetComplete(id int64) error {
	r.Lock()
	defer r.Unlock()

	for _, req := range r.storage {
		if req.ID == id {
			req.IsComplete = true

			break
		}
	}

	return nil
}
