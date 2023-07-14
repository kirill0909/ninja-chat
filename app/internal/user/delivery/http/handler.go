package http

import (
	"ninja-chat/config"
	"ninja-chat/internal/user"
)

type ProductHandler struct {
	cfg       *config.Config
	productUC user.Usecase
}

func NewProductHandler(cfg *config.Config, productUC user.Usecase) user.Handler {
	return &ProductHandler{cfg: cfg, productUC: productUC}
}
