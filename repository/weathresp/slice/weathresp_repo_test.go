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
		resp   *models.WeatherResponse
		wantID int64
	}{
		{
			resp: &models.WeatherResponse{
				RequestID: 1,
				Weather: models.Weather{
					Humidity:    45,
					Pressure:    34,
					Temperature: 27,
				},
			},
			wantID: 1,
		},
		{
			resp: &models.WeatherResponse{
				RequestID: 2,
				Weather: models.Weather{
					Humidity:    23,
					Pressure:    11,
					Temperature: 22,
				},
			},
			wantID: 2,
		},
		{
			resp: &models.WeatherResponse{
				RequestID: 3,
				Weather: models.Weather{
					Humidity:    87,
					Pressure:    34,
					Temperature: 20,
				},
			},
			wantID: 3,
		},
	}

	for _, test := range tests {
		id, err := repo.Add(test.resp)
		if err != nil {
			t.Errorf("Got err inserting %v", test.resp)
		}

		if id != test.wantID {
			t.Errorf("Got %v, expected %v", id, test.wantID)
		}
	}
}

func TestByRequestID(t *testing.T) {
	repo := New()
	if repo == nil {
		t.Errorf("Got err on creating repo")
	}

	tests := []struct {
		resp   *models.WeatherResponse
		wantID int64
	}{
		{
			resp: &models.WeatherResponse{
				RequestID: 1,
				Weather: models.Weather{
					Humidity:    45,
					Pressure:    34,
					Temperature: 27,
				},
			},
			wantID: 1,
		},
		{
			resp: &models.WeatherResponse{
				RequestID: 2,
				Weather: models.Weather{
					Humidity:    23,
					Pressure:    11,
					Temperature: 22,
				},
			},
			wantID: 2,
		},
		{
			resp: &models.WeatherResponse{
				RequestID: 3,
				Weather: models.Weather{
					Humidity:    87,
					Pressure:    34,
					Temperature: 20,
				},
			},
			wantID: 3,
		},
	}

	for _, test := range tests {
		id, err := repo.Add(test.resp)
		if err != nil {
			t.Errorf("Got err inserting %v", test.resp)
		}

		if id != test.wantID {
			t.Errorf("Got %v, expected %v", id, test.wantID)
		}
	}

	got, err := repo.GetByRequestID(2)
	if err != nil {
		t.Error("Got err in GetByRequestID")
	}

	if got.RequestID != 2 {
		t.Error("Got wrong response")
	}
}
