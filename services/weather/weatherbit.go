package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/Darkren/weatherservice/config"
	"github.com/Darkren/weatherservice/models"
	"github.com/Darkren/weatherservice/services"
)

type Weatherbit struct {
	key     string
	baseURL string
}

func New(serviceConfig config.Config) services.Weather {
	return &Weatherbit{
		key:     serviceConfig.MustGetString("key"),
		baseURL: serviceConfig.MustGetString("baseURL"),
	}
}

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
		Temperature float64 `json:"temp"`
		Humidity    float64 `json:"rh"`
		Pressure    float64 `json:"pres"`
	}{}
	err = json.Unmarshal(respBytes, &weather)
	if err != nil {
		return nil, err
	}

	return &models.Weather{
		Temperature: int(math.Round(weather.Temperature)),
		Humidity:    int(math.Round(weather.Humidity)),
		Pressure:    int(math.Round(weather.Pressure)),
	}, nil
}
