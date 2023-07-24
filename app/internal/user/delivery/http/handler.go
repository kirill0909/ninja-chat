package http

import (
	"log"
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"
	"ninja-chat-core-api/pkg/reqvalidator"

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
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if res := req.Validate(); !res {
			log.Printf("empty string for login(%s) or password(%s)", req.Login, req.Password)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := h.userUC.Registration(c.Context(), req)
		if err != nil {
			log.Printf("%s:%s users.delivery.http.Registration", err.Error(), result.Error)
			return c.Status(result.Code).JSON(result)
		}

		return c.Status(result.Code).JSON(result)
	}
}

func (h *UserHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req models.UserLoginRequest
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		result, err := h.userUC.Login(c.Context(), req)
		if err != nil {
			log.Printf("%s:%s user.delivery.http.Login", err.Error(), result.Error)
			return c.Status(result.Code).JSON(result)
		}

		return c.Status(result.Code).JSON(result)
	}
}

func (h *UserHandler) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {

		userID, ok := c.Locals("userID").(int)
		if !ok {
			log.Println("Cannot get userID from fiber ctx. user.delivery.http.Logout")
		}

		result, err := h.userUC.Logout(c.Context(), userID)
		if err != nil {
			log.Printf("%s:%s user.delivery.http.Logout", err.Error(), result.Error)
			return c.Status(result.Code).JSON(result)
		}

		return c.Status(result.Code).JSON(result)
	}
}
