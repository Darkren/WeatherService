package server

import (
	"database/sql"
	"fmt"

	"github.com/Darkren/weatherservice/config"
)

type Server struct {
	config config.Config
	db     *sql.DB
}

func New(config config.Config) *Server {
	return &Server{config: config}
}

func (s *Server) Start() {
	dbConfigSection, err := s.config.Section("db")
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

	s.db = db

}
