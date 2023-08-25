package bootstrap

import (
	controller "erp/api/controllers"
	"erp/api/middlewares"
	config "erp/config"
	infrastructure "erp/infrastructure"
	lib "erp/lib"
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
