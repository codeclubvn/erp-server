package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"go.uber.org/zap"
)

type OrderItemRepo interface {
	CreateMultiple(tx *TX, ctx context.Context, orderItems []*models.OrderItem) error
	GetOrderItemByOrderId(ctx context.Context, orderId string) ([]*models.OrderItem, error)
}

type orderItemRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewOrderItemRepo(db *infrastructure.Database, logger *zap.Logger) OrderItemRepo {
	return &orderItemRepo{
		db:     db,
		logger: logger,
	}
}

func (p *orderItemRepo) CreateMultiple(tx *TX, ctx context.Context, orderItems []*models.OrderItem) error {
	if err := tx.db.Create(orderItems).Error; err != nil {
		return err
	}
	return nil
}

func (r *orderItemRepo) GetOrderItemByOrderId(ctx context.Context, orderId string) ([]*models.OrderItem, error) {
	var res []*models.OrderItem
	if err := r.db.Where("order_id = ?", orderId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
