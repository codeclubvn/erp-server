package middlewares

import (
	"erp/cmd/infrastructure"
	config "erp/config"
	"go.uber.org/zap"
)

type GinMiddleware struct {
	logger *zap.Logger
	config *config.Config
	db     *infrastructure.Database
}

func NewMiddleware(config *config.Config, db *infrastructure.Database, logger *zap.Logger) *GinMiddleware {
	return &GinMiddleware{
		logger: logger,
		config: config,
		db:     db,
	}
}
