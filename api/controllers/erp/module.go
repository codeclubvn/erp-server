package erpcontroller

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(
	NewERPEmployeeManagementController,
	NewERPStoreController,
	NewERPProductController,
	NewERPCategoryController,
	NewERPCustomerController,
	NewOrderController,
	NewPromoteController,
))
