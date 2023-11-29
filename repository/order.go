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
	GetOrderById(ctx context.Context, id string) (*models.Order, error)
	GetList(ctx context.Context, req erpdto.GetListOrderRequest) ([]*models.Order, error)
}

type erpOrderRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewOrderRepository(db *infrastructure.Database, logger *zap.Logger) OrderRepo {
	return &erpOrderRepository{
		db:     db,
		logger: logger,
	}
}

func (r *erpOrderRepository) Create(tx *TX, ctx context.Context, order *models.Order) error {
	return tx.db.WithContext(ctx).Create(order).Error
}

func (r *erpOrderRepository) Update(tx *TX, ctx context.Context, order *models.Order) error {
	return tx.db.WithContext(ctx).Updates(order).Error
}

func (r *erpOrderRepository) GetOrderById(ctx context.Context, id string) (*models.Order, error) {
	var res models.Order
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&res).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, "Find order failed")
	}

	return &res, nil
}

func (r *erpOrderRepository) GetList(ctx context.Context, req erpdto.GetListOrderRequest) ([]*models.Order, error) {
	var res []*models.Order
	var count int64
	query := r.db.WithContext(ctx).Model(&models.Order{})

	if err := query.Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "Count order failed")
	}

	if err := query.Offset(req.Page).Limit(req.Limit).Find(&res).Error; err != nil {
		return nil, errors.Wrap(err, "Find order failed")
	}

	return res, nil
}
