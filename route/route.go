package route

import (
	controller "erp/api/controllers"
	"erp/api/middlewares"
	"erp/lib"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Invoke(
	NewRoute,
))

type Route struct {
	handler            *lib.Handler
	categoryController *controller.ERPCategoryController
	customerController *controller.ERPCustomerController
	authController     *controller.AuthController
	employeeController *controller.ERPEmployeeManagementController
	orderController    *controller.OrderController
	productController  *controller.ERPProductController
	promoteController  *controller.PromoteController
	storeController    *controller.ERPStoreController
	healthController   *controller.HealthController
	middleware         *middlewares.GinMiddleware
	revenueController  *controller.TransactionController
}

func NewRoute(
	handler *lib.Handler,
	categoryController *controller.ERPCategoryController,
	customerController *controller.ERPCustomerController,
	authController *controller.AuthController,
	employeeController *controller.ERPEmployeeManagementController,
	orderController *controller.OrderController,
	productController *controller.ERPProductController,
	promoteController *controller.PromoteController,
	storeController *controller.ERPStoreController,
	healthController *controller.HealthController,
	middleware *middlewares.GinMiddleware,
	transactionController *controller.TransactionController,
	walletController *controller.WalletController,
	budgetController *controller.BudgetController,
	transactionCategoryController *controller.TransactionCategoryController,
) *Route {

	handler.POST("/v1/auth/register", authController.Register)
	handler.POST("/v1/auth/login", authController.Login)

	handler.POST("/v1/product/", middleware.Auth(true), productController.Create)
	handler.PUT("/v1/product/", middleware.Auth(true), productController.Update)
	handler.DELETE("/v1/product/:id", middleware.Auth(true), productController.Delete)
	handler.GET("/v1/product/:id", middleware.Auth(false), productController.GetOne)
	handler.GET("/v1/product/", middleware.Auth(false), productController.GetList)

	handler.POST("/v1/category/", middleware.Auth(true), categoryController.Create)
	handler.PUT("/v1/category/", middleware.Auth(true), categoryController.Update)
	handler.GET("/v1/category/", middleware.Auth(false), categoryController.GetList)
	handler.GET("/v1/category/:id", middleware.Auth(false), categoryController.GetOne)
	handler.DELETE("/v1/category/:id", middleware.Auth(true), categoryController.Delete)

	handler.GET("/v1/customer/", middleware.Auth(false), customerController.GetList)
	handler.GET("/v1/customer/:id", middleware.Auth(false), customerController.GetOne)
	handler.POST("/v1/customer/", middleware.Auth(false), customerController.Create)
	handler.PUT("/v1/customer/:id", middleware.Auth(false), customerController.Update)
	handler.DELETE("/v1/customer/:id", middleware.Auth(false), customerController.Delete)

	handler.GET("/v1/permission/", middleware.Auth(false), employeeController.GetList)

	handler.POST("/v1/role/", middleware.Auth(true), employeeController.CreateRole)

	handler.POST("/v1/employee/", middleware.Auth(true), employeeController.CreateEmployee)

	handler.POST("/v1/order/", middleware.Auth(true), orderController.Create)
	handler.PUT("/v1/order/", middleware.Auth(true), orderController.Update)
	handler.GET("/v1/order/", middleware.Auth(true), orderController.GetList)
	handler.GET("/v1/order/:id", middleware.Auth(true), orderController.GetOne)

	handler.POST("/v1/promote/", middleware.Auth(true), promoteController.Create)

	handler.POST("/v1/store/", middleware.Auth(false), storeController.Create)
	handler.PUT("/v1/store/", middleware.Auth(true), storeController.Update)
	handler.GET("/v1/store/", middleware.Auth(false), storeController.List)
	handler.DELETE("/v1/store/", middleware.Auth(true), storeController.Delete)

	handler.POST("/v1/cashbook/", middleware.Auth(false), transactionController.Create)
	handler.PUT("/v1/cashbook/", middleware.Auth(true), transactionController.Update)
	handler.GET("/v1/cashbook/", middleware.Auth(false), transactionController.List)
	handler.DELETE("/v1/cashbook/:id", middleware.Auth(true), transactionController.Delete)
	handler.GET("/v1/cashbook/:id", middleware.Auth(true), transactionController.GetOne)

	handler.POST("/v1/debt-book/", middleware.Auth(false), transactionController.Create)
	handler.PUT("/v1/debt-book/", middleware.Auth(true), transactionController.Update)
	handler.GET("/v1/debt-book/", middleware.Auth(false), transactionController.List)
	handler.DELETE("/v1/debt-book/:id", middleware.Auth(true), transactionController.Delete)
	handler.GET("/v1/debt-book/:id", middleware.Auth(true), transactionController.GetOne)

	handler.POST("/v1/wallet/", middleware.Auth(false), walletController.Create)
	handler.PUT("/v1/wallet/", middleware.Auth(true), walletController.Update)
	handler.GET("/v1/wallet/", middleware.Auth(false), walletController.List)
	handler.DELETE("/v1/wallet/:id", middleware.Auth(true), walletController.Delete)
	handler.GET("/v1/wallet/:id", middleware.Auth(true), walletController.GetOne)

	handler.POST("/v1/budget/", middleware.Auth(false), budgetController.Create)
	handler.PUT("/v1/budget/", middleware.Auth(true), budgetController.Update)
	handler.GET("/v1/budget/", middleware.Auth(false), budgetController.List)
	handler.DELETE("/v1/budget/:id", middleware.Auth(true), budgetController.Delete)
	handler.GET("/v1/budget/:id", middleware.Auth(true), budgetController.GetOne)

	handler.POST("/v1/transaction_category/", middleware.Auth(false), transactionCategoryController.Create)
	handler.PUT("/v1/transaction_category/", middleware.Auth(true), transactionCategoryController.Update)
	handler.GET("/v1/transaction_category/", middleware.Auth(false), transactionCategoryController.List)
	handler.DELETE("/v1/transaction_category/:id", middleware.Auth(true), transactionCategoryController.Delete)
	handler.GET("/v1/transaction_category/:id", middleware.Auth(true), transactionCategoryController.GetOne)

	handler.GET("/v1/health/", healthController.Health)

	return &Route{
		handler: handler,
	}
}
