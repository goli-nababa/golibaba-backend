package context

import (
	"context"

	"gorm.io/gorm"
)

type appContext struct {
	context.Context
	db *gorm.DB
}

type AppContextType func(*appContext) *appContext

func WithDB(db *gorm.DB) AppContextType {
	return func(ac *appContext) *appContext {
		ac.db = db
		return ac
	}
}

func NewAppContext(parent context.Context, opts ...AppContextType) context.Context {
	ctx := &appContext{Context: parent}
	for _, opt := range opts {
		ctx = opt(ctx)
	}
	return ctx
}

func SetDB(ctx context.Context, db *gorm.DB) {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return
	}
	appCtx.db = db
}

func GetDB(ctx context.Context) *gorm.DB {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return nil
	}
	return appCtx.db
}
