package main

import (
	"blog-management-system/internal/config"
	"blog-management-system/internal/handler/httperror"
	"blog-management-system/internal/middleware"
	"blog-management-system/internal/platform/mongo"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	app := fiber.New(fiber.Config{ErrorHandler: httperror.Handler})

	app.Use(middleware.RequestIdMiddleware())
	app.Use(middleware.LoggerMiddleware())
	app.Use(recover.New())
	app.Use(middleware.HealthCheckMiddleware(mongoClient))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/panic", func(c *fiber.Ctx) error {
		panic("I'm an error")
	})

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	log.Info().Msgf("Server starting on port %d", cfg.AppPort)
	log.Fatal().Err(app.Listen(addr))
}
