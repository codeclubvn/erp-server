package repository

import (
	"context"
	"erp/api_errors"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"
	"go.uber.org/zap"

	"github.com/pkg/errors"
)

type OrderRepo interface {
	Create(tx *TX, ctx context.Context, order *models.Order) error
	Update(tx *TX, ctx context.Context, order *models.Order) error
	GetOneById(ctx context.Context, id string) (*models.Order, error)
	GetList(ctx context.Context, req erpdto.GetListOrderRequest) (res []*models.Order, total int64, err error)
}

type orderRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewOrderRepository(db *infrastructure.Database, logger *zap.Logger) OrderRepo {
	return &orderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *orderRepo) Create(tx *TX, ctx context.Context, order *models.Order) error {
	return tx.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepo) Update(tx *TX, ctx context.Context, order *models.Order) error {
	return tx.db.WithContext(ctx).Save(order).Error
}

func (r *orderRepo) GetOneById(ctx context.Context, id string) (*models.Order, error) {
	var res models.Order
	if err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("OrderItems").
		Preload("Customers").
		First(&res).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, "Find order failed")
	}

	return &res, nil
}

func (r *orderRepo) GetList(ctx context.Context, req erpdto.GetListOrderRequest) (res []*models.Order, total int64, err error) {
	query := r.db.Model(&models.Order{})
	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	query = query.Preload("OrderItems").Preload("Customers")

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}
