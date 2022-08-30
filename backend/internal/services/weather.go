package services

import "github.com/arthurshafikov/appcreative/backend/internal/core"

type WeatherService struct{}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetCurrentWeather(city string) *core.WeatherResponse {
	return &core.WeatherResponse{}
}
