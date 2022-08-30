package handler

import (
	"context"

	"github.com/arthurshafikov/appcreative/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctx      context.Context
	services *services.Services
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewHandler(
	ctx context.Context,
	services *services.Services,
) *Handler {
	return &Handler{
		ctx:      ctx,
		services: services,
	}
}

func (h *Handler) InitRoutes(e *gin.Engine) {
}

func (h *Handler) setErrorJSONResponse(ctx *gin.Context, code int, err string) {
	ctx.JSON(code, ErrorResponse{
		Error: err,
	})
}
