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
	revenueController  *controller.RevenueController
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
	revenueController *controller.RevenueController,
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

	handler.GET("/v1/customer/", middleware.Auth(false), customerController.ListCustomer)
	handler.GET("/v1/customer/:id", middleware.Auth(false), customerController.CustomerDetail)
	handler.POST("/v1/customer/", middleware.Auth(false), customerController.CreateCustomer)
	handler.PUT("/v1/customer/:id", middleware.Auth(false), customerController.UpdateCustomer)
	handler.DELETE("/v1/customer/:id", middleware.Auth(false), customerController.DeleteCustomer)

	handler.GET("/v1/permission/", middleware.Auth(false), employeeController.ListPermission)

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

	handler.POST("/v1/revenue/", middleware.Auth(false), revenueController.Create)
	handler.PUT("/v1/revenue/", middleware.Auth(true), revenueController.Update)
	handler.GET("/v1/revenue/", middleware.Auth(false), revenueController.List)
	handler.DELETE("/v1/revenue/", middleware.Auth(true), revenueController.Delete)

	handler.GET("/v1/health/", healthController.Health)

	return &Route{
		handler: handler,
	}
}
