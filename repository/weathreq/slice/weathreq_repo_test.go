package slice

import (
	"testing"

	"github.com/Darkren/weatherservice/models"
)

func TestNew(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("Got err on creating repo")
	}
}

func TestAdd(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []struct {
		req    *models.WeatherRequest
		wantID int64
	}{
		{
			req: &models.WeatherRequest{
				Lat: 15.46,
				Lon: 13.45,
			},
			wantID: 1,
		},
		{
			req: &models.WeatherRequest{
				Lat: 19.28,
				Lon: 10.287,
			},
			wantID: 2,
		},
		{
			req: &models.WeatherRequest{
				Lat: 12.22222,
				Lon: 90.87,
			},
			wantID: 3,
		},
	}

	for _, test := range tests {
		id, err := repo.Add(test.req)
		if err != nil {
			t.Errorf("Got err inserting %v", test.req)
		}

		if id != test.wantID {
			t.Errorf("Got %v, expected %v", id, test.wantID)
		}
	}
}

func TestGetNotComplete(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []struct {
		req    *models.WeatherRequest
		wantID int64
	}{
		{
			req: &models.WeatherRequest{
				Lat: 15.46,
				Lon: 13.45,
			},
			wantID: 1,
		},
		{
			req: &models.WeatherRequest{
				Lat: 19.28,
				Lon: 10.287,
			},
			wantID: 2,
		},
		{
			req: &models.WeatherRequest{
				Lat: 12.22222,
				Lon: 90.87,
			},
			wantID: 3,
		},
	}

	for _, test := range tests {
		id, err := repo.Add(test.req)
		if err != nil {
			t.Errorf("Got err inserting %v", test.req)
		}

		if id != test.wantID {
			t.Errorf("Got %v, expected %v", id, test.wantID)
		}
	}

	for _, test := range tests {
		test.req.IsComplete = true
	}

	notComplete, err := repo.GetNotComplete()
	if err != nil {
		t.Error("Got err on GetNotComplete")
	}

	if notComplete != nil {
		t.Errorf("Got %v, expected none", notComplete)
	}

	tests[1].req.IsComplete = false

	notComplete, err = repo.GetNotComplete()
	if err != nil {
		t.Error("Got err on GetNotComplete")
	}

	if notComplete == nil {
		t.Error("Got none")
	}
	if notComplete.IsComplete {
		t.Errorf("Got %v, IsComplete should be false", notComplete)
	}
}

func TestSetComplete(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []struct {
		req    *models.WeatherRequest
		wantID int64
	}{
		{
			req: &models.WeatherRequest{
				Lat: 15.46,
				Lon: 13.45,
			},
			wantID: 1,
		},
		{
			req: &models.WeatherRequest{
				Lat: 19.28,
				Lon: 10.287,
			},
			wantID: 2,
		},
		{
			req: &models.WeatherRequest{
				Lat: 12.22222,
				Lon: 90.87,
			},
			wantID: 3,
		},
	}

	for _, test := range tests {
		id, err := repo.Add(test.req)
		if err != nil {
			t.Errorf("Got err inserting %v", test.req)
		}

		if id != test.wantID {
			t.Errorf("Got %v, expected %v", id, test.wantID)
		}
	}

	err := repo.SetComplete(3)
	if err != nil {
		t.Error("Got err in SetComplete")
	}

	if !tests[2].req.IsComplete {
		t.Error("IsComplete should be true")
	}
}
