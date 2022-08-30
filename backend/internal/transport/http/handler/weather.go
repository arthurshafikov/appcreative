package handler

import (
	"net/http"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initWeatherRoutes(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		v1.GET("/getCurrentWeather", h.getCurrentWeather)
	}
}

func (h *Handler) getCurrentWeather(ctx *gin.Context) {
	city := ctx.Query("city")
	if city == "" {
		h.setErrorJSONResponse(ctx, http.StatusUnprocessableEntity, core.ErrorBag{
			"city": []string{core.CityNotDefined},
		})
		return
	}

	response := h.services.Weather.GetCurrentWeather(city)

	ctx.JSON(200, response)
}
