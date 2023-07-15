package usecase

import (
	"context"
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	cfg        *config.Config
	userPGRepo user.PGRepo
}

func NewUserUsecase(cfg *config.Config, userPGRepo user.PGRepo) user.Usecase {
	return &UserUsecase{cfg: cfg, userPGRepo: userPGRepo}
}

func (u *UserUsecase) Registration(ctx context.Context, req models.RegistrationRequest) error {
	// TODO: add solt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	return u.userPGRepo.Registration(ctx, req)
}
