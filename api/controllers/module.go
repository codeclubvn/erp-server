package controller

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(NewHealthController, NewUserController, NewAuthController, NewCategoryController),
)
