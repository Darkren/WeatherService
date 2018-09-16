package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/Darkren/weatherservice/config/json"
	"github.com/Darkren/weatherservice/server"
)

func main() {
	config, err := json.Load("config.json")
	if err != nil {
		log.Fatal("Error reading config file")
	}

	server := server.New(config)

	server.Start()
}
