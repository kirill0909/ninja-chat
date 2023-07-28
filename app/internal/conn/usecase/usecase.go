package usecase

import (
	"context"
	"ninja-chat-core-api/config"
	"ninja-chat-core-api/internal/conn"
	models "ninja-chat-core-api/internal/models/conn"
)

type ConnUsecase struct {
	cfg      *config.Config
	connRepo conn.PGRepo
}

func NewConnUsecase(cfg *config.Config, connRepo conn.PGRepo) conn.Usecase {
	return &ConnUsecase{cfg: cfg, connRepo: connRepo}
}

func (u *ConnUsecase) SendMessage(ctx context.Context, sendMessageRequest models.SendMessageRequest) (
	result models.SendMessageResponse, err error) {
	return u.connRepo.SendMessage(ctx, sendMessageRequest)
}
