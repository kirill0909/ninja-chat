package http

import (
	"log"
	"ninja-chat-core-api/config"
	"ninja-chat-core-api/internal/conn"
	"ninja-chat-core-api/pkg/reqvalidator"

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
		userID, ok := c.Locals("userID").(int)
		if !ok {
			log.Println("Cannot cust userID from fiber ctx to int. conn.delivery.http.SendMessage")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		request.UserID = userID

		if err := reqvalidator.ReadRequest(c, &request); err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := h.connUC.SendMessage(c.Context(), request)
		if err != nil {
			log.Printf("%s:%s conn.delivery.http.SendMessage", err.Error(), result.Error)
			return c.Status(result.Code).JSON(result)
		}

		return c.Status(result.Code).JSON(result)
	}
}
