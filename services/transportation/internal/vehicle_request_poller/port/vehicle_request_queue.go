package port

import (
	"context"
	"transportation/internal/trip/domain"
)

type QueueRepo interface {
	PublishRequest(ctx context.Context, msg domain.VehicleRequest) error
}
