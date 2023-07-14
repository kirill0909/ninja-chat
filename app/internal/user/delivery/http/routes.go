package http

import (
	"github.com/gofiber/fiber/v2"
	"ninja-chat-core-api/internal/user"
)

func MapUserRoutes(userRoutes fiber.Router, h user.Handler) {
	userRoutes.Post("/registration", h.Registration())
}
