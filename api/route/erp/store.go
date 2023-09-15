package erproute

import (
	erpcontroller "erp/api/controllers/erp"
	"erp/api/middlewares"
)

type StoreRoutes struct {
	handler *ERPHandler
}

func NewStoreRoutes(handler *ERPHandler, controller *erpcontroller.ERPStoreController, middleware *middlewares.GinMiddleware) *StoreRoutes {
	g := handler.Group("/store")

	g.POST("/", middleware.Auth(false), controller.CreateStore)
	g.PUT("/", middleware.Auth(true), controller.UpdateStore)
	g.GET("/", middleware.Auth(false), controller.ListStore)
	g.DELETE("/", middleware.Auth(true), controller.DeleteStore)

	return &StoreRoutes{
		handler: handler,
	}
}
