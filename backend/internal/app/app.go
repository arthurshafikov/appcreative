package app

import (
	"context"

	"github.com/arthurshafikov/appcreative/backend/internal/services"
	"github.com/arthurshafikov/appcreative/backend/internal/transport/http"
	"github.com/arthurshafikov/appcreative/backend/internal/transport/http/handler"
)

func Run() {
	ctx := context.Background()
	services := services.NewServices()

	handler := handler.NewHandler(ctx, services)
	http.NewServer(handler).Serve(ctx, 8123)
}
