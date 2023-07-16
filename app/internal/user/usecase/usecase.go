package usecase

import (
	"context"
	"database/sql"
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"

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

func (u *UserUsecase) Login(ctx context.Context, req models.UserLoginRequest) (models.UserLoginResponse, error) {
	authData, err := u.userPGRepo.Login(ctx, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.UserLoginResponse{Error: "Unable to find user", Code: fiber.ErrNotFound.Code}, sql.ErrNoRows
		}
		return models.UserLoginResponse{Error: "Unable to get auth data", Code: fiber.ErrInternalServerError.Code}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(authData.PasswordHash), []byte(req.Password)); err != nil {
		return models.UserLoginResponse{Error: "Invalid password", Code: fiber.ErrBadRequest.Code}, err
	}

	var tokenData models.TokenData
	if tokenData, err = u.createSession(authData); err != nil {
		return models.UserLoginResponse{Error: "Unable to create session"}, err
	}

	if err = u.userRedisRepo.SaveUserSession(ctx, models.ClientSession{
		UserID:      authData.UserID,
		AccessToken: tokenData.AccessToken,
		ExpireAt:    int(time.Hour * 24),
	}); err != nil {
		return models.UserLoginResponse{Error: "Unable to save user sesssion", Code: fiber.ErrInternalServerError.Code}, err
	}

	return models.UserLoginResponse{Success: true, AccessToken: tokenData.AccessToken}, nil
}

func (u *UserUsecase) createSession(authData models.AuthData) (models.TokenData, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     authData.UserID,
		"expairedAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return models.TokenData{}, err
	}

	return models.TokenData{AccessToken: tokenString}, nil
}
