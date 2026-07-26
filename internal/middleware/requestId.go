package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

func RequestIdMiddleware() fiber.Handler {
	return requestid.New(requestid.Config{
		Generator: utils.UUIDv4,
	})
}
