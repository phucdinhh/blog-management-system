package http

import (
	"blog-management-system/internal/handler/response"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type HealthHandler struct {
	mongoClient *mongo.Client
}

func NewHealthHandler(mongoClient *mongo.Client) *HealthHandler {

	return &HealthHandler{
		mongoClient: mongoClient,
	}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	if err := h.mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Err(err).Msg("ping MongoDB")
		return c.Status(fiber.StatusServiceUnavailable).JSON(response.ErrorResponse{
			Error: response.ErrorBody{
				Code:    fiber.ErrServiceUnavailable.Code,
				Message: "MongoDB connection failed",
				Details: err.Error(),
			},
		})
	}

	log.Info().Msg("health check OK")
	return c.Status(fiber.StatusOK).JSON(response.Response[string]{
		Data: "Connect successfully",
	})
}
