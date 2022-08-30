package app

import (
	"context"
	"flag"

	"github.com/arthurshafikov/appcreative/backend/internal/config"
	"github.com/arthurshafikov/appcreative/backend/internal/services"
	"github.com/arthurshafikov/appcreative/backend/internal/transport/http"
	"github.com/arthurshafikov/appcreative/backend/internal/transport/http/handler"
)

var envFileLocation string

func init() {
	flag.StringVar(&envFileLocation, "env", "./.env", "Path to .env file")
}

func Run() {
	flag.Parse()

	ctx := context.Background()
	config := config.NewConfig(envFileLocation)
	services := services.NewServices()

	handler := handler.NewHandler(ctx, services)
	http.NewServer(handler).Serve(ctx, config.ServerConfig.Port)
}
