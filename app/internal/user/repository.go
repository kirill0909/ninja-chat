package user

import (
	"context"
	"ninja-chat-core-api/internal/models/user"
)

type PGRepo interface {
	Registration(ctx context.Context, req models.RegistrationRequest) (err error)
}
