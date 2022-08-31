package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenWeatherMapConfig
	ServerConfig
	RedisConfig
}

type OpenWeatherMapConfig struct {
	APIKey string
}

type ServerConfig struct {
	Port string
}

type RedisConfig struct {
	Address      string
	TTLInSeconds int
}

func NewConfig(envFileLocation string) *Config {
	if err := godotenv.Load(envFileLocation); err != nil {
		log.Fatalln(err)
	}

	redisTTLInSeconds, err := strconv.Atoi(os.Getenv("REDIS_CACHE_TTL_IN_SECONDS"))
	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		OpenWeatherMapConfig: OpenWeatherMapConfig{
			APIKey: os.Getenv("OPEN_WEATHER_MAP_API_KEY"),
		},
		ServerConfig: ServerConfig{
			Port: os.Getenv("APP_PORT"),
		},
		RedisConfig: RedisConfig{
			Address:      os.Getenv("REDIS_ADDRESS"),
			TTLInSeconds: redisTTLInSeconds,
		},
	}
}
