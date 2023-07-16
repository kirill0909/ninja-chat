package redis

import (
	"context"
	"fmt"
	"log"
	"ninja-chat-core-api/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return redisClient, nil
}
