package usecase

import (
	"context"
	"fmt"
	"ninja-chat-core-api/config"
	"ninja-chat-core-api/internal/conn"
	models "ninja-chat-core-api/internal/models/conn"

	"github.com/gofiber/fiber/v2"
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

func (u *ConnUsecase) SaveMessage(ctx context.Context, request models.SaveMessageRequest) (
	result models.SaveMessageResponse, err error) {

	result, err = u.connPGRepo.SaveMessage(ctx, request)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == violatesForeignKeyCode {
			return models.SaveMessageResponse{
				Error: fmt.Sprintf(saveMessageForNonExistsUser, request.UserID, request.Message, request.RecipientID),
				Code:  fiber.ErrBadRequest.Code}, err
		}
		return models.SaveMessageResponse{Error: saveMessagePGError, Code: fiber.ErrInternalServerError.Code}, err
	}
	request.MessageID = result.MessageID

	if err = u.connRedisRepo.SaveMessage(ctx, request); err != nil {
		return models.SaveMessageResponse{Error: saveMessageRedisError, Code: fiber.ErrInternalServerError.Code}, err
	}

	return
}
