package repository

import (
	"context"
	"ninja-chat-core-api/config"

	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"

	"github.com/jmoiron/sqlx"
)

type UserPGRepo struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewUserPGRepo(cfg *config.Config, db *sqlx.DB) user.PGRepo {
	return &UserPGRepo{cfg: cfg, db: db}
}

func (r *UserPGRepo) Registration(ctx context.Context, req models.RegistrationRequest) error {
	if _, err := r.db.ExecContext(ctx, queryRegistration, req.Login, req.Password); err != nil {
		return err
	}
	return nil
}

func (r *UserPGRepo) Login(ctx context.Context, req models.UserLoginRequest) (models.AuthData, error) {
	var result models.AuthData
	if err := r.db.GetContext(ctx, &result, queryLogin, req.Login); err != nil {
		return models.AuthData{}, err
	}

	return result, nil
}
