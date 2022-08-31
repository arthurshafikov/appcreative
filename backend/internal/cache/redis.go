package cache

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/arthurshafikov/appcreative/backend/internal/core"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
	ttl    time.Duration
}

type RedisConfig struct {
	Address      string
	Password     string
	TTLInSeconds int
}

func NewRedis(config RedisConfig) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Fatalln(err)
	}

	return &Redis{
		client: client,
		ttl:    time.Second * time.Duration(config.TTLInSeconds),
	}
}

func (c *Redis) GetAndUnmarshal(key string, pointer any) error {
	result, err := c.client.Get(key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return core.ErrNotFound
		}

		return err
	}

	return json.Unmarshal([]byte(result), pointer)
}

func (c *Redis) MarshalAndSet(key string, value any) error {
	marshalled, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(key, marshalled, c.ttl).Err()
}
