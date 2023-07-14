package usecase

import (
	"ninja-chat/config"
	"ninja-chat/internal/user"
	"sync"
)

var (
	count int64
)

type ProductUsecase struct {
	cfg        *config.Config
	productMap sync.Map
}

func NewProducetUsecase(cfg *config.Config) user.Usecase {
	return &ProductUsecase{cfg: cfg}
}
