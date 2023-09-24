package lib

import (
	"context"
	"erp/api/middlewares"
	config "erp/config"
	constants "erp/constants"
	"erp/infrastructure"
	models "erp/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

func NewServerGroup(instance *gin.Engine) *gin.RouterGroup {
	return instance.Group("/api")
}

func NewServer(lifecycle fx.Lifecycle, zap *zap.Logger, config *config.Config, db *infrastructure.Database, middlewares *middlewares.GinMiddleware) *gin.Engine {
	switch config.Server.Env {
	case constants.Dev, constants.Local:
		gin.SetMode(gin.DebugMode)
	case constants.Prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	//gin.LoggerWithConfig(gin.LoggerConfig{
	//	Formatter: nil,
	//	Output:    nil,
	//	SkipPaths: nil,
	//})
	instance := gin.New()

	//instance.Use(gozap.RecoveryWithZap(zap, true))

	instance.Use(middlewares.ErrorHandler(zap))
	instance.Use(middlewares.JSONMiddleware)
	instance.Use(middlewares.CORS)
	instance.Use(middlewares.Logger(zap))
	instance.Use(middlewares.JWT(config))

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.Info("Starting HTTP server")

			SeedRoutes(instance, db)
			go func() {
				addr := fmt.Sprint(config.Server.Host, ":", config.Server.Port)
				if err := instance.Run(addr); err != nil {
					zap.Fatal(fmt.Sprint("HTTP server failed to start %w", err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.Info("Stopping HTTP server")
			return nil
		},
	})

	return instance
}

func SeedRoutes(engine *gin.Engine, db *infrastructure.Database) error {
	// Seed routes
	routes := []models.Routes{}

	// Delete all routes
	db.DB.Delete(&routes, "1=1")

	for _, r := range engine.Routes() {
		routes = append(routes, models.Routes{
			Method: r.Method,
			Path:   r.Path,
		})
	}

	err := db.DB.Create(&routes).Error
	if err != nil {
		panic(err)
	}
	return nil
}
