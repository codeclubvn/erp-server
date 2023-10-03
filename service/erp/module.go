package erpservice

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewERPEmployeeManagementService, NewERPStoreService, NewTransactionService))
