package repository

import (
	"context"
	"erp/cmd/infrastructure"
	models "erp/domain"
	"erp/handler/dto/erp"
	"erp/utils"
	"github.com/pkg/errors"
)

type ERPProductRepository interface {
	Create(ctx context.Context, ERPProduct *models.Product) (err error)
	Update(ctx context.Context, tx *TX, product *models.Product) (err error)
	UpdateMulti(ctx context.Context, product []*models.Product) (err error)
	Delete(ctx context.Context, id string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.Product, err error)
	GetList(ctx context.Context, req erpdto.GetListProductRequest) (res []*models.Product, total int64, err error)
	GetListProductById(ctx context.Context, productIds []string) (res []*models.Product, err error)
}

type productRepo struct {
	db *infrastructure.Database
}

func NewERPProductRepository(db *infrastructure.Database) ERPProductRepository {
	return &productRepo{db}
}

func (r *productRepo) Create(ctx context.Context, product *models.Product) (err error) {
	err = r.db.Create(&product).Error
	return errors.Wrap(err, "create product failed")
}

func (r *productRepo) Update(ctx context.Context, tx *TX, product *models.Product) (err error) {
	tx = GetTX(tx, *r.db)
	err = tx.db.Save(&product).Error
	return errors.Wrap(err, "update product failed")
}

func (r *productRepo) UpdateMulti(ctx context.Context, product []*models.Product) (err error) {
	err = r.db.Save(&product).Error
	return errors.Wrap(err, "update product failed")
}

func (r *productRepo) Delete(ctx context.Context, id string) (err error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return errors.Wrap(err, "Delete product failed")
	}
	return nil
}

func (r *productRepo) GetOneByID(ctx context.Context, id string) (res *models.Product, err error) {
	err = r.db.Where("id = ?", id).First(&res).Error
	return res, err
}

func (r *productRepo) GetList(ctx context.Context, req erpdto.GetListProductRequest) (res []*models.Product, total int64, err error) {
	query := r.db.Model(&models.Product{})
	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res).Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *productRepo) GetListProductById(ctx context.Context, productIds []string) (res []*models.Product, err error) {
	if err = r.db.Model(&models.Product{}).Where("id in (?)", productIds).Find(&res).Error; err != nil {
		return nil, errors.Wrap(err, "get product by id failed")
	}
	return res, nil
}
