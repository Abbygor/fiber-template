package redis

import (
	"context"
	"fiber-template/internal/config"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Create(context.Context) (*redis.Client, error)
}

type redisClient struct {
	config *config.Config
}

func NewRedisClient(cfg *config.Config) RedisClient {
	return &redisClient{
		config: cfg,
	}
}

func (r *redisClient) Create(ctx context.Context) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr: r.buildRedisConnectionString(),
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}
	log.Println("redis connection success")

	return rdb, nil
}

func (r *redisClient) buildRedisConnectionString() string {
	return fmt.Sprintf("%s:%s",
		r.config.RedisConnection.Server,
		r.config.RedisConnection.Port,
	)
}
