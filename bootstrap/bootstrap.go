package bootstrap

import (
	controller "erp/api/controllers"
	"erp/api/middlewares"
	"erp/api/route"
	config "erp/config"
	infrastructure "erp/infrastructure"
	"erp/lib"
	repository "erp/repository"
	service "erp/service"
	utils "erp/utils"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func inject() fx.Option {
	return fx.Options(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		//fx.NopLogger,
		fx.Provide(
			config.NewConfig,
			utils.NewTimeoutContext,
		),
		route.Module,
		lib.Module,
		repository.Module,
		service.Module,
		controller.Module,
		middlewares.Module,
		infrastructure.Module,
	)
}

func Run() {
	fx.New(inject()).Run()
}
