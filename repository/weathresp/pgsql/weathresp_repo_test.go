package pgsql

import (
	"testing"

	_ "github.com/lib/pq"

	"github.com/Darkren/weatherservice/models"
)

func TestNew(t *testing.T) {
	host, port, user, password, dbName := "127.0.0.1", 5432, "postgres", "1234", "weatherservice"

	_, err := New(user, password, dbName, host, port)
	if err != nil {
		t.Errorf("Got err on creating repo")
	}

}

func TestAdd(t *testing.T) {
	host, port, user, password, dbName := "127.0.0.1", 5432, "postgres", "1234", "weatherservice"

	repo, err := New(user, password, dbName, host, port)
	if err != nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []*models.WeatherResponse{
		{
			RequestID: 1,
			Weather: models.Weather{
				Humidity:    45,
				Pressure:    34,
				Temperature: 27,
			},
		},
		{
			RequestID: 2,
			Weather: models.Weather{
				Humidity:    23,
				Pressure:    11,
				Temperature: 22,
			},
		},
		{
			RequestID: 3,
			Weather: models.Weather{
				Humidity:    87,
				Pressure:    34,
				Temperature: 20,
			},
		},
	}

	for _, test := range tests {
		id, err := repo.Add(test)
		if err != nil {
			t.Errorf("Got err inserting %v", test)
		}

		if id == 0 {
			t.Error("Got 0 as ID")
		}
	}
}

func TestGetByRequestID(t *testing.T) {
	host, port, user, password, dbName := "127.0.0.1", 5432, "postgres", "1234", "weatherservice"

	repo, err := New(user, password, dbName, host, port)
	if err != nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []*models.WeatherResponse{
		{
			RequestID: 1,
			Weather: models.Weather{
				Humidity:    45,
				Pressure:    34,
				Temperature: 27,
			},
		},
		{
			RequestID: 2,
			Weather: models.Weather{
				Humidity:    23,
				Pressure:    11,
				Temperature: 22,
			},
		},
		{
			RequestID: 3,
			Weather: models.Weather{
				Humidity:    87,
				Pressure:    34,
				Temperature: 20,
			},
		},
	}

	for _, test := range tests {
		id, err := repo.Add(test)
		if err != nil {
			t.Errorf("Got err inserting %v", test)
		}

		if id == 0 {
			t.Error("Got 0 as ID")
		}

		test.ID = id
	}

	_, err = repo.GetByRequestID(tests[1].RequestID)
	if err != nil {
		t.Error("Got err in GetByRequestID")
	}
}
