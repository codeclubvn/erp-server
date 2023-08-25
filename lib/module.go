package lib

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewZapLogger), fx.Provide(NewServerGroup), fx.Provide(NewServer))
