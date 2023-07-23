package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"

	models "ninja-chat-core-api/internal/models/user"
)

func (md *MDWManager) AuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeaders := models.AuthHeaders{
			APIKey:      c.Get("ApiKey"),
			AccessToken: c.Get("AccessToken"),
		}

		if err := md.validateAuthHeaders(authHeaders); err != nil {
			log.Println(err)
			return err
		}

		return nil
	}
}
