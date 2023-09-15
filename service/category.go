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
	CategoryService interface {
		Create(ctx context.Context, req dto.CategoryRequest) (*dto.CategoryResponse, error)
		Update(ctx context.Context, req dto.CategoryRequest) (*dto.CategoryResponse, error)
		Delete(ctx context.Context, req dto.CategoryRequest) error
		GetOne(ctx context.Context, req dto.CategoryRequest) (*dto.CategoryResponse, error)
		GetList(ctx context.Context, req dto.CategoryRequest) (*dto.CategoriesResponse, error)
	}
	CategoryServiceImpl struct {
		categoryRepo repository.CategoryRepository
		config       *config.Config
	}
)

func NewCategoryService(itemRepo repository.CategoryRepository, config *config.Config) CategoryService {
	return &CategoryServiceImpl{
		categoryRepo: itemRepo,
		config:       config,
	}
}

func (u *CategoryServiceImpl) Create(ctx context.Context, req dto.CategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{}
	var err error

	// todo: check user_id co thuoc cua hang do hay ko

	if err = utils.Copy(req, category); err != nil {
		return nil, err
	}
	if err = u.categoryRepo.Create(ctx, &category); err != nil {
		return nil, err
	}

	res := dto.CategoryResponse{}
	if err = utils.Copy(category, res); err != nil {
		return nil, err
	}

	return &res, err
}

func (u *CategoryServiceImpl) Update(ctx context.Context, req dto.CategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{}
	var err error

	// todo: check user_id co thuoc cua hang do hay ko

	if err = utils.Copy(req, category); err != nil {
		return nil, err
	}
	if err = u.categoryRepo.Update(ctx, &category); err != nil {
		return nil, err
	}

	res := dto.CategoryResponse{}
	if err = utils.Copy(category, res); err != nil {
		return nil, err
	}

	return &res, err
}

func (u *CategoryServiceImpl) Delete(ctx context.Context, req dto.CategoryRequest) error {
	// todo: check user_id co thuoc cua hang do hay ko

	err := u.categoryRepo.Delete(ctx, req.ID)
	return err
}

func (u *CategoryServiceImpl) GetOne(ctx context.Context, req dto.CategoryRequest) (*dto.CategoryResponse, error) {
	category := &models.Category{}
	var err error
	res := dto.CategoryResponse{}

	if req.ID != "" {
		category, err = u.categoryRepo.GetOneByID(ctx, req.ID)
		return &res, err
	}

	if err = utils.Copy(category, res); err != nil {
		return nil, err
	}

	return &res, err
}

// todo:
func (u *CategoryServiceImpl) GetList(ctx context.Context, req dto.CategoryRequest) (*dto.CategoriesResponse, error) {
	res := &dto.CategoriesResponse{}
	var err error
	//r, err := u.categoryRepo.GetList(ctx, category)
	return res, err
}
