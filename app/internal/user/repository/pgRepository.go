package repository

import (
	"context"
	"ninja-chat/config"

	models "ninja-chat/internal/models/user"
	"ninja-chat/internal/user"

	"github.com/jmoiron/sqlx"
)

type UserPGRepo struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewUserPGRepo(cfg *config.Config, db *sqlx.DB) user.PGRepo {
	return &UserPGRepo{cfg: cfg, db: db}
}

func (r *UserPGRepo) Registration(ctx context.Context, req models.RegistrationRequest) (err error) {
	return nil
}
