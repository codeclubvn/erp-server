package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	models "erp/models"
	"erp/utils"
	"github.com/pkg/errors"
	"time"
)

type ERPProductRepository interface {
	Create(ctx context.Context, ERPProduct *models.Product) (err error)
	Update(ctx context.Context, product *models.Product) (err error)
	Delete(ctx context.Context, id string, userId string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.Product, err error)
	GetList(ctx context.Context, product erpdto.GetListProductRequest) (res *erpdto.ProductsResponse, err error)
}

type productRepositoryImpl struct {
	db *infrastructure.Database
}

func NewERPProductRepository(db *infrastructure.Database) ERPProductRepository {
	if db == nil {
		panic("Database engine is null")
	}
	return &productRepositoryImpl{db}
}

func (u *productRepositoryImpl) Create(ctx context.Context, product *models.Product) (err error) {
	err = u.db.Create(&product).Error
	return errors.Wrap(err, "create product failed")
}

func (u *productRepositoryImpl) Update(ctx context.Context, product *models.Product) (err error) {
	err = u.db.Updates(&product).Error
	return errors.Wrap(err, "update product failed")
}

func (u *productRepositoryImpl) Delete(ctx context.Context, id string, userId string) (err error) {
	err = u.db.Where("id = ?", id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": userId}).Error
	return errors.Wrap(err, "delete product failed")
}

func (u *productRepositoryImpl) GetOneByID(ctx context.Context, id string) (res *models.Product, err error) {
	err = u.db.Where("id = ?", id).First(&res).Error
	return res, errors.Wrap(err, "get product by id failed")
}

func (u *productRepositoryImpl) GetList(ctx context.Context, req erpdto.GetListProductRequest) (res *erpdto.ProductsResponse, err error) {
	var total int64 = 0

	query := u.db.Model(&models.Product{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(u.db, req.PageOptions, &res.Data).Count(&total).Error(); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, err
}
