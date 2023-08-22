package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewServer), fx.Provide(NewDatabase))
