package conn

import "github.com/gofiber/fiber/v2"

type Handler interface {
	SaveMessage() fiber.Handler
}
