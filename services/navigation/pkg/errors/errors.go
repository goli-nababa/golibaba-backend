package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func httpToGRPCCode(httpCode int) codes.Code {
	switch httpCode {
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	default:
		return codes.Unknown
	}
}

func GRPCErrorHandler(err error) error {
	var customErr *CustomError

	var e *CustomError

	switch {
	case errors.As(err, &e):
		customErr = e
	default:
		customErr = NewError(ErrInternal, err.Error(), http.StatusInternalServerError)
	}

	grpcCode := httpToGRPCCode(customErr.StatusCode)

	errDetails := &ErrorDetails{
		ErrorCode:    string(customErr.Code),
		ErrorMessage: customErr.Message,
		Details:      customErr.Details,
	}

	detailsBytes, _ := json.Marshal(errDetails)

	_ = detailsBytes

	st := status.New(grpcCode, string(detailsBytes))
	return st.Err()
}
