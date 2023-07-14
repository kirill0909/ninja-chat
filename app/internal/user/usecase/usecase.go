package usecase

import (
	"context"
	"ninja-chat/config"
	models "ninja-chat/internal/models/user"
	"ninja-chat/internal/user"
)

type UserUsecase struct {
	cfg        *config.Config
	userPGRepo user.PGRepo
}

func NewUserUsecase(cfg *config.Config, userPGRepo user.PGRepo) user.Usecase {
	return &UserUsecase{cfg: cfg, userPGRepo: userPGRepo}
}

func (u *UserUsecase) Registration(ctx context.Context, req models.RegistrationRequest) (err error) {
	return u.userPGRepo.Registration(ctx, req)
}
