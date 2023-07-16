package http

import (
	"log"
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"
	"ninja-chat-core-api/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"

	"github.com/pkg/errors"
)

type UserHandler struct {
	cfg    *config.Config
	userUC user.Usecase
}

func NewUserHandler(cfg *config.Config, userUC user.Usecase) user.Handler {
	return &UserHandler{cfg: cfg, userUC: userUC}
}

func (h *UserHandler) Registration() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req models.RegistrationRequest
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if err := h.userUC.Registration(c.Context(), req); err != nil {
			log.Printf("users.delivery.http.Registration:%s", err.Error())
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return c.Status(fiber.StatusBadRequest).JSON(models.RegistrationResponse{Error: "This login alread exists"})
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusOK).JSON(models.RegistrationResponse{Success: true})
	}
}

func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req models.UserLoginRequest
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := h.userUC.Login(c.Context(), req)
		if err != nil {
			log.Printf("user.delivery.http.Login: %s", err.Error())
			return c.Status(result.Code).JSON(result)
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}
