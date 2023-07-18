package usecase

import (
	"context"
	"database/sql"
	"github.com/jackc/pgconn"
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

func (u *UserUsecase) Registration(ctx context.Context, req models.RegistrationRequest) (models.RegistrationResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.RegistrationResponse{Code: fiber.ErrInternalServerError.Code}, err
	}

	req.Password = string(hashedPassword)

	if err = u.userPGRepo.Registration(ctx, req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return models.RegistrationResponse{Error: "This login already exists", Code: fiber.ErrBadRequest.Code}, err
		}
		return models.RegistrationResponse{Error: "Registration cannot be performed", Code: fiber.ErrInternalServerError.Code}, err
	}

	return models.RegistrationResponse{Success: true, Code: fiber.StatusOK}, nil
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
		return models.UserLoginResponse{Error: "Unable to create session", Code: fiber.ErrInternalServerError.Code}, err
	}

	if err = u.userRedisRepo.SaveUserSession(ctx, models.UserSession{
		UserID:      authData.UserID,
		AccessToken: tokenData.AccessToken,
		ExpireAt:    int(time.Hour * 24),
	}); err != nil {
		return models.UserLoginResponse{Error: "Unable to save user sesssion", Code: fiber.ErrInternalServerError.Code}, err
	}

	return models.UserLoginResponse{Success: true, Code: fiber.StatusOK, AccessToken: tokenData.AccessToken}, nil
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
