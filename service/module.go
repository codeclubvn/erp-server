package service

import (
	erpservice "erp/service/erp"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(
	NewAuthService,
	NewUserService,
	NewJwtService,
	erpservice.NewCategoryService,
	erpservice.NewCategoryProductService,
), erpservice.Module)
