package pgsql

/*import (
	"testing"

	_ "github.com/lib/pq"

	"github.com/Darkren/weatherservice/models"
)

func TestNew(t *testing.T) {
	host, port, user, password, dbName := "127.0.0.1", 5432, "postgres", "1234", "weatherservice"

	_, err := New(user, password, dbName, host, port)
	if err != nil {
		t.Errorf("Got err connecting to DB: %v", err)
	}
}

func TestAdd(t *testing.T) {
	host, port, user, password, dbName := "127.0.0.1", 5432, "postgres", "1234", "weatherservice"

	repo, err := New(user, password, dbName, host, port)
	if err != nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []*models.WeatherRequest{
		{
			Lat: 15.46,
			Lon: 13.45,
		},
		{
			Lat: 19.28,
			Lon: 10.287,
		},
		{
			Lat: 12.22222,
			Lon: 90.87,
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

func TestGetForProcessing(t *testing.T) {
	host, port, user, password, dbName := "127.0.0.1", 5432, "postgres", "1234", "weatherservice"

	repo, err := New(user, password, dbName, host, port)
	if err != nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []*models.WeatherRequest{
		{
			Lat: 15.46,
			Lon: 13.45,
		},
		{
			Lat: 19.28,
			Lon: 10.287,
		},
		{
			Lat: 12.22222,
			Lon: 90.87,
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

	notComplete, err := repo.GetForProcessing()
	if err != nil {
		t.Error("Got err on GetForProcessing")
	}

	if notComplete == nil {
		t.Error("Got none")
	}
	if notComplete.IsComplete {
		t.Errorf("Got %v, IsComplete should be false", notComplete)
	}
	if !notComplete.IsInProgress {
		t.Errorf("Got %v, IsInProgress should be true", notComplete)
	}
}

func TestProcessingFinished(t *testing.T) {
	host, port, user, password, dbName := "127.0.0.1", 5432, "postgres", "1234", "weatherservice"

	repo, err := New(user, password, dbName, host, port)
	if err != nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []*models.WeatherRequest{
		{
			Lat: 15.46,
			Lon: 13.45,
		},
		{
			Lat: 19.28,
			Lon: 10.287,
		},
		{
			Lat: 12.22222,
			Lon: 90.87,
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

	err = repo.ProcessingFinished(tests[2].ID)
	if err != nil {
		t.Error("Got err in ProcessingFinished")
	}
}
*/
