package repository

import (
	"fmt"
	"ninja-chat-core-api/config"
	"ninja-chat-core-api/internal/user"
	"time"

	models "ninja-chat-core-api/internal/models/user"

	"context"

	"github.com/redis/go-redis/v9"
)

var (
	userSessionPrefix = "USER_SESSION"
)

type RedisRepo struct {
	cfg *config.Config
	db  *redis.Client
}

func NewRedisRepo(cfg *config.Config, db *redis.Client) user.RedisRepo {
	return &RedisRepo{cfg: cfg, db: db}
}

func (r *RedisRepo) SaveUserSession(ctx context.Context, params models.UserSession) error {

	key := fmt.Sprintf("%s_%d", userSessionPrefix, params.UserID)
	if err := r.db.Set(ctx, key, params.AccessToken, time.Duration(params.ExpireAt)).Err(); err != nil {
		return err
	}
	return nil
}
