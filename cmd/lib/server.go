package lib

import (
	"context"
	config "erp/config"
	"erp/handler/middlewares"
	"erp/utils/constants"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

type Handler struct {
	*gin.RouterGroup
}

func NewServerGroup(instance *gin.Engine) *Handler {
	return &Handler{
		instance.Group("/handler/"),
	}
}

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}

func NewServer(lifecycle fx.Lifecycle, zap *zap.Logger, config *config.Config, middlewares *middlewares.GinMiddleware) *gin.Engine {
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

	instance.Use(middlewares.JSONMiddleware)
	instance.Use(CORS())
	instance.Use(middlewares.Logger)
	instance.Use(middlewares.ErrorHandler)
	// instance.Use(middlewares.JWT(config, db))

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.Info("Starting HTTP server")

			//if err := SeedRoutes(instance, db); err != nil {
			//	zap.Fatal(fmt.Sprint("HTTP server failed to start %w", err))
			//}
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
