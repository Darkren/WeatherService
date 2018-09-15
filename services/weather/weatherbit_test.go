package weather

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	tests := []struct {
		lat float64
		lon float64
	}{
		{
			lat: 12.675,
			lon: 17.854,
		},
		{
			lat: 34.678,
			lon: 23.432,
		},
		{
			lat: 30.768,
			lon: 23.454,
		},
	}

	weatherService := New("ea3a97ff41524f7499584aa5403ba4b4",
		"https://api.weatherbit.io/v2.0")

	for _, test := range tests {
		resp, err := weatherService.Get(test.lat, test.lon)
		if err != nil {
			t.Errorf("Got err with weatherbit.io: %v", err)
		}

		if resp == nil {
			t.Errorf("Got no response from weatherbit.io")
		}

		if resp.Humidity == 0 || resp.Pressure == 0 || resp.Temperature == 0 {
			t.Errorf("Got incorrect values from weatherbit.io: %v", *resp)
		}

		// timeout for a bit of safety
		time.Sleep(1 * time.Second)
	}
}
