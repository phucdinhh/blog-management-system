package main

import (
	_ "blog-management-system/cmd/api/docs"
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

// @title Blog management system API
// @version 1.0
// @description This is an API document of Blog management system

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	app.Use(middleware.SwaggerMiddleware())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/panic", panicHandler)

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	log.Info().Msgf("Server starting on port %d", cfg.AppPort)
	log.Fatal().Err(app.Listen(addr))
}

// @Summary Health check panic endpoint
// @Description Used to verify error handling and route wiring.
// @Tags system
// @Produce json
// @Success 200 {object} map[string]string
// @Router /panic [get]
func panicHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "ok"})
}
