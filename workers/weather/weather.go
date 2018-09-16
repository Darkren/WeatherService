// Package weather contains wheather info resolving worker
package weather

import (
	"log"
	"time"

	"github.com/Darkren/weatherservice/models"

	"github.com/Darkren/weatherservice/services"

	"github.com/Darkren/weatherservice/repository"
)

// Worker fetches next not complete request from the repository.
// Then it gets weather info from the external service and stores
// response to the repository. In case external service returned an
// error - the response is written with default values and
// IsSucceeded equaling false
type Worker struct {
	requestRepository  repository.WeatherRequestRepository
	responseRepository repository.WeatherResponseRepository
	weatherService     services.Weather
	fetchTimeoutMs     int
}

// New constructs and returns worker
func New(reqRepo repository.WeatherRequestRepository,
	respRepo repository.WeatherResponseRepository,
	weatherService services.Weather,
	fetchTimeoutMs int) *Worker {
	return &Worker{
		requestRepository:  reqRepo,
		responseRepository: respRepo,
		weatherService:     weatherService,
		fetchTimeoutMs:     fetchTimeoutMs,
	}
}

// Run starts the worker
func (w *Worker) Run() {
	for {
		// get next request
		next, err := w.requestRepository.GetForProcessing()
		if err != nil {
			log.Printf("Got err fetching next request: %v", err)

			continue
		}

		if next == nil {
			time.Sleep(time.Duration(w.fetchTimeoutMs) * time.Millisecond)

			continue
		}

		resp := &models.WeatherResponse{
			RequestID: next.ID,
		}

		// get service response
		if serviceResp, err := w.weatherService.Get(next.Lat, next.Lon); err != nil {
			log.Println("Got err requesting weather from service")

			w.responseRepository.Add(resp)
		} else {
			resp.IsSucceeded = true
			resp.Pressure = serviceResp.Pressure
			resp.Humidity = serviceResp.Humidity
			resp.Temperature = serviceResp.Temperature

			w.responseRepository.Add(resp)
		}

		w.requestRepository.ProcessingFinished(next.ID)
	}
}
