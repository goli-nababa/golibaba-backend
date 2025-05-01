package transaction

import (
	"context"
	"gorm.io/gorm"
)

type txKey struct{}

var TransactionKey = txKey{}

func GetTransaction(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TransactionKey).(*gorm.DB); ok {
		return tx
	}
	return nil
}

func ContextWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TransactionKey, tx)
}
