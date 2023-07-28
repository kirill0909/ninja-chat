package conn

import (
	"context"
	models "ninja-chat-core-api/internal/models/conn"
)

type Usecase interface {
	SendMessage(ctx context.Context, sendMessageRequest models.SendMessageRequest) (result models.SendMessageResponse, err error)
}
