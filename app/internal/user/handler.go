package user

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Registration() fiber.Handler
	Login() fiber.Handler
	Logout() fiber.Handler
}
