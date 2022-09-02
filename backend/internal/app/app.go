package app

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/arthurshafikov/appcreative/backend/internal/cache"
	"github.com/arthurshafikov/appcreative/backend/internal/clients"
	"github.com/arthurshafikov/appcreative/backend/internal/config"
	"github.com/arthurshafikov/appcreative/backend/internal/logger"
	"github.com/arthurshafikov/appcreative/backend/internal/services"
	"github.com/arthurshafikov/appcreative/backend/internal/transport/http"
	"github.com/arthurshafikov/appcreative/backend/internal/transport/http/handler"
	"golang.org/x/sync/errgroup"
)

var envFileLocation string

func init() {
	flag.StringVar(&envFileLocation, "env", "./.env", "Path to .env file")
}

func Run() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)
	config := config.NewConfig(envFileLocation)
	logger := logger.NewLogger()
	openWeatherMapClient := clients.NewOpenWeatherMap(config.OpenWeatherMapConfig.APIKey)
	redis := cache.NewRedis(cache.RedisConfig{
		Address:      config.RedisConfig.Address,
		TTLInSeconds: config.RedisConfig.TTLInSeconds,
	})
	services := services.NewServices(&services.Dependencies{
		Logger:        logger,
		WeatherClient: openWeatherMapClient,
		Cache:         redis,
	})

	handler := handler.NewHandler(ctx, services)
	http.NewServer(handler, logger).Serve(ctx, group, config.ServerConfig.Port)

	if err := group.Wait(); err != nil {
		logger.Error(err)
	}
}
