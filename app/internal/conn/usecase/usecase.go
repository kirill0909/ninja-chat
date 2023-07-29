package usecase

import (
	"context"
	"fmt"
	"ninja-chat-core-api/config"
	"ninja-chat-core-api/internal/conn"
	models "ninja-chat-core-api/internal/models/conn"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
)

type ConnUsecase struct {
	cfg           *config.Config
	connPGRepo    conn.PGRepo
	connRedisRepo conn.RedisRepo
}

func NewConnUsecase(cfg *config.Config, connRepo conn.PGRepo, connRedisRepo conn.RedisRepo) conn.Usecase {
	return &ConnUsecase{cfg: cfg, connPGRepo: connRepo, connRedisRepo: connRedisRepo}
}

func (u *ConnUsecase) SendMessage(ctx context.Context, request models.SendMessageRequest) (
	result models.SendMessageResponse, err error) {

	request.MessageUUID = uuid.New().String()

	result, err = u.connPGRepo.SendMessage(ctx, request)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == violatesForeignKeyCode {
			return models.SendMessageResponse{
				Error: fmt.Sprintf(sendMessageNonExistsUser, request.RecipientID),
				Code:  fiber.ErrBadRequest.Code}, err
		}
		return models.SendMessageResponse{Error: sendMessagePGError, Code: fiber.ErrInternalServerError.Code}, err
	}

	if err = u.connRedisRepo.SendMessage(ctx, request); err != nil {
		return models.SendMessageResponse{Error: sendMessageRedisError, Code: fiber.ErrInternalServerError.Code}, err
	}

	return
}
