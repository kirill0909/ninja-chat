package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	models "ninja-chat-core-api/internal/models/user"
)

func (m *MDWManager) NonAuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		nonAuthHeaders := models.NonAuthHeaders{
			APIKey: c.Get("ApiKey"),
		}

		if err := m.validateNonAuthHeaders(nonAuthHeaders); err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
