package user

import (
	"context"
	models "ninja-chat-core-api/internal/models/user"
)

type Usecase interface {
	Registration(ctx context.Context, req models.RegistrationRequest) (err error)
	Login(ctx context.Context, req models.UserLoginRequest) (result models.UserLoginResponse, err error)
}
