package repository

import (
	"context"
	"erp/infrastructure"
	models "erp/models"
)

type CategoryProductRepository interface {
	Create(ctx context.Context, CategoryProduct *models.CategoryProduct) (err error)
	Update(ctx context.Context, category *models.CategoryProduct) (err error)
	Delete(ctx context.Context, id string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.CategoryProduct, err error)
	GetList(ctx context.Context, category models.CategoryProduct) (res *models.CategoryProduct, err error)
}

type catProRepoImpl struct {
	*infrastructure.Database
}

func NewCategoryProductRepository(db *infrastructure.Database) CategoryProductRepository {
	if db == nil {
		panic("Database engine is null")
	}
	return &catProRepoImpl{db}
}

func (u *catProRepoImpl) Create(ctx context.Context, category *models.CategoryProduct) (err error) {
	err = u.DB.Create(&category).Error
	return err
}

func (u *catProRepoImpl) Update(ctx context.Context, category *models.CategoryProduct) (err error) {
	err = u.DB.Updates(&category).Error
	return err
}

func (u *catProRepoImpl) Delete(ctx context.Context, id string) (err error) {
	err = u.DB.Where("id = ?", id).Delete(&models.CategoryProduct{}).Error
	return err
}

func (u *catProRepoImpl) GetOneByID(ctx context.Context, id string) (res *models.CategoryProduct, err error) {
	err = u.DB.Where("id = ?", id).First(&res).Error
	return res, err
}

// todo: fix this
func (u *catProRepoImpl) GetList(ctx context.Context, category models.CategoryProduct) (res *models.CategoryProduct, err error) {
	err = u.DB.Find(&category).Error
	return &category, err
}
