package lib

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewZapLogger, NewServer, NewServerGroup, NewCloudinaryRepository, NewSendinblue))
