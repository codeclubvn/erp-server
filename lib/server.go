package lib

import (
	"context"
	"erp/api/middlewares"
	config "erp/config"
	constants "erp/constants"
	"erp/infrastructure"
	models "erp/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

type Handler struct {
	*gin.RouterGroup
}

func NewServerGroup(instance *gin.Engine) *Handler {
	return &Handler{
		instance.Group("/v1/api"),
	}
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

	instance.Use(middlewares.JSONMiddleware)
	instance.Use(middlewares.CORS)
	instance.Use(middlewares.Logger)
	instance.Use(middlewares.ErrorHandler)
	// instance.Use(middlewares.JWT(config, db))

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.Info("Starting HTTP server")

			SeedRoutes(instance, db)
			go func() {
				//addr := fmt.Sprint(config.Server.Host, ":", config.Server.Port)
				addr := ":" + strconv.Itoa(config.Server.Port)
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
	// Seed permissions
	permissions := []models.Permission{}
	newPermissions := []models.Permission{}
	db.Find(&permissions)

	mapRoutes := make(map[string]models.Permission)
	for _, r := range permissions {
		mapRoutes[r.RoutePath] = r
	}

	for _, r := range engine.Routes() {
		// permission name
		// if method is GET, path is /api/v1/erp/users, permission name is get:users
		// if method is POST, path is /api/v1/erp/users, permission name is create:users
		// ...

		_, isPublic := constants.PublicRoutes[r.Path]
		_, isExist := mapRoutes[r.Path]

		s := strings.Split(r.Path, "/")

		if isExist {
			continue
		}

		if isPublic || s[1]+"/"+s[2]+"/"+s[3] != "v1/api/erp" {
			continue
		}

		last := s[len(s)-1]
		if last == "" {
			s = s[:len(s)-1]
		}

		permissionPrefix := ""
		switch r.Method {
		case "GET":
			permissionPrefix = "get"
		case "POST":
			permissionPrefix = "create"
		case "PUT":
			permissionPrefix = "update"
		case "DELETE":
			permissionPrefix = "delete"
		}

		pn := permissionPrefix + ":" + s[len(s)-1]

		newPermissions = append(newPermissions, models.Permission{
			Method:    r.Method,
			RoutePath: r.Path,
			Name:      pn,
		})
	}

	if len(newPermissions) == 0 {
		return nil
	}

	err := db.DB.Create(&newPermissions).Error
	if err != nil {
		panic(err)
	}
	return nil
}
