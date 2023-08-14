package route

import (
	"erp-server/conf"
	"erp-server/handler"
	"erp-server/middleware"
	"erp-server/repo"
	"erp-server/service"
)

type Service struct {
	*conf.App
}

type IRoute interface {
	NewService() *Service
}

func NewService() *Service {
	s := Service{
		conf.NewApp(),
	}

	db := s.GetDB()
	repository := repo.NewRepo(db)

	userService := service.NewUser(repository)
	user := handler.NewUser(userService)

	businessService := service.NewBusiness(repository)
	business := handler.NewBusiness(businessService)

	productService := service.NewProduct(repository)
	product := handler.NewProduct(productService)

	orderService := service.NewOrder(repository)
	order := handler.NewOrder(orderService)

	moneyService := service.NewMoney(repository)
	money := handler.NewMoney(moneyService)

	migrate := handler.NewMigration(db)
	// migration

	router := s.Router
	v1 := router.Group("/v1")

	// auth
	v1.POST("/token", user.GenerateToken)
	v1.POST("/user/register", user.Register)

	// user
	secured := v1.Group("/secure").Use(middleware.Auth())
	{
		secured.GET("/user", handler.Ping)
	}

	// business
	v1.POST("/business", business.CreateBusiness)
	v1.PUT("/business", business.UpdateBusiness)
	v1.GET("/business", business.GetBusiness)

	// product
	v1.POST("/product", product.CreateProduct)
	v1.PUT("/product", product.UpdateProduct)
	v1.GET("/products", product.GetProducts)
	v1.GET("/product", product.GetProduct)

	// order
	v1.POST("/order", order.CreateOrder)
	v1.PUT("/order", order.UpdateOrder)
	v1.GET("/orders", order.GetOrders)
	v1.GET("/order", order.GetOrder)

	// money
	v1.GET("/money", money.GetMoney)
	v1.POST("/money", money.CreateMoney)
	v1.PUT("/money", money.UpdateMoney)
	v1.GET("/moneys", money.GetMoneys)
	v1.DELETE("/money", money.DeleteMoney)

	// migration
	router.POST("/migrate", migrate.Migrate)

	return &s
}
