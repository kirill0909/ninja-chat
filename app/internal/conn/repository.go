package conn

import (
	"context"
	models "ninja-chat-core-api/internal/models/conn"
)

type PGRepo interface {
	SendMessage(ctx context.Context, request models.SendMessageRequest) (result models.SendMessageResponse, err error)
}

type RedisRepo interface {
	SendMessage(ctx context.Context, request models.SendMessageRequest) (err error)
}
