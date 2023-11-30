package controller

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewHealthController, NewUserController, NewAuthController,
		NewERPCategoryController, NewERPCustomerController, NewERPEmployeeManagementController,
		NewERPProductController, NewERPStoreController, NewOrderController, NewPromoteController,
	))
