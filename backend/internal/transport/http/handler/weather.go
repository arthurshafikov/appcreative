package handler

import (
	"errors"
	"net/http"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initWeatherRoutes(e *gin.Engine) {
	v1 := e.Group("v1")
	{
		v1.GET("getCurrentWeather", h.getCurrentWeather)
	}
}

func (h *Handler) getCurrentWeather(ctx *gin.Context) {
	city := ctx.Query("city")
	if city == "" {
		h.setErrorJSONResponse(ctx, http.StatusUnprocessableEntity, core.ErrorBag{
			"city": []string{core.ErrCityNotDefined.Error()},
		})
		return
	}

	response, err := h.services.Weather.GetCurrentWeather(city)
	if err != nil {
		if errors.Is(err, core.ErrCityNotFound) {
			h.setErrorJSONResponse(ctx, http.StatusUnprocessableEntity, core.ErrorBag{
				"city": []string{core.ErrCityNotFound.Error()},
			})
			return
		}

		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, response)
}
