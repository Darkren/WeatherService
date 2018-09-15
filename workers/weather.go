package workers

import (
	"log"

	"github.com/Darkren/weatherservice/models"

	"github.com/Darkren/weatherservice/services"

	"github.com/Darkren/weatherservice/repository"
)

// WeatherWorker fetches next not complete request from the repository.
// Then it gets weather info from the external service and stores
// response to the repository. In case external service returned an
// error - the response is written with default values and
// IsSucceeded equaling false. Not fully thread-safe.
// Concurrent workers will end up processing same request
type WeatherWorker struct {
	requestRepository  repository.WeatherRequestRepository
	responseRepository repository.WeatherResponseRepository
	weatherService     services.Weather
}

// New constructs and returns worker
func New(reqRepo repository.WeatherRequestRepository,
	respRepo repository.WeatherResponseRepository) *WeatherWorker {
	return &WeatherWorker{
		requestRepository:  reqRepo,
		responseRepository: respRepo,
	}
}

// Run starts the worker
func (w *WeatherWorker) Run() {
	for {
		// get next request
		next, err := w.requestRepository.GetNotComplete()
		if err != nil {
			log.Println("Got err fetching next request")

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

		w.requestRepository.SetComplete(next.ID)
	}
}
