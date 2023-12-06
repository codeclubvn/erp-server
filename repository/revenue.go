package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type RevenueRepository interface {
	Create(tx *TX, ctx context.Context, trans *models.Revenue) error
	Update(tx *TX, ctx context.Context, trans *models.Revenue) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetOneById(ctx context.Context, id string) (*models.Revenue, error)
	GetRevenueByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Revenue, error)
	GetList(ctx context.Context, req erpdto.ListRevenueRequest) (res []*models.Revenue, total int64, err error)
}

type revenueRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewRevenueRepository(db *infrastructure.Database, logger *zap.Logger) RevenueRepository {
	return &revenueRepo{
		db:     db,
		logger: logger,
	}
}

func (r *revenueRepo) Create(tx *TX, ctx context.Context, trans *models.Revenue) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(trans).Error
}

func (r *revenueRepo) GetOneById(ctx context.Context, id string) (*models.Revenue, error) {
	trans := &models.Revenue{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(trans).Error
	return trans, err
}

func (r *revenueRepo) GetRevenueByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Revenue, error) {
	tx = GetTX(tx, *r.db)
	trans := &models.Revenue{}
	err := tx.db.WithContext(ctx).Where("order_id = ?", orderId).First(trans).Error
	return trans, err
}

func (r *revenueRepo) GetList(ctx context.Context, req erpdto.ListRevenueRequest) (res []*models.Revenue, total int64, err error) {
	query := r.db.Model(&models.Order{})
	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *revenueRepo) Update(tx *TX, ctx context.Context, trans *models.Revenue) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", trans.ID).Save(trans).Error
}

func (r *revenueRepo) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Revenue{}).Error
}
