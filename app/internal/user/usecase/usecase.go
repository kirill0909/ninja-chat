package usecase

import (
	"context"
	"database/sql"
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	cfg           *config.Config
	userPGRepo    user.PGRepo
	userRedisRepo user.RedisRepo
}

func NewUserUsecase(cfg *config.Config, userPGRepo user.PGRepo, userRedisRepo user.RedisRepo) user.Usecase {
	return &UserUsecase{cfg: cfg, userPGRepo: userPGRepo, userRedisRepo: userRedisRepo}
}

func (u *UserUsecase) Registration(ctx context.Context, req models.RegistrationRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	return u.userPGRepo.Registration(ctx, req)
}

func (u *UserUsecase) Login(ctx context.Context, req models.UserLoginRequest) (userID int, err error) {
	authData, err := u.userPGRepo.Login(ctx, req)
	if err != nil {
		return 0, err
	}
	if authData.UserID == 0 {
		return 0, sql.ErrNoRows
	}

	if err = bcrypt.CompareHashAndPassword([]byte(authData.PasswordHash), []byte(req.Password)); err != nil {
		return 0, err
	}

	return authData.UserID, nil
}
