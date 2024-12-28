package port

import (
	"context"
)

type WalletEventPublisher interface {
	PublishEvent(ctx context.Context, event interface{}) error
}
