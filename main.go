package main

import (
	"WeatherService/config"
	"fmt"
)

func main() {
	test := `{
		"id": 1,
		"name": "qwerty",
		"birthday": "12.09.2018",
		"address": 
		{
			"city": "Moscow",
			"street": "Lenina str."
		}
	}`

	configuration, err := config.New(test)
	if err != nil {
		panic(err)
	}

	name, err := configuration.GetString("name")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %v\n", name)

	id, err := configuration.GetInt("id")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Id: %v\n", id)

	addressSection, err := configuration.Section("address")
	if err != nil {
		panic(err)
	}

	city, err := addressSection.GetString("city")
	if err != nil {
		panic(err)
	}

	fmt.Printf("City: %v\n", city)
}
