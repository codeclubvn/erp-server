package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(
	NewUserRepository,
	NewErpPermissionRepo,
	NewErpRoleRepo,
	NewERPStoreRepository,
	NewERPCustomerRepository,
))
