package user

import (
	"context"
	models "ninja-chat-core-api/internal/models/user"
)

type Usecase interface {
	Registration(ctx context.Context, req models.RegistrationRequest) (result models.RegistrationResponse, err error)
	Login(ctx context.Context, req models.UserLoginRequest) (result models.UserLoginResponse, err error)
	GetUserSession(ctx context.Context, req models.AuthHeaders) (result models.UserSession, err error)
	Logout(ctx context.Context, userID int) (result models.LogoutResponse, err error)
}
