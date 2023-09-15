package controller

import (
	erpcontroller "erp/api/controllers/erp"

<<<<<<< HEAD
var Module = fx.Options(
	fx.Invoke(NewHealthController, NewUserController, NewAuthController, NewCategoryController),
=======
	"go.uber.org/fx"
>>>>>>> develop
)

var Module = fx.Options(
	fx.Provide(NewHealthController, NewUserController, NewAuthController), erpcontroller.Module)
