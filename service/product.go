package service

import (
	"context"
	config "erp/config"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	models "erp/models"
	repository "erp/repository"
	"erp/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type (
	IProductService interface {
		Create(ctx context.Context, req erpdto.CreateProductRequest) (*models.Product, error)
		Update(ctx context.Context, req erpdto.UpdateProductRequest) (*models.Product, error)
		UpdateMulti(tx *repository.TX, ctx context.Context, req []*models.Product) error
		Delete(ctx context.Context, id string) error
		GetOne(ctx context.Context, id string) (*models.Product, error)
		GetList(ctx context.Context, req erpdto.GetListProductRequest) ([]*models.Product, int64, error)
		GetListProductById(ctx context.Context, productIds []string) ([]*models.Product, error)
	}
	productService struct {
		productRepo repository.ERPProductRepository
		config      *config.Config
		db          *infrastructure.Database
	}
)

func NewProductService(productRepo repository.ERPProductRepository, config *config.Config, db *infrastructure.Database) IProductService {
	return &productService{
		productRepo: productRepo,
		config:      config,
		db:          db,
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

	err = repository.WithTransaction(u.db, func(tx *repository.TX) error {
		if err = u.productRepo.Update(ctx, tx, product); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return product, err
}

func (u *productService) UpdateMulti(tx *repository.TX, ctx context.Context, req []*models.Product) error {
	for _, product := range req {
		if err := u.productRepo.Update(ctx, tx, product); err != nil {
			return err
		}
	}
	return nil
}

func (u *productService) Delete(ctx context.Context, id string) error {
	return u.productRepo.Delete(ctx, id)
}

func (u *productService) GetOne(ctx context.Context, id string) (*models.Product, error) {
	return u.productRepo.GetOneByID(ctx, id)
}

func (u *productService) GetList(ctx context.Context, req erpdto.GetListProductRequest) ([]*models.Product, int64, error) {
	return u.productRepo.GetList(ctx, req)
}

func (u *productService) GetListProductById(ctx context.Context, productIds []string) ([]*models.Product, error) {
	return u.productRepo.GetListProductById(ctx, productIds)
}
