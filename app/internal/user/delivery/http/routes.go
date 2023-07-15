package http

import (
	"github.com/gofiber/fiber/v2"
	"ninja-chat-core-api/internal/middleware"
	"ninja-chat-core-api/internal/user"
)

func MapUserRoutes(userRoutes fiber.Router, mw *middleware.MDWManager, h user.Handler) {
	userRoutes.Post("/registration", mw.NonAuthedMiddleware(), h.Registration())
}
