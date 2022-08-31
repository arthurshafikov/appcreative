package services

import (
	"errors"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
)

type WeatherService struct {
	logger Logger
	client WeatherClient
	cache  Cache
}

func NewWeatherService(logger Logger, client WeatherClient, cache Cache) *WeatherService {
	return &WeatherService{
		logger: logger,
		client: client,
		cache:  cache,
	}
}

func (s *WeatherService) GetCurrentWeather(city string) (*core.WeatherResponse, error) {
	weatherResponse := &core.WeatherResponse{}
	err := s.cache.GetAndUnmarshal(city, weatherResponse)
	if err == nil {
		return weatherResponse, nil
	}
	if !errors.Is(err, core.ErrNotFound) {
		s.logger.Error(err)
		return nil, core.ErrInternalServer
	}

	weatherResponse, err = s.client.GetCurrentWeather(city)
	if err != nil {
		if errors.Is(err, core.ErrCityNotFound) {
			return nil, err
		}

		s.logger.Error(err)
		return nil, core.ErrInternalServer
	}

	if err := s.cache.MarshalAndSet(city, *weatherResponse); err != nil {
		s.logger.Error(err)
		return nil, core.ErrInternalServer
	}

	return weatherResponse, nil
}
