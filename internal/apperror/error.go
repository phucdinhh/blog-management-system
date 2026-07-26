package apperror

import "net/http"

type Error struct {
	Status  int
	Code    string
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}

	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func New(status int, code, message string, cause error) *Error {
	return &Error{
		Status:  status,
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

func Validation(message string, cause error) *Error {
	return New(http.StatusUnprocessableEntity, "VALIDATION_ERROR", message, cause)
}

func NotFound(message string) *Error {
	return New(http.StatusNotFound, "NOT_FOUND", message, nil)
}

func Internal(cause error) *Error {
	return New(http.StatusInternalServerError, "INTERNAL_ERROR", "internal server error", cause)
}
