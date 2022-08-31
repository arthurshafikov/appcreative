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

type Cache interface {
	GetAndUnmarshal(key string, pointer any) error
	MarshalAndSet(key string, value any) error
}

type Dependencies struct {
	Logger
	WeatherClient
	Cache
}

func NewServices(deps *Dependencies) *Services {
	return &Services{
		Weather: NewWeatherService(deps.Logger, deps.WeatherClient, deps.Cache),
	}
}
