package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"
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
	return u.userPGRepo.Login(ctx, req)
}
