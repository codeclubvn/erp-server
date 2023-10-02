package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(
	NewUserRepository,
	NewCategoryRepository,
	NewCategoryProductRepository,
	NewErpPermissionRepo,
	NewErpRoleRepo,
	NewERPStoreRepository,
	NewERPProductRepository,
	NewERPCustomerRepository,
))
