package middleware

import (
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() fiber.Handler {
	appEnv := os.Getenv("APP_ENV")

	var logger zerolog.Logger

	if appEnv == "development" {
		logger = zerolog.New(os.Stdout).Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
		}).With().Timestamp().Logger()
	} else {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	return fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
		Fields: []string{"latency", "status", "method", "path", "error"},
	})
}
