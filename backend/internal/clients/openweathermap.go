package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
)

const appEndpoint = "https://api.openweathermap.org/data/2.5"

type openWeatherMapResponse struct {
	City     string        `json:"name"`
	Weather  []weatherInfo `json:"weather"`
	MainInfo mainInfo      `json:"main"`
	Wind     windInfo      `json:"wind"`
}

type weatherInfo struct {
	Main string `json:"main"`
}

type mainInfo struct {
	Temperature float64 `json:"temp"`
	Pressure    int     `json:"pressure"`
	Humidity    uint8   `json:"humidity"`
}

type windInfo struct {
	Speed   float64 `json:"speed"`
	Degrees int     `json:"deg"`
}

type OpenWeatherMap struct {
	apiKey string
}

func NewOpenWeatherMap(apiKey string) *OpenWeatherMap {
	return &OpenWeatherMap{
		apiKey: apiKey,
	}
}

func (owm *OpenWeatherMap) GetCurrentWeather(city string) (*core.WeatherResponse, error) {
	res, err := http.Get(fmt.Sprintf("%s/weather?appid=%s&q=%s&units=metric", appEndpoint, owm.apiKey, city))
	if err != nil {
		return nil, err
	}

	var responseJSON openWeatherMapResponse
	if err := json.NewDecoder(res.Body).Decode(&responseJSON); err != nil {
		return nil, err
	}

	if len(responseJSON.Weather) < 1 {
		return nil, core.ErrCityNotFound
	}

	return &core.WeatherResponse{
		City:        city,
		Temperature: responseJSON.MainInfo.Temperature,
		WeatherCondition: core.WeatherCondition{
			Type:     responseJSON.Weather[0].Main,
			Pressure: responseJSON.MainInfo.Pressure,
			Humidity: responseJSON.MainInfo.Humidity,
		},
		Wind: core.WindStatus{
			Speed:     responseJSON.Wind.Speed,
			Direction: owm.getWindString(responseJSON.Wind.Degrees),
		},
	}, nil
}

func (owm *OpenWeatherMap) getWindString(degrees int) string {
	value := int((float64(degrees) / 22.5) + 0.5)
	directions := []string{
		"N",
		"NNE",
		"NE",
		"ENE",
		"E",
		"ESE",
		"SE",
		"SSE",
		"S",
		"SSW",
		"SW",
		"WSW",
		"W",
		"WNW",
		"NW",
		"NNW",
	}

	return directions[value%16]
}
