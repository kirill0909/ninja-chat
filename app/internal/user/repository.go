package user

import (
	"context"
	models "ninja-chat-core-api/internal/models/user"
)

type PGRepo interface {
	Registration(ctx context.Context, req models.RegistrationRequest) (err error)
	Login(ctx context.Context, req models.UserLoginRequest) (userID int, err error)
}

type RedisRepo interface{}
