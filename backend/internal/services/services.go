package services

import "github.com/arthurshafikov/appcreative/backend/internal/core"

type Weather interface {
	GetCurrentWeather(city string) *core.WeatherResponse
}

type Services struct {
	Weather
}

func NewServices() *Services {
	return &Services{
		Weather: NewWeatherService(),
	}
}
