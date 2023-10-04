package erpservice

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(
	NewERPEmployeeManagementService,
	NewERPStoreService,
	NewProductService,
	NewERPCustomerService,
	NewTransactionService,
	NewDebtService,
	NewPromoteService,
	NewOrderItemService,
	NewOrderService,
))
