package route

import (
	erproute "erp/api/route/erp"

	"go.uber.org/fx"
)

var Module = fx.Options(fx.Invoke(NewAuthRoutes, NewUserRoutes, NewHealthRoutes), erproute.Module)
