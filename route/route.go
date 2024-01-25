package route

import (
	"erp/cmd/lib"
	controller "erp/handler/controllers"
	"erp/handler/middlewares"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Invoke(
	NewRoute,
))

type Route struct {
	handler                   *lib.Handler
	categoryController        *controller.ERPCategoryController
	customerController        *controller.CustomerController
	authController            *controller.AuthController
	employeeController        *controller.ERPEmployeeManagementController
	orderController           *controller.OrderController
	productController         *controller.ERPProductController
	promoteController         *controller.PromoteController
	storeController           *controller.ERPStoreController
	healthController          *controller.HealthController
	middleware                *middlewares.GinMiddleware
	transactionController     *controller.CashbookController
	categoryProductController *controller.CategoryProductController
}

func NewRoute(
	handler *lib.Handler,
	categoryController *controller.ERPCategoryController,
	categoryProductController *controller.CategoryProductController,
	customerController *controller.CustomerController,
	authController *controller.AuthController,
	employeeController *controller.ERPEmployeeManagementController,
	orderController *controller.OrderController,
	productController *controller.ERPProductController,
	promoteController *controller.PromoteController,
	storeController *controller.ERPStoreController,
	healthController *controller.HealthController,
	middleware *middlewares.GinMiddleware,
	transactionController *controller.CashbookController,
	walletController *controller.WalletController,
	budgetController *controller.BudgetController,
	transactionCategoryController *controller.TransactionCategoryController,
) *Route {

	v1 := handler.Group("/v1")
	v1.POST("/auth/register", authController.Register)
	v1.POST("/auth/login", authController.Login)

	v1.POST("/product/", middleware.Auth(true), productController.Create)
	v1.PUT("/product/", middleware.Auth(true), productController.Update)
	v1.DELETE("/product/:id", middleware.Auth(true), productController.Delete)
	v1.GET("/product/:id", middleware.Auth(false), productController.GetOne)
	v1.GET("/product/", middleware.Auth(false), productController.GetList)

	v1.POST("/category/", middleware.Auth(true), categoryController.Create)
	v1.PUT("/category/", middleware.Auth(true), categoryController.Update)
	v1.GET("/category/", middleware.Auth(false), categoryController.GetList)
	v1.GET("/category/:id", middleware.Auth(false), categoryController.GetOne)
	v1.DELETE("/category/:id", middleware.Auth(true), categoryController.Delete)

	v1.POST("/category_product/", middleware.Auth(true), categoryProductController.Create)
	v1.PUT("/category_product/", middleware.Auth(true), categoryProductController.Update)
	v1.GET("/category_product/", middleware.Auth(false), categoryProductController.GetList)
	v1.DELETE("/category_product/:id", middleware.Auth(true), categoryProductController.Delete)

	v1.GET("/customer/", middleware.Auth(false), customerController.GetList)
	v1.GET("/customer/:id", middleware.Auth(false), customerController.GetOne)
	v1.POST("/customer/", middleware.Auth(false), customerController.Create)
	v1.PUT("/customer/", middleware.Auth(false), customerController.Update)
	v1.DELETE("/customer/:id", middleware.Auth(false), customerController.Delete)

	v1.GET("/permission/", middleware.Auth(false), employeeController.GetList)

	v1.POST("/role/", middleware.Auth(true), employeeController.CreateRole)

	v1.POST("/employee/", middleware.Auth(true), employeeController.CreateEmployee)

	v1.POST("/order/", middleware.Auth(true), orderController.Create)
	v1.PUT("/order/", middleware.Auth(true), orderController.Update)
	v1.GET("/order/", middleware.Auth(true), orderController.GetList)
	v1.GET("/order/:id", middleware.Auth(true), orderController.GetOne)
	v1.GET("/order/overview/", middleware.Auth(true), orderController.GetOverview)
	v1.GET("/order/best_seller", middleware.Auth(true), orderController.GetBestSeller)

	v1.POST("/promote/", middleware.Auth(true), promoteController.Create)

	v1.POST("/store/", middleware.Auth(false), storeController.Create)
	v1.PUT("/store/", middleware.Auth(true), storeController.Update)
	v1.GET("/store/", middleware.Auth(false), storeController.List)
	v1.DELETE("/store/", middleware.Auth(true), storeController.Delete)

	v1.POST("/cashbook/", middleware.Auth(false), transactionController.Create)
	v1.PUT("/cashbook/", middleware.Auth(true), transactionController.Update)
	v1.GET("/cashbook/", middleware.Auth(false), transactionController.List)
	v1.DELETE("/cashbook/:id", middleware.Auth(true), transactionController.Delete)
	v1.GET("/cashbook/:id", middleware.Auth(true), transactionController.GetOne)

	v1.GET("/debt/", middleware.Auth(false), transactionController.ListDebt)

	v1.POST("/wallet/", middleware.Auth(false), walletController.Create)
	v1.PUT("/wallet/", middleware.Auth(true), walletController.Update)
	v1.GET("/wallet/", middleware.Auth(false), walletController.List)
	v1.DELETE("/wallet/:id", middleware.Auth(true), walletController.Delete)
	v1.GET("/wallet/:id", middleware.Auth(true), walletController.GetOne)

	v1.POST("/budget/", middleware.Auth(false), budgetController.Create)
	v1.PUT("/budget/", middleware.Auth(true), budgetController.Update)
	v1.GET("/budget/", middleware.Auth(false), budgetController.List)
	v1.DELETE("/budget/:id", middleware.Auth(true), budgetController.Delete)
	v1.GET("/budget/:id", middleware.Auth(true), budgetController.GetOne)

	v1.POST("/cashbook_category/", middleware.Auth(false), transactionCategoryController.Create)
	v1.PUT("/cashbook_category/", middleware.Auth(true), transactionCategoryController.Update)
	v1.GET("/cashbook_category/", middleware.Auth(false), transactionCategoryController.List)
	v1.DELETE("/cashbook_category/:id", middleware.Auth(true), transactionCategoryController.Delete)
	v1.GET("/cashbook_category/:id", middleware.Auth(true), transactionCategoryController.GetOne)

	v1.GET("/health/", healthController.Health)

	return &Route{
		handler: handler,
	}
}
