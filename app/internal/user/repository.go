package user

import (
	"context"
	models "ninja-chat-core-api/internal/models/user"
)

type PGRepo interface {
	Registration(ctx context.Context, req models.RegistrationRequest) (err error)
	Login(ctx context.Context, req models.UserLoginRequest) (result models.AuthData, err error)
}

type RedisRepo interface {
	SaveUserSession(context.Context, models.UserSession) error
	GetUserSession(ctx context.Context, req models.AuthHeaders) (result models.UserSession, err error)
	Logout(ctx context.Context, userID int) (result models.LogoutResponse, err error)
}
