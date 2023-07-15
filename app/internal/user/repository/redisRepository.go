package repository

import (
	"github.com/go-redis/redis"
	"ninja-chat-core-api/config"
)

type RedisRepo struct {
	cfg *config.Config
	rdb *redis.Client
}

func NewRedisRepo(cfg *config.Config, rdb *redis.Client) *RedisRepo {
	return &RedisRepo{cfg: cfg, rdb: rdb}
}
