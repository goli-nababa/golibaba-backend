package logging

import (
	"api_gateway/pkg/logging/ports"
	"encoding/json"
	"fmt"
	"time"
)

type LogService struct {
	Publisher ports.LogPublisher
}

func NewLogService(publisher ports.LogPublisher) *LogService {
	return &LogService{
		Publisher: publisher,
	}
}

// LogRequest creates a log entry for an API request and publishes it.
func (s *LogService) LogRequest(userID, companyID uint, method, path string) error {
	logData := s.createLogData(userID, companyID, method, path)

	logBytes, err := json.Marshal(logData)
	if err != nil {
		return fmt.Errorf("failed to marshal log data: %w", err)
	}

	if err := s.Publisher.PublishLog(logBytes); err != nil {
		return fmt.Errorf("failed to publish log: %w", err)
	}

	return nil
}

// createLogData constructs the log entry with the provided details.
func (s *LogService) createLogData(userID, companyID uint, method, path string) map[string]interface{} {
	return map[string]interface{}{
		"user_id":    userID,
		"company_id": companyID,
		"method":     method,
		"path":       path,
		"timestamp":  time.Now().UTC(),
	}
}
