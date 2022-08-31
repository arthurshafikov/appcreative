package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
	"github.com/arthurshafikov/appcreative/backend/internal/services"
	mock_services "github.com/arthurshafikov/appcreative/backend/internal/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetCurrentWeather(t *testing.T) {
	ctrl := gomock.NewController(t)
	weatherServiceMock := mock_services.NewMockWeather(ctrl)
	services := &services.Services{
		Weather: weatherServiceMock,
	}
	writer, ctx, engine := getWriterContextAndHandler(t, services)
	city := "Oslo"
	temperature := float64(20.5)
	gomock.InOrder(
		weatherServiceMock.EXPECT().GetCurrentWeather(city).Times(1).Return(&core.WeatherResponse{
			City:        city,
			Temperature: temperature,
		}, nil),
	)
	ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/getCurrentWeather?city=%s", city), nil)

	engine.ServeHTTP(writer, ctx.Request)

	var jsonResponse core.WeatherResponse
	require.NoError(t, json.Unmarshal(writer.Body.Bytes(), &jsonResponse))
	require.Equal(t, http.StatusOK, writer.Code)
	require.Equal(t, city, jsonResponse.City)
	require.Equal(t, temperature, jsonResponse.Temperature)
}

func TestGetCurrentWeatherMissingCity(t *testing.T) {
	services := &services.Services{}
	writer, ctx, engine := getWriterContextAndHandler(t, services)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/getCurrentWeather", nil)
	expectedErrors := core.ErrorBag{
		"city": []string{core.ErrCityNotDefined.Error()},
	}

	engine.ServeHTTP(writer, ctx.Request)

	var errorResponse core.ErrorResponse
	require.NoError(t, json.Unmarshal(writer.Body.Bytes(), &errorResponse))
	require.Equal(t, http.StatusUnprocessableEntity, writer.Code)
	require.Equal(t, expectedErrors, errorResponse.Errors)
}

func getWriterContextAndHandler(
	t *testing.T,
	services *services.Services,
) (*httptest.ResponseRecorder, *gin.Context, *gin.Engine) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	writer := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(writer)
	handler := NewHandler(context.Background(), services)
	handler.InitRoutes(engine)

	return writer, ctx, engine
}
