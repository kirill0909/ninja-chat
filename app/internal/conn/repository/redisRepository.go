package repository

import (
	"context"
	"fmt"
	"ninja-chat-core-api/internal/conn"

	models "ninja-chat-core-api/internal/models/conn"

	"github.com/redis/go-redis/v9"
)

type ConnRedisRepo struct {
	db *redis.Client
}

func NewConnRedisRepo(db *redis.Client) conn.RedisRepo {
	return &ConnRedisRepo{db: db}
}

func (r *ConnRedisRepo) SendMessage(ctx context.Context, request models.SendMessageRequest) (err error) {

	// TODO: give user abillity to set expired time for message
	key := fmt.Sprintf("%d_%s", request.RecipientID, request.MessageUUID)
	if err = r.db.Set(ctx, key, request.Message, 0).Err(); err != nil {
		return
	}

	return
}
