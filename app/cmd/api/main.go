package main

import (
	"context"
	"log"
	"ninja-chat/config"
	"ninja-chat/internal/server"
	httpUser "ninja-chat/internal/user/delivery/http"
	pgRepoUser "ninja-chat/internal/user/repository"
	usecaseUser "ninja-chat/internal/user/usecase"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

func main() {
	viper, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(viper)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config loaded")

	app, deps := mapHandler(cfg)
	server := server.NewServer(app, deps, cfg)

	ctx := context.Background()
	if err := server.Run(ctx); err != nil {
		log.Println(err)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	server.Shutdown()
}

func mapHandler(cfg *config.Config) (*fiber.App, server.Deps) {
	// create App
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(logger.New())

	// repository
	userPGRepo := pgRepoUser.NewUserPGRepo(cfg, &sqlx.DB{}) // TODO: rewrite kostil

	// usecase
	userUC := usecaseUser.NewUserUsecase(cfg, userPGRepo)

	// handler
	userHTTP := httpUser.NewUserHandler(cfg, userUC)
	// productGRPC := grpcProduct.NewProductHandler(productUC)

	// groups
	apiGroup := app.Group("api")
	userGroup := apiGroup.Group("user")

	// routes
	httpUser.MapUserRoutes(userGroup, userHTTP)

	// create grpc dependencyes
	deps := server.Deps{ /* ProductDeps: productGRPC */ }
	return app, deps
}
