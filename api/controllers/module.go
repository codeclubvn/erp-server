package controller

import (
	erpcontroller "erp/api/controllers/erp"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHealthController, NewUserController, NewAuthController), erpcontroller.Module)
