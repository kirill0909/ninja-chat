package main

import (
	"context"
	"log"
	"ninja-chat-core-api/config"
	"ninja-chat-core-api/internal/server"
	httpUser "ninja-chat-core-api/internal/user/delivery/http"
	pgRepoUser "ninja-chat-core-api/internal/user/repository"
	usecaseUser "ninja-chat-core-api/internal/user/usecase"
	"ninja-chat-core-api/pkg/storage/postgres"
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

	ctx := context.Background()
	psqlDB, err := postgres.InitPsqlDB(ctx, cfg)
	if err != nil {
		log.Printf("PostgreSQL error connection: %s", err.Error())
		return
	} else {
		log.Println("PostgreSQL successful connection")
	}

	app, deps := mapHandler(cfg, psqlDB)
	server := server.NewServer(app, deps, cfg)

	if err := server.Run(ctx); err != nil {
		log.Println(err)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	server.Shutdown()
}

func mapHandler(cfg *config.Config, db *sqlx.DB) (*fiber.App, server.Deps) {
	// create App
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(logger.New())

	// repository
	userPGRepo := pgRepoUser.NewUserPGRepo(cfg, db)

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
