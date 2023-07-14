package user

import (
	"context"
	"ninja-chat/internal/models/user"
)

type Repository interface {
	Registration(ctx context.Context, req models.RegistrationRequest) (err error)
}
