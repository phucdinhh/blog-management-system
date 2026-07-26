package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func HealthCheckMiddleware(mongoClient *mongo.Client) fiber.Handler {
	return healthcheck.New(healthcheck.Config{
		LivenessEndpoint: "/live",
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/ready",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
				log.Err(err).Msgf("Readiness probe failed. MongoDB unreachable: %v", err)
				return false
			}

			return true
		},
	})
}
