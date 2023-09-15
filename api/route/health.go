package route

import (
	controller "erp/api/controllers"
	"erp/lib"
)

type HealthRoutes struct {
	handler *lib.Handler
}

func NewHealthRoutes(handler *lib.Handler, controller *controller.HealthController) *HealthRoutes {
	g := handler.Group("/health")

	g.GET("/", controller.Health)

	return &HealthRoutes{
		handler: handler,
	}
}
