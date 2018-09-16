// Package server contains Sever struct and all needed methods
// to start the application
package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Darkren/weatherservice/graceful"
	"github.com/Darkren/weatherservice/services"

	"github.com/gorilla/mux"

	"github.com/Darkren/weatherservice/config"
	"github.com/Darkren/weatherservice/controllers"
	weatherReqRepo "github.com/Darkren/weatherservice/repository/weathreq/pgsql"
	weatherRespRepo "github.com/Darkren/weatherservice/repository/weathresp/pgsql"
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
	dbConfigSection, err := s.config.Section("db")
	if err != nil {
		log.Fatalf("Error reading DB config: %v", err)
	}

	dbHost := dbConfigSection.MustGetString("host")
	dbPort := dbConfigSection.MustGetInt("port")
	dbUser := dbConfigSection.MustGetString("user")
	dbPassword := dbConfigSection.MustGetString("password")
	dbName := dbConfigSection.MustGetString("dbName")
	sslmode := dbConfigSection.MustGetString("sslmode")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, sslmode)

	// get db connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Couldn't establish connection to DB: %v", err)
	}

	// create repositories
	weatherRequestRepo := weatherReqRepo.New(db)
	weatherResponseRepo := weatherRespRepo.New(db)

	// get weather service config
	weatherServicesConfig, err := s.config.Section("weatherServices")
	if err != nil {
		log.Fatalf("Got err reading weather services config section: %v", err)
	}

	weatherServiceName := s.config.MustGetString("weatherService")

	// create weather resolving service
	weatherService := instantiateWeatherService(weatherServicesConfig, weatherServiceName)

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

	server := http.Server{Addr: fmt.Sprintf(":%d", port), Handler: router}

	shutdown := graceful.Shutdown(&server)

	log.Printf("Server started listening on port: %v", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Couldn't start server on port %v: %v", port, err)
	}

	<-shutdown

	log.Println("Server stopped")
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
