package erpservice

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
		Delete(ctx context.Context, id string, userId string) error
		GetList(ctx context.Context, req erpdto.GetListCatProRequest) (*erpdto.CatProductsResponse, error)
	}
	CategoryProductServiceImpl struct {
		categoryProductRepo repository.CategoryProductRepository
		config              *config.Config
	}
)

func NewCategoryProductService(itemRepo repository.CategoryProductRepository, config *config.Config) ERPCategoryProductService {
	return &CategoryProductServiceImpl{
		categoryProductRepo: itemRepo,
		config:              config,
	}
}

func (u *CategoryProductServiceImpl) Create(ctx context.Context, req erpdto.CategoryProductRequest) (*models.CategoryProduct, error) {
	res := models.CategoryProduct{}
	var err error

	if err = utils.Copy(req, res); err != nil {
		return nil, err
	}
	if err = u.categoryProductRepo.Create(ctx, &res); err != nil {
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
	if err = u.categoryProductRepo.Update(ctx, &res); err != nil {
		return nil, err
	}

	return &res, err
}

func (u *CategoryProductServiceImpl) Delete(ctx context.Context, id string, userId string) error {
	err := u.categoryProductRepo.Delete(ctx, id, userId)
	return err
}

func (u *CategoryProductServiceImpl) GetList(ctx context.Context, req erpdto.GetListCatProRequest) (*erpdto.CatProductsResponse, error) {
	var err error
	res, err := u.categoryProductRepo.GetList(ctx, req)
	return res, err
}
