package services

import (
	"testing"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
	mock_services "github.com/arthurshafikov/appcreative/backend/internal/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var city = "Oslo"

func TestGetCurrentWeatherFoundInCache(t *testing.T) {
	weatherService, _, _, cacheMock := getWeatherServiceWithMocks(t)
	gomock.InOrder(
		cacheMock.EXPECT().GetAndUnmarshal(city, gomock.Any()).Return(nil),
	)

	_, err := weatherService.GetCurrentWeather(city)
	require.NoError(t, err)
}

func TestGetCurrentWeatherCacheReturnedError(t *testing.T) {
	weatherService, loggerMock, _, cacheMock := getWeatherServiceWithMocks(t)
	gomock.InOrder(
		cacheMock.EXPECT().GetAndUnmarshal(city, gomock.Any()).Return(core.ErrInternalServerError),
		loggerMock.EXPECT().Error(core.ErrInternalServerError),
	)

	_, err := weatherService.GetCurrentWeather(city)
	require.ErrorIs(t, err, core.ErrInternalServerError)
}

func TestGetCurrentWeatherNotInCache(t *testing.T) {
	weatherResponse := &core.WeatherResponse{
		City: city,
	}
	weatherService, _, clientMock, cacheMock := getWeatherServiceWithMocks(t)
	gomock.InOrder(
		cacheMock.EXPECT().GetAndUnmarshal(city, gomock.Any()).Return(core.ErrNotFound),
		clientMock.EXPECT().GetCurrentWeather(city).Return(weatherResponse, nil),
		cacheMock.EXPECT().MarshalAndSet(city, *weatherResponse).Return(nil),
	)

	weatherResponse, err := weatherService.GetCurrentWeather(city)
	require.NoError(t, err)
	require.Equal(t, city, weatherResponse.City)
}

func TestGetCurrentWeatherNotInCacheCityNotFound(t *testing.T) {
	weatherService, _, clientMock, cacheMock := getWeatherServiceWithMocks(t)
	gomock.InOrder(
		cacheMock.EXPECT().GetAndUnmarshal(city, gomock.Any()).Return(core.ErrNotFound),
		clientMock.EXPECT().GetCurrentWeather(city).Return(nil, core.ErrCityNotFound),
	)

	weatherResponse, err := weatherService.GetCurrentWeather(city)
	require.Nil(t, weatherResponse)
	require.ErrorIs(t, err, core.ErrCityNotFound)
}

func TestGetCurrentWeatherNotInCacheClientServerError(t *testing.T) {
	weatherService, loggerMock, clientMock, cacheMock := getWeatherServiceWithMocks(t)
	gomock.InOrder(
		cacheMock.EXPECT().GetAndUnmarshal(city, gomock.Any()).Return(core.ErrNotFound),
		clientMock.EXPECT().GetCurrentWeather(city).Return(nil, core.ErrInternalServerError),
		loggerMock.EXPECT().Error(core.ErrInternalServerError),
	)

	weatherResponse, err := weatherService.GetCurrentWeather(city)
	require.Nil(t, weatherResponse)
	require.ErrorIs(t, err, core.ErrInternalServerError)
}

func TestGetCurrentWeatherNotInCacheSetCacheServerError(t *testing.T) {
	weatherService, loggerMock, clientMock, cacheMock := getWeatherServiceWithMocks(t)
	weatherResponse := &core.WeatherResponse{
		City: city,
	}
	gomock.InOrder(
		cacheMock.EXPECT().GetAndUnmarshal(city, gomock.Any()).Return(core.ErrNotFound),
		clientMock.EXPECT().GetCurrentWeather(city).Return(weatherResponse, nil),
		cacheMock.EXPECT().MarshalAndSet(city, *weatherResponse).Return(core.ErrInternalServerError),
		loggerMock.EXPECT().Error(core.ErrInternalServerError),
	)

	weatherResponse, err := weatherService.GetCurrentWeather(city)
	require.Nil(t, weatherResponse)
	require.ErrorIs(t, err, core.ErrInternalServerError)
}

func getWeatherServiceWithMocks(
	t *testing.T,
) (*WeatherService, *mock_services.MockLogger, *mock_services.MockWeatherClient, *mock_services.MockCache) {
	t.Helper()

	ctrl := gomock.NewController(t)
	loggerMock := mock_services.NewMockLogger(ctrl)
	clientMock := mock_services.NewMockWeatherClient(ctrl)
	cacheMock := mock_services.NewMockCache(ctrl)
	weatherService := NewWeatherService(loggerMock, clientMock, cacheMock)

	return weatherService, loggerMock, clientMock, cacheMock
}
