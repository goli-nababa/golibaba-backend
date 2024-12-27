package port

import (
	"context"
	"transportation/internal/vehicle_request_poller/domain"
)

type Service interface {
	PollAndPublish(ctx context.Context, req domain.PollerRequest) error
}
