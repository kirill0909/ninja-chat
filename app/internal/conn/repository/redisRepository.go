package repository

import (
	"context"
	"ninja-chat-core-api/internal/conn"

	"github.com/redis/go-redis/v9"
	models "ninja-chat-core-api/internal/models/conn"
)

type ConnRedisRepo struct {
	db *redis.Client
}

func NewConnRedisRepo(db *redis.Client) conn.RedisRepo {
	return &ConnRedisRepo{db: db}
}

func (r *ConnRedisRepo) SendMessage(ctx context.Context, request models.SendMessageRequest) (err error) {
	return
}
