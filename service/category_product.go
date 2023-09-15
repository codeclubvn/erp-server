package service

import (
	"context"
	config "erp/config"
	"erp/dto"
	models "erp/models"
	repository "erp/repository"
	"erp/utils"
)

type (
	CategoryProductService interface {
		Create(ctx context.Context, req dto.CategoryProductRequest) (*dto.CategoryResponse, error)
		Update(ctx context.Context, req dto.CategoryProductRequest) (*dto.CategoryResponse, error)
		Delete(ctx context.Context, req dto.DeleteCatagoryProductRequest) error
		GetList(ctx context.Context, req dto.CategoryProductRequest) (*dto.CategoriesResponse, error)
	}
	CategoryProductServiceImpl struct {
		categoryProductRepo repository.CategoryProductRepository
		config              *config.Config
	}
)

func NewCategoryProductService(itemRepo repository.CategoryProductRepository, config *config.Config) CategoryProductService {
	return &CategoryProductServiceImpl{
		categoryProductRepo: itemRepo,
		config:              config,
	}
}

func (u *CategoryProductServiceImpl) Create(ctx context.Context, req dto.CategoryProductRequest) (*dto.CategoryResponse, error) {
	category := models.CategoryProduct{}
	var err error

	// todo: check user_id co thuoc cua hang do hay ko

	if err = utils.Copy(req, category); err != nil {
		return nil, err
	}
	if err = u.categoryProductRepo.Create(ctx, &category); err != nil {
		return nil, err
	}

	res := dto.CategoryResponse{}
	if err = utils.Copy(category, res); err != nil {
		return nil, err
	}

	return &res, err
}

func (u *CategoryProductServiceImpl) Update(ctx context.Context, req dto.CategoryProductRequest) (*dto.CategoryResponse, error) {
	category := models.CategoryProduct{}
	var err error

	// todo: check user_id co thuoc cua hang do hay ko

	if err = utils.Copy(req, category); err != nil {
		return nil, err
	}
	if err = u.categoryProductRepo.Update(ctx, &category); err != nil {
		return nil, err
	}

	res := dto.CategoryResponse{}
	if err = utils.Copy(category, res); err != nil {
		return nil, err
	}

	return &res, err
}

func (u *CategoryProductServiceImpl) Delete(ctx context.Context, req dto.DeleteCatagoryProductRequest) error {
	// todo: check user_id co thuoc cua hang do hay ko

	err := u.categoryProductRepo.Delete(ctx, req.ID)
	return err
}

// todo:
func (u *CategoryProductServiceImpl) GetList(ctx context.Context, req dto.CategoryProductRequest) (*dto.CategoriesResponse, error) {
	res := &dto.CategoriesResponse{}
	var err error
	//r, err := u.categoryProductRepo.GetList(ctx, category)
	return res, err
}
