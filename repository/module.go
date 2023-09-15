package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(
	NewUserRepository,
<<<<<<< HEAD
	NewCategoryRepository,
	NewCategoryProductRepository,
)
=======
	NewErpPermissionRepo,
	NewErpRoleRepo,
	NewERPStoreRepository,
))
>>>>>>> develop
