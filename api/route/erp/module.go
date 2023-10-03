package erproute

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewERPHandler), fx.Invoke(
	NewEmployeeManagementRoutes,
	NewStoreRoutes,
	NewCategoryRoutes,
	NewProductRoutes,
	NewCustomerRoutes,
))
