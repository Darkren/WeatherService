// Package mock contains WeatherResponseRepository implementation
// with mock data
package mock

import (
	"sync"

	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/repository"
)

// WeatherResponseRepository is WeatherResponseRepository interface implementation
// containing mock data
type WeatherResponseRepository struct {
	storage []*models.WeatherResponse
	serial  int64
	sync.Mutex
}

// New constructs and returns pointer to mock repository
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

// GetByID returns response with the specified ID
func (r *WeatherResponseRepository) GetByID(id int64) (*models.WeatherResponse, error) {
	r.Lock()
	defer r.Unlock()

	for _, resp := range r.storage {
		if resp.ID == id {
			return resp, nil
		}
	}

	return nil, nil
}
