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

	key := fmt.Sprintf("%s_%d_%s", userSessionPrefix, params.UserID, params.AccessToken)
	if err := r.db.Set(ctx, key, params.AccessToken, time.Duration(params.ExpiredAt)).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) GetUserSession(ctx context.Context, req models.AuthHeaders) (result models.UserSession, err error) {

	var key string
	iter := r.db.Scan(ctx, 0, fmt.Sprintf("%s_*_%s", userSessionPrefix, req.AccessToken), 1).Iterator()
	if iter.Next(ctx) {
		key = iter.Val()
	} else {
		return models.UserSession{}, redis.Nil
	}

	userSessionString, err := r.db.Get(ctx, key).Result()
	if err != nil {
		return models.UserSession{}, err
	}

	result.AccessToken = userSessionString

	return
}
