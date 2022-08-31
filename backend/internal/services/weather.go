package services

import "github.com/arthurshafikov/appcreative/backend/internal/core"

type WeatherService struct {
	client WeatherClient
}

func NewWeatherService(client WeatherClient) *WeatherService {
	return &WeatherService{
		client: client,
	}
}

func (s *WeatherService) GetCurrentWeather(city string) (*core.WeatherResponse, error) {
	return s.client.GetCurrentWeather(city)
}
