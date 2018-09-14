package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/repository"

	"github.com/gorilla/mux"
)

// WeatherController processes all the requests connected with weather
type WeatherController struct {
	WeatherRequestRepository  repository.WeatherRequestRepository
	WeatherResponseRepository repository.WeatherResponseRepository
}

// Search creates new weather search requests and saves it to storage
func (c *WeatherController) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)

		return
	}

	// anonymous struct to check whether incoming request body
	// contains all needed parameters or not
	decoder := json.NewDecoder(r.Body)
	var req struct {
		Lat *float64 `json:"lat"`
		Lon *float64 `json:"lon"`
	}

	err := decoder.Decode(&req)
	if err != nil || req.Lon == nil || req.Lat == nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// save request to storage for further processing
	reqID, err := c.WeatherRequestRepository.Add(&models.WeatherRequest{
		Lat: *req.Lat,
		Lon: *req.Lon,
	})
	if err != nil {
		log.Printf("%v Error inserting request with values lat: %v lon: %v", time.Now(),
			req.Lat, req.Lon)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	resp := struct {
		RequestID int64 `json:"requestId"`
	}{
		RequestID: reqID,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("%v Error marshalling response: %v", time.Now(), resp)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

// Result returns the search result looking for the specified request id
// If nothing found - returns 404
func (c *WeatherController) Result(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	// get request id paramete from route
	vars := mux.Vars(r)

	reqIDStr, ok := vars["reqId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	reqID, err := strconv.ParseInt(reqIDStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	result, err := c.WeatherResponseRepository.GetByRequestID(reqID)
	if err != nil {
		log.Printf("%v Got err while selecting result with reqID: %v", time.Now(), reqID)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// nothing found, id either invalid or its request
	// is strill being processed
	if result == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	resp := struct {
		RequestID   int64 `json:"requestId"`
		Pressure    int   `json:"pressure"`
		Temperature int   `json:"temperature"`
		Humidity    int   `json:"humidity"`
	}{
		RequestID:   reqID,
		Pressure:    result.Pressure,
		Temperature: result.Temperature,
		Humidity:    result.Humidity,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("%v Error marshalling response: %v", time.Now(), resp)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}
