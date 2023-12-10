package service

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(
	NewAuthService, NewUserService, NewJwtService,
	NewCategoryService, NewCategoryProductService, NewProductService,
	NewCustomerService, NewOrderService, NewOrderItemService,
	NewPromoteService, NewStoreService, NewERPEmployeeManagementService,
	NewCashbookService, NewTransactionCategoryService,
	NewWalletService, NewBudgetService,
))
