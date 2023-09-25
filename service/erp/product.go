package erpservice

import (
	"context"
	config "erp/config"
	erpdto "erp/dto/erp"
	models "erp/models"
	repository "erp/repository"
	"erp/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	ERPProductService interface {
		Create(ctx context.Context, req erpdto.CreateProductRequest) (*models.Product, error)
		Update(ctx context.Context, req erpdto.UpdateProductRequest) (*models.Product, error)
		Delete(ctx context.Context, id string) error
		GetOne(ctx context.Context, id string) (*models.Product, error)
		GetList(ctx context.Context, req erpdto.GetListProductRequest) ([]*models.Product, *int64, error)
	}
	productService struct {
		productRepo repository.ERPProductRepository
		config      *config.Config
	}
)

func NewERPProductService(productRepo repository.ERPProductRepository, config *config.Config) ERPProductService {
	return &productService{
		productRepo: productRepo,
		config:      config,
	}
}

func (u *productService) Create(ctx context.Context, req erpdto.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{}
	var err error

	if err = utils.Copy(product, req); err != nil {
		return nil, errors.WithStack(err)
	}
	product.UpdaterID = uuid.FromStringOrNil(req.UserId)

	if err = u.productRepo.Create(ctx, product); err != nil {
		return nil, err
	}
	return product, err
}

func (u *productService) Update(ctx context.Context, req erpdto.UpdateProductRequest) (*models.Product, error) {
	product := &models.Product{}
	var err error

	if err = utils.Copy(product, req); err != nil {
		return nil, err
	}
	product.UpdaterID = uuid.FromStringOrNil(req.UserId)

	if err = u.productRepo.Update(ctx, product); err != nil {
		return nil, err
	}

	return product, err
}

func (u *productService) Delete(ctx context.Context, id string) error {
	return u.productRepo.Delete(ctx, id)
}

func (u *productService) GetOne(ctx context.Context, id string) (*models.Product, error) {
	return u.productRepo.GetOneByID(ctx, id)
}

func (u *productService) GetList(ctx context.Context, req erpdto.GetListProductRequest) ([]*models.Product, *int64, error) {
	return u.productRepo.GetList(ctx, req)
}
