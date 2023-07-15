package middleware

import (
	"ninja-chat-core-api/config"
	models "ninja-chat-core-api/internal/models/user"
	"ninja-chat-core-api/internal/user"

	"github.com/pkg/errors"
)

var (
	ApiKey = "NinjaApiKey" // TODO: remove kostil
)

type MDWManager struct {
	cfg    *config.Config
	userUC user.Usecase
}

func NewMDWManager(cfg *config.Config, userUC user.Usecase) *MDWManager {
	return &MDWManager{cfg: cfg, userUC: userUC}
}

func (m *MDWManager) validateHeaders(headers models.NonAuthHeaders) error {
	if headers.APIKey != ApiKey {
		return errors.New("Invalid ApiKey for NonAuthHeaders")
	}
	return nil
}
