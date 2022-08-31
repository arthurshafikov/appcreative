package services

import (
	"github.com/arthurshafikov/appcreative/backend/internal/core"
)

type WeatherService struct {
	logger Logger
	client WeatherClient
}

func NewWeatherService(logger Logger, client WeatherClient) *WeatherService {
	return &WeatherService{
		logger: logger,
		client: client,
	}
}

func (s *WeatherService) GetCurrentWeather(city string) (*core.WeatherResponse, error) {
	weatherResponse, err := s.client.GetCurrentWeather(city)
	if err != nil {
		s.logger.Error(err)
		return nil, core.ErrInternalServerError
	}

	return weatherResponse, nil
}
