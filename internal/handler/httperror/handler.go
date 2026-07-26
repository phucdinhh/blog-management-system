package httperror

import (
	"blog-management-system/internal/apperror"
	"blog-management-system/internal/constants"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type response struct {
	Error detail `json:"error"`
}

type detail struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func Handler(c *fiber.Ctx, err error) error {
	status, code, message := classify(err)
	requestID, _ := c.Locals(constants.RequestIDLocalKey).(string)

	event := log.With().
		Err(err).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("request_id", requestID).
		Str("error_code", code).
		Int("status", status).
		Logger()

	if status >= http.StatusInternalServerError {
		event.Error().Msg("request failed")
	} else {
		event.Warn().Msg("request failed")
	}

	return c.Status(status).JSON(response{
		Error: detail{
			Code:      code,
			Message:   message,
			RequestID: requestID,
		},
	})
}

func classify(err error) (int, string, string) {
	var appErr *apperror.Error
	if errors.As(err, &appErr) {
		return appErr.Status, appErr.Code, appErr.Message
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		switch fiberErr.Code {
		case http.StatusBadRequest:
			return http.StatusBadRequest, "BAD_REQUEST", fiberErr.Message
		case http.StatusNotFound:
			return http.StatusNotFound, "NOT_FOUND", "resource not found"
		case http.StatusServiceUnavailable:
			return http.StatusServiceUnavailable, "SERVICE_UNAVAILABLE", "service unavailable"
		default:
			return fiberErr.Code, "REQUEST_ERROR", fiberErr.Message
		}
	}

	return http.StatusInternalServerError, "INTERNAL_ERROR", "internal server error"
}
