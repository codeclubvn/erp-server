package bootstrap

import (
	"erp/cmd/infrastructure"
	"erp/cmd/lib"
	"erp/config"
	controller "erp/handler/controllers"
	"erp/handler/middlewares"
	"erp/repository"
	"erp/route"
	"erp/service"
	"erp/utils"

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
		route.Module,
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
