package conn

import (
	"context"
	models "ninja-chat-core-api/internal/models/conn"
)

type Usecase interface {
	SaveMessage(ctx context.Context, request models.SaveMessageRequest) (result models.SaveMessageResponse, err error)
}
