package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"ninja-chat-core-api/internal/conn"
	models "ninja-chat-core-api/internal/models/conn"
)

type ConnPGRepo struct {
	db *sqlx.DB
}

func NewConnPGRepo(db *sqlx.DB) conn.PGRepo {
	return &ConnPGRepo{db: db}
}

func (r *ConnPGRepo) SendMessage(ctx context.Context, sendMessageRequest models.SendMessageRequest) (
	result models.SendMessageResponse, err error) {
	return models.SendMessageResponse{Success: true, Code: 200}, nil
}
