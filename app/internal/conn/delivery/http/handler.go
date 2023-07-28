package http

import (
	"log"
	"ninja-chat-core-api/config"
	"ninja-chat-core-api/internal/conn"

	models "ninja-chat-core-api/internal/models/conn"

	"github.com/gofiber/fiber/v2"
)

type ConnHandler struct {
	cfg    *config.Config
	connUC conn.Usecase
}

func NewConnHandler(cfg *config.Config, connUC conn.Usecase) conn.Handler {
	return &ConnHandler{cfg: cfg, connUC: connUC}
}

func (h *ConnHandler) SendMessage() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.SendMessageRequest
		result, err := h.connUC.SendMessage(c.Context(), request)
		if err != nil {
			log.Println(err)
			return c.Status(result.Code).JSON(result)
		}

		return c.Status(result.Code).JSON(result)
	}
}
