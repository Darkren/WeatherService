// Package weatherbit contains weatherbit.io implementation of
// weather service
package weatherbit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/services"
)

const (
	hgMlsInMb = 0.750030101
)

// Weatherbit is an implementation of weather service which
// relise on weatherbit.io API
type Weatherbit struct {
	key     string
	baseURL string
}

// New constructs service using provided API key and base URL
func New(key, baseURL string) services.Weather {
	return &Weatherbit{
		key:     "ea3a97ff41524f7499584aa5403ba4b4",
		baseURL: "https://api.weatherbit.io/v2.0",
	}
}

// Get returns the weather info got from weatherbit.io
func (w *Weatherbit) Get(lat, lon float64) (*models.Weather, error) {
	url := fmt.Sprintf("%s%s?key=%s&lon=%s&lat=%s", w.baseURL, "/current", w.key,
		strconv.FormatFloat(lon, 'f', -1, 64),
		strconv.FormatFloat(lat, 'f', -1, 64))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	weather := struct {
		ObservationsCount int `json:"count"`
		Observations      []struct {
			Temperature float64 `json:"temp"`
			Humidity    float64 `json:"rh"`
			Pressure    float64 `json:"pres"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(respBytes, &weather)
	if err != nil {
		return nil, err
	}

	if weather.ObservationsCount == 0 {
		return nil, fmt.Errorf("No ovservations found")
	}

	// needed conversion ot hg mls cause weatherbit returns pressure in millibars
	weather.Observations[0].Pressure = weather.Observations[0].Pressure * hgMlsInMb

	return &models.Weather{
		Temperature: int(math.Round(weather.Observations[0].Temperature)),
		Humidity:    int(math.Round(weather.Observations[0].Humidity)),
		Pressure:    int(math.Round(weather.Observations[0].Pressure)),
	}, nil
}
