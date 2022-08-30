package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig
}

type ServerConfig struct {
	Port string
}

func NewConfig(envFileLocation string) *Config {
	if err := godotenv.Load(envFileLocation); err != nil {
		log.Fatalln(err)
	}

	return &Config{
		ServerConfig: ServerConfig{
			Port: os.Getenv("APP_PORT"),
		},
	}
}
