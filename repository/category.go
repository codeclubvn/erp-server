package repository

import (
	"context"
	"erp/infrastructure"
	models "erp/models"
)

type CategoryRepository interface {
	Create(ctx context.Context, Category *models.Category) (err error)
	Update(ctx context.Context, category *models.Category) (err error)
	Delete(ctx context.Context, id string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.Category, err error)
	GetList(ctx context.Context, category models.Category) (res *models.Category, err error)
}

type categoryRepositoryImpl struct {
	*infrastructure.Database
}

func NewCategoryRepository(db *infrastructure.Database) CategoryRepository {
	if db == nil {
		panic("Database engine is null")
	}
	return &categoryRepositoryImpl{db}
}

func (u *categoryRepositoryImpl) Create(ctx context.Context, category *models.Category) (err error) {
	err = u.DB.Create(&category).Error
	return err
}

func (u *categoryRepositoryImpl) Update(ctx context.Context, category *models.Category) (err error) {
	err = u.DB.Updates(&category).Error
	return err
}

func (u *categoryRepositoryImpl) Delete(ctx context.Context, id string) (err error) {
	err = u.DB.Where("id = ?", id).Delete(&models.Category{}).Error
	return err
}

func (u *categoryRepositoryImpl) GetOneByID(ctx context.Context, id string) (res *models.Category, err error) {
	err = u.DB.Where("id = ?", id).First(&res).Error
	return res, err
}

// todo:
func (u *categoryRepositoryImpl) GetList(ctx context.Context, category models.Category) (res *models.Category, err error) {
	err = u.DB.Find(&category).Error
	return &category, err
}
