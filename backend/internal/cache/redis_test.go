package cache

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/arthurshafikov/appcreative/backend/internal/core"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/require"
)

var (
	key                   = "someKey"
	cachedWeatherResponse = core.WeatherResponse{
		City:        "SomeCity",
		Temperature: 20,
	}
)

func getRedisWithClient(t *testing.T) (*Redis, *redis.Client) {
	t.Helper()
	mr, err := miniredis.Run()
	require.NoError(t, err)
	redisClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return &Redis{
		client: redisClient,
	}, redisClient
}

func TestGetAndUnmarshal(t *testing.T) {
	redis, redisClient := getRedisWithClient(t)
	cachedWeatherResponseJSON, err := json.Marshal(cachedWeatherResponse)
	require.NoError(t, err)
	err = redisClient.Set(key, cachedWeatherResponseJSON, time.Hour).Err()
	require.NoError(t, err)

	var unmarshaledResult core.WeatherResponse
	err = redis.GetAndUnmarshal(key, &unmarshaledResult)
	require.NoError(t, err)

	require.Equal(t, cachedWeatherResponse.City, unmarshaledResult.City)
	require.Equal(t, cachedWeatherResponse.Temperature, unmarshaledResult.Temperature)
}

func TestMarshalAndSet(t *testing.T) {
	redis, redisClient := getRedisWithClient(t)
	cachedWeatherResponseJSON, err := json.Marshal(cachedWeatherResponse)
	require.NoError(t, err)

	err = redis.MarshalAndSet(key, cachedWeatherResponse)
	require.NoError(t, err)

	result, err := redisClient.Get(key).Result()
	require.NoError(t, err)
	require.Equal(t, cachedWeatherResponseJSON, []byte(result))
}
