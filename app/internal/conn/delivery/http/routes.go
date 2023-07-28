package http

import (
	"github.com/gofiber/fiber/v2"
	"ninja-chat-core-api/internal/conn"
	"ninja-chat-core-api/internal/middleware"
)

func MapConnRoutes(connRoutes fiber.Router, mw *middleware.MDWManager, h conn.Handler) {
	connRoutes.Post("/send_message", mw.AuthedMiddleware(), h.SendMessage())
}
