package services

import "github.com/arthurshafikov/appcreative/backend/internal/core"

type Weather interface {
	GetCurrentWeather(city string) (*core.WeatherResponse, error)
}

type Services struct {
	Weather
}

type Logger interface {
	Error(err error)
}

type WeatherClient interface {
	GetCurrentWeather(city string) (*core.WeatherResponse, error)
}

type Dependencies struct {
	Logger
	WeatherClient
}

func NewServices(deps *Dependencies) *Services {
	return &Services{
		Weather: NewWeatherService(deps.Logger, deps.WeatherClient),
	}
}
