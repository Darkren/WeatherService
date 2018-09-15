// Package server contains Sever struct and all needed methods
// to start the application
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Darkren/weatherservice/services"

	"github.com/gorilla/mux"

	"github.com/Darkren/weatherservice/config"
	"github.com/Darkren/weatherservice/controllers"
	weatherReqRepo "github.com/Darkren/weatherservice/repository/weathreq/slice"
	weatherRespRepo "github.com/Darkren/weatherservice/repository/weathresp/slice"
	weatherService "github.com/Darkren/weatherservice/services/weatherbit"
	weatherWorker "github.com/Darkren/weatherservice/workers/weather"
)

// Server is a server itself. Needs configuration to run
type Server struct {
	config            config.Config
	weatherController *controllers.WeatherController
}

// New constructs new server
func New(config config.Config) *Server {
	return &Server{config: config}
}

// Start does all the needed setup and runs the server
func (s *Server) Start() {
	/*dbConfigSection, err := s.config.Section("db")
	if err != nil {
		panic(err)
	}

	dbLogin := dbConfigSection.MustGetString("login")
	dbPassword := dbConfigSection.MustGetString("password")
	dbName := dbConfigSection.MustGetString("name")

	dbConnStr := fmt.Sprintf("user=%s password=%s dbname=%s", dbLogin,
		dbPassword, dbName)

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		panic(err)
	}

	s.db = db*/

	// get weather service config
	weatherServicesConfig, err := s.config.Section("weatherServices")
	if err != nil {
		log.Fatalf("Got err reading weather services config section: %v", err)
	}

	weatherServiceName := s.config.MustGetString("weatherService")

	// create weather resolving service
	weatherService := instantiateWeatherService(weatherServicesConfig, weatherServiceName)

	// create repositories
	weatherRequestRepo := weatherReqRepo.New()
	weatherResponseRepo := weatherRespRepo.New()

	// create weather worker
	weatherWorker := weatherWorker.New(weatherRequestRepo, weatherResponseRepo,
		weatherService, s.config.MustGetInt("workerFetchTimeoutMs"))

	// start worker
	go weatherWorker.Run()

	// create controllers
	s.weatherController = &controllers.WeatherController{
		WeatherRequestRepository:  weatherRequestRepo,
		WeatherResponseRepository: weatherResponseRepo,
	}

	// setup routing
	router := mux.NewRouter()

	sub := router.PathPrefix("/weather").Subrouter()

	sub.HandleFunc("/search", s.weatherController.Search)
	sub.HandleFunc("/search/", s.weatherController.Search)

	sub.HandleFunc("/result/{reqId}", s.weatherController.Result)
	sub.HandleFunc("/result/{reqId}/", s.weatherController.Result)

	port := s.config.MustGetInt("port")

	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func instantiateWeatherService(weatherServicesConfig config.Config,
	serviceName string) services.Weather {

	weatherServiceConfig, err := weatherServicesConfig.Section(serviceName)
	if err != nil {
		log.Fatalf("Weather service %v not configured", serviceName)
	}

	return weatherService.New(weatherServiceConfig.MustGetString("key"),
		weatherServiceConfig.MustGetString("baseURL"))
}
