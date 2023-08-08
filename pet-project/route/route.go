package route

import (
	"pet-project/config"
	"pet-project/handler"
	repo2 "pet-project/repo"
	"pet-project/service"
)

type Service struct {
	*config.App
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

	router := s.Router
	v1 := router.Group("/api/v1")

	// user
	v1.POST("/login", user.Login)

	// migration
	v1.POST("/migrate", migrate.Migrate)
	return &s
}
