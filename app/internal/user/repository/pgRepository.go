package repository

import (
	"context"
	"log"
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

func (r *UserPGRepo) Registration(ctx context.Context, req models.RegistrationRequest) (err error) {
	log.Printf("%+v", req)
	return nil
}
