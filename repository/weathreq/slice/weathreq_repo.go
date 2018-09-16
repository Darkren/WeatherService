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

// GetForProcessing returns first available not complete request and
// as a side effect marks it as IsInProgress
func (r *WeatherRequestRepository) GetForProcessing() (*models.WeatherRequest, error) {
	r.Lock()
	defer r.Unlock()

	for _, req := range r.storage {
		if !req.IsComplete && !req.IsInProgress {
			req.IsInProgress = true

			return req, nil
		}
	}

	return nil, nil
}

// ProcessingFinished marks the request as complete, also setting IsInProgress to false
// which is not neccessary but actually right
func (r *WeatherRequestRepository) ProcessingFinished(id int64) error {
	r.Lock()
	defer r.Unlock()

	for _, req := range r.storage {
		if req.ID == id {
			req.IsComplete = true
			req.IsInProgress = false

			break
		}
	}

	return nil
}
