package service

import (
	"context"
	config "erp/config"
	erpdto "erp/dto/erp"
	models "erp/models"
	repository "erp/repository"
	"erp/utils"
)

type (
	ERPCategoryProductService interface {
		Create(ctx context.Context, req erpdto.CategoryProductRequest) (*models.CategoryProduct, error)
		Update(ctx context.Context, req erpdto.CategoryProductRequest) (*models.CategoryProduct, error)
		Delete(ctx context.Context, id string) error
		GetList(ctx context.Context, req erpdto.GetListCatProRequest) ([]*models.CategoryProduct, *int64, error)
	}
	CategoryProductServiceImpl struct {
		catProductRepo repository.CategoryProductRepository
		config         *config.Config
	}
)

func NewCategoryProductService(catProductRepo repository.CategoryProductRepository, config *config.Config) ERPCategoryProductService {
	return &CategoryProductServiceImpl{
		catProductRepo: catProductRepo,
		config:         config,
	}
}

func (u *CategoryProductServiceImpl) Create(ctx context.Context, req erpdto.CategoryProductRequest) (*models.CategoryProduct, error) {
	res := models.CategoryProduct{}
	var err error

	if err = utils.Copy(req, res); err != nil {
		return nil, err
	}
	if err = u.catProductRepo.Create(nil, ctx, &res); err != nil {
		return nil, err
	}

	return &res, err
}

func (u *CategoryProductServiceImpl) Update(ctx context.Context, req erpdto.CategoryProductRequest) (*models.CategoryProduct, error) {
	res := models.CategoryProduct{}
	var err error

	if err = utils.Copy(req, res); err != nil {
		return nil, err
	}
	if err = u.catProductRepo.Update(nil, ctx, &res); err != nil {
		return nil, err
	}

	return &res, err
}

func (u *CategoryProductServiceImpl) Delete(ctx context.Context, id string) error {
	return u.catProductRepo.Delete(nil, ctx, id)
}

func (u *CategoryProductServiceImpl) GetList(ctx context.Context, req erpdto.GetListCatProRequest) ([]*models.CategoryProduct, *int64, error) {
	return u.catProductRepo.GetList(ctx, req)
}
