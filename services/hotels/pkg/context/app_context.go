package context

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type appContext struct {
	context.Context
	db     *gorm.DB
	logger *zap.Logger
}

type AppContextType func(*appContext) *appContext

var defaultLogger *zap.Logger

func WithDB(db *gorm.DB) AppContextType {
	return func(ac *appContext) *appContext {
		ac.db = db
		return ac
	}
}

func WithLogger(logger *zap.Logger) AppContextType {
	return func(ac *appContext) *appContext {
		ac.logger = logger
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

func SetLogger(ctx context.Context, logger *zap.Logger) {
	appCtx, ok := ctx.(*appContext)
	if ok {
		appCtx.logger = logger
	}
}

func GetLogger(ctx context.Context) *zap.Logger {
	appCtx, ok := ctx.(*appContext)
	if !ok || appCtx.logger == nil {
		return defaultLogger
	}
	return appCtx.logger
}
