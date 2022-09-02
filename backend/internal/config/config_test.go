package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	envFilePath    = ".env.testing"
	envFileContent = []byte(
		"APP_PORT=9999\nOPEN_WEATHER_MAP_API_KEY=SomeKey\nREDIS_ADDRESS=redis:6379\nREDIS_CACHE_TTL_IN_SECONDS=60",
	)
)

func TestNewConfig(t *testing.T) {
	createFakeEnvFile(t)
	defer deleteFakeEnvFile(t)

	config := NewConfig(envFilePath)

	require.Equal(t, config.OpenWeatherMapConfig.APIKey, "SomeKey")
	require.Equal(t, config.ServerConfig.Port, "9999")
	require.Equal(t, config.RedisConfig.Address, "redis:6379")
	require.Equal(t, config.RedisConfig.TTLInSeconds, 60)

}

func createFakeEnvFile(t *testing.T) {
	t.Helper()
	if err := os.WriteFile(envFilePath, envFileContent, 0600); err != nil { //nolint:gofumpt
		t.Fatal(err)
	}
}

func deleteFakeEnvFile(t *testing.T) {
	t.Helper()
	if err := os.Remove(envFilePath); err != nil {
		t.Fatal(err)
	}
}
