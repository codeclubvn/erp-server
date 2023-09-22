package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type CategoryProductRoutes struct {
	handler *ERPHandler
}

func NewCategoryProductRoutes(handler *ERPHandler, controller *erpcontroller.ERPCategoryProductController, middleware *middlewares.GinMiddleware) *CategoryProductRoutes {
	g := handler.Group("/category_product")

	g.POST("/", middleware.Auth(true), controller.Create)
	g.PUT("/", middleware.Auth(true), controller.Update)
	g.GET("/", middleware.Auth(false), controller.GetList)
	g.DELETE("/:id", middleware.Auth(true), controller.Delete)

	return &CategoryProductRoutes{
		handler: handler,
	}
}
