package http

import (
	"log"
	"ninja-chat/config"
	models "ninja-chat/internal/models/user"
	"ninja-chat/internal/user"

	"github.com/gofiber/fiber/v2"
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
		if err := h.userUC.Registration(c.Context(), req); err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
