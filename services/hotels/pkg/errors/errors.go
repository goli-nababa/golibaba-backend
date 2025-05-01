package errors

import (
	"encoding/json"

	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrInternal     ErrorCode = "INTERNAL_ERROR"
	ErrInvalidInput ErrorCode = "INVALID_INPUT"
	ErrNotFound     ErrorCode = "NOT_FOUND"
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"
)

type ErrorDetails struct {
	ErrorCode    string                 `json:"code"`
	ErrorMessage string                 `json:"message"`
	Details      map[string]interface{} `json:"details,omitempty"`
}

type CustomError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	Details    map[string]interface{} `json:"details,omitempty"`
	StatusCode int                    `json:"-"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

func NewError(code ErrorCode, message string, statusCode int) *CustomError {
	return &CustomError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Details:    make(map[string]interface{}),
	}
}

func (e *CustomError) WithDetails(details map[string]interface{}) *CustomError {
	e.Details = details
	return e
}

func HTTPErrorHandler(err error, w http.ResponseWriter) {
	var customErr *CustomError

	switch e := err.(type) {
	case *CustomError:
		customErr = e
	default:
		customErr = NewError(ErrInternal, "خطای داخلی سرور", http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"success": false,
		"error": map[string]interface{}{
			"code":    customErr.Code,
			"message": customErr.Message,
			"details": customErr.Details,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(customErr.StatusCode)
	json.NewEncoder(w).Encode(response)
}
