package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	models "erp/models"
	"erp/utils"
	"github.com/pkg/errors"
)

type ERPProductRepository interface {
	Create(ctx context.Context, ERPProduct *models.Product) (err error)
	Update(ctx context.Context, product *models.Product) (err error)
	Delete(ctx context.Context, id string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.Product, err error)
	GetList(ctx context.Context, product erpdto.GetListProductRequest) (res []*models.Product, total *int64, err error)
}

type productRepo struct {
	db *infrastructure.Database
}

func NewERPProductRepository(db *infrastructure.Database) ERPProductRepository {
	return &productRepo{db}
}

func (u *productRepo) Create(ctx context.Context, product *models.Product) (err error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return err
	}
	product.UpdaterID = currentUID

	err = u.db.Create(&product).Error
	return errors.Wrap(err, "create product failed")
}

func (u *productRepo) Update(ctx context.Context, product *models.Product) (err error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return err
	}
	product.UpdaterID = currentUID

	err = u.db.Updates(&product).Error
	return errors.Wrap(err, "update product failed")
}

func (u *productRepo) Delete(ctx context.Context, id string) (err error) {
	if err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return errors.Wrap(err, "Delete product failed")
	}
	return nil
}

func (u *productRepo) GetOneByID(ctx context.Context, id string) (res *models.Product, err error) {
	err = u.db.Where("id = ?", id).First(&res).Error
	return res, errors.Wrap(err, "get product by id failed")
}

func (u *productRepo) GetList(ctx context.Context, req erpdto.GetListProductRequest) (res []*models.Product, total *int64, err error) {
	query := u.db.Model(&models.Product{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res).Count(total).Error(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return res, total, err
}
