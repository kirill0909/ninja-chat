package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"ninja-chat/config"
	"ninja-chat/internal/server"
	httpUser "ninja-chat/internal/user/delivery/http"
	usecaseUser "ninja-chat/internal/user/usecase"
	"os"
	"os/signal"
	"syscall"
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

	// usecase
	productUC := usecaseUser.NewProducetUsecase(cfg)

	// handler
	productHTTP := httpUser.NewProductHandler(cfg, productUC)
	// productGRPC := grpcProduct.NewProductHandler(productUC)

	// groups
	apiGroup := app.Group("api")
	productGroup := apiGroup.Group("product")

	// routes
	httpUser.MapProductRoutes(productGroup, productHTTP)

	// create grpc dependencyes
	deps := server.Deps{ /* ProductDeps: productGRPC */ }
	return app, deps
}
