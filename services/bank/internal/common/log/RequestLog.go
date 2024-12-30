package log

import (
	"bank_service/pkg/logging"
	"time"
)

type RequestLog struct {
	UserID       uint          `json:"user_id,omitempty"`
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	IP           string        `json:"ip"`
	StatusCode   int           `json:"status_code"`
	ResponseTime time.Duration `json:"response_time"`
	Error        string        `json:"error,omitempty"`
	HandlerLogs  []HandlerLog  `json:"handler_logs,omitempty"`
	RequestBody  interface{}   `json:"request_body,omitempty"`
	ResponseBody interface{}   `json:"response_body,omitempty"`
	StartTime    time.Time     `json:"start_time"`
}

type HandlerLog struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type contextKey string

const (
	LoggerKey     contextKey = "logger"
	RequestLogKey contextKey = "request_log"
)

const (
	CategoryHTTP   logging.Category    = "http"
	SubCategoryAPI logging.SubCategory = "api"
)
