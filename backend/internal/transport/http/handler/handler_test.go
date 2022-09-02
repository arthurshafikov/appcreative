package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestInitRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	writer := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(writer)
	handler := NewHandler(context.Background(), nil)
	expectedRoute := gin.RouteInfo{
		Method: "GET",
		Path:   "/v1/getCurrentWeather",
	}

	handler.InitRoutes(engine)

	firstRoute := engine.Routes()[0]
	require.Equal(t, expectedRoute.Method, firstRoute.Method)
	require.Equal(t, expectedRoute.Path, firstRoute.Path)
}

func TestSetErrorJSONResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	handler := NewHandler(context.Background(), nil)
	errorBag := core.ErrorBag{
		"someField": []string{"some error"},
	}
	expected := core.ErrorResponse{
		Errors: errorBag,
	}
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)

	handler.setErrorJSONResponse(ctx, http.StatusForbidden, errorBag)

	require.Equal(t, http.StatusForbidden, ctx.Writer.Status())
	writerBuffer := writer.Body
	require.Equal(t, expectedJSON, writerBuffer.Bytes())
}
