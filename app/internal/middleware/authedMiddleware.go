package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	models "ninja-chat-core-api/internal/models/user"
)

func (md *MDWManager) AuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeaders := models.AuthHeaders{
			APIKey:      c.Get("ApiKey"),
			AccessToken: c.Get("AccessToken"),
		}

		if err := md.validateAuthHeaders(authHeaders); err != nil {
			log.Printf("%s: middleware.AuthedMiddleware.validateAuthHeaders", err.Error())
			return c.SendStatus(fiber.StatusBadRequest)
		}

		userSession, err := md.userUC.GetUserSession(c.Context(), authHeaders)
		if err != nil {
			log.Printf("%s: middleware.AuthedMiddleware.GetUserSession", err.Error())
			if errors.Is(err, redis.Nil) {
				return c.SendStatus(fiber.StatusNotFound)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		userID, expiredAt, err := md.parseToken(userSession.AccessToken)
		if err != nil {
			log.Printf("%s: middleware.AuthedMiddleware.parseToken", err.Error())
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if userID == 0 {
			log.Println("userID == 0. middleware.AuthedMiddleware")
			return c.SendStatus(fiber.StatusNotFound)
		}

		// TODO: check this. Because redis deletes the keys that have expired on its own
		if int64(expiredAt) <= time.Now().Unix() {
			log.Printf("session time(%d) expired for user(%d). middleware.AuthedMiddleware", expiredAt, userID)
			return c.SendStatus(fiber.StatusBadRequest)
		}

		c.Locals("userID", userID)
		return c.Next()
	}
}

func (md *MDWManager) parseToken(tokenString string) (userID, expiredAt int, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(md.cfg.JWTSecret), nil
	})
	if err != nil {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tmp, ok := claims["userID"].(float64); ok {
			userID = int(tmp)
		}
		if tmp, ok := claims["expairedAt"].(float64); ok {
			expiredAt = int(tmp)
		}
		return
	} else {
		err = errors.New("unable to parse token claims")
		return
	}
}
