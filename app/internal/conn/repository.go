package conn

import (
	"context"
	models "ninja-chat-core-api/internal/models/conn"
)

type PGRepo interface {
	SaveMessage(ctx context.Context, request models.SaveMessageRequest) (result models.SaveMessageResponse, err error)
}

type RedisRepo interface {
	SaveMessage(ctx context.Context, request models.SaveMessageRequest) (err error)
}
