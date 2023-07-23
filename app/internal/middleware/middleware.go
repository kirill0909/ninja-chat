package middleware

import (
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"
	"strings"

	"github.com/pkg/errors"
)

type MDWManager struct {
	cfg    *config.Config
	userUC user.Usecase
}

func NewMDWManager(cfg *config.Config, userUC user.Usecase) *MDWManager {
	return &MDWManager{cfg: cfg, userUC: userUC}
}

func (m *MDWManager) validateNonAuthHeaders(headers models.NonAuthHeaders) error {
	if headers.APIKey != m.cfg.ApiKey {
		return errors.New("Invalid ApiKey for NonAuthHeaders")
	}
	return nil
}

func (m *MDWManager) validateAuthHeaders(headers models.AuthHeaders) error {
	if headers.APIKey != m.cfg.ApiKey {
		return errors.New("Invalid ApiKey for AuthHeaders")
	}
	if strings.TrimSpace(headers.AccessToken) == "" {
		return errors.New("Invalid AccessToken for AuthHeaders")
	}

	return nil
}
