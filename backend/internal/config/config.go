package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenWeatherMapConfig
	ServerConfig
}

type OpenWeatherMapConfig struct {
	APIKey string
}

type ServerConfig struct {
	Port string
}

func NewConfig(envFileLocation string) *Config {
	if err := godotenv.Load(envFileLocation); err != nil {
		log.Fatalln(err)
	}

	return &Config{
		OpenWeatherMapConfig: OpenWeatherMapConfig{
			APIKey: os.Getenv("OPEN_WEATHER_MAP_API_KEY"),
		},
		ServerConfig: ServerConfig{
			Port: os.Getenv("APP_PORT"),
		},
	}
}
