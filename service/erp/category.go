package service

import (
	"context"
	config "erp/config"
	"erp/dto/erp"
	models "erp/models"
	repository "erp/repository"
	"erp/utils"
	"github.com/pkg/errors"
)

type (
	ERPCategoryService interface {
		Create(ctx context.Context, req erpdto.CreateCategoryRequest) (*models.Category, error)
		Update(ctx context.Context, req erpdto.UpdateCategoryRequest) (*models.Category, error)
		Delete(ctx context.Context, id string, userId string) error
		GetOne(ctx context.Context, id string) (*models.Category, error)
		GetList(ctx context.Context, req erpdto.GetListCategoryRequest) (*erpdto.CategoriesResponse, error)
	}
	categoryService struct {
		categoryRepo repository.CategoryRepository
		config       *config.Config
	}
)

func NewCategoryService(itemRepo repository.CategoryRepository, config *config.Config) ERPCategoryService {
	return &categoryService{
		categoryRepo: itemRepo,
		config:       config,
	}
}

func (u *categoryService) Create(ctx context.Context, req erpdto.CreateCategoryRequest) (*models.Category, error) {
	category := &models.Category{}
	var err error

	if err = utils.Copy(category, req); err != nil {
		return nil, errors.WithStack(err)
	}
	if err = u.categoryRepo.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, err
}

func (u *categoryService) Update(ctx context.Context, req erpdto.UpdateCategoryRequest) (*models.Category, error) {
	category := &models.Category{}
	var err error

	if err = utils.Copy(category, req); err != nil {
		return nil, err
	}
	if err = u.categoryRepo.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, err
}

func (u *categoryService) Delete(ctx context.Context, id string, userId string) error {
	err := u.categoryRepo.Delete(ctx, id, userId)
	return err
}

func (u *categoryService) GetOne(ctx context.Context, id string) (*models.Category, error) {
	category := &models.Category{}
	var err error

	category, err = u.categoryRepo.GetOneByID(ctx, id)
	return category, err
}

func (u *categoryService) GetList(ctx context.Context, req erpdto.GetListCategoryRequest) (*erpdto.CategoriesResponse, error) {
	res := &erpdto.CategoriesResponse{}
	var err error

	res, err = u.categoryRepo.GetList(ctx, req)
	return res, err
}
