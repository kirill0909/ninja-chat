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

func (r *ConnRedisRepo) SaveMessage(ctx context.Context, request models.SaveMessageRequest) (err error) {

	// TODO: give user abillity to set expired time for message
	key := fmt.Sprintf("%d_%d_%d", request.UserID, request.RecipientID, request.MessageID)
	if err = r.db.Set(ctx, key, request.Message, 0).Err(); err != nil {
		return
	}

	return
}
