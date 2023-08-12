package route

import (
	"pet-project/config"
	"pet-project/handler"
	"pet-project/middleware"
	repo2 "pet-project/repo"
	"pet-project/service"
)

type Service struct {
	*config.App
}

type IRoute interface {
	NewService() *Service
}

func NewService() *Service {
	s := Service{
		config.NewApp(),
	}

	db := s.GetDB()
	repo := repo2.NewRepo(db)

	userService := service.NewUser(repo)
	user := handler.NewUser(userService)
	migrate := handler.NewMigration(db)
	// migration

	router := s.Router
	api := router.Group("/api/v1")
	{
		api.POST("/migrate", migrate.Migrate)
		api.POST("/token", handler.IUser(user).GenerateToken)
		api.POST("/user/register", user.Register)
		// user
		secured := api.Group("/secure").Use(middleware.Auth())
		{
			secured.GET("/user", handler.Ping)
		}

	}
	return &s
}
