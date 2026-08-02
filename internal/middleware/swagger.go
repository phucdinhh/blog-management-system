package middleware

import (
	"os"
	"path/filepath"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func SwaggerMiddleware() fiber.Handler {
	wd, _ := os.Getwd()
	swaggerPath := filepath.Join(wd, "cmd/api/docs/swagger.json")

	return swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: swaggerPath,
		Title:    "Blog management system API docs",
		Path:     "docs",
	})
}
