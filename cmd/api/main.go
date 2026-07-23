package main

import (
	"blog-management-system/internal/config"
	"blog-management-system/internal/handler/http"
	"blog-management-system/internal/middleware"
	"blog-management-system/internal/platform/mongo"
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize configuration")
	}

	mongoClient := mongo.Config(cfg)
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Error().Err(err).Msg("disconnect from MongoDB")
		}
	}()

	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	api := app.Group("/api", middleware.LoggerMiddleware())
	v1 := api.Group("/v1")
	v1.Get("/health", http.NewHealthHandler(mongoClient).Check)

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	log.Info().Msgf("Server starting on port %d", cfg.AppPort)
	log.Fatal().Err(app.Listen(addr))
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		status = fiberErr.Code
	}

	log.Error().Err(err).Str("path", ctx.Path()).Msg("request failed")
	return ctx.Status(status).JSON(fiber.Map{"message": fiber.ErrInternalServerError.Message})
}
