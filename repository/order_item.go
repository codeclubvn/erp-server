package repository

import (
	"context"
	"erp/domain"
	"erp/infrastructure"
	"go.uber.org/zap"
)

type OrderItemRepo interface {
	CreateMultiple(tx *TX, ctx context.Context, orderItems []*domain.OrderItem) error
	GetOrderItemByOrderId(ctx context.Context, orderId string) ([]*domain.OrderItem, error)
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

func (p *orderItemRepo) CreateMultiple(tx *TX, ctx context.Context, orderItems []*domain.OrderItem) error {
	if err := tx.db.Create(orderItems).Error; err != nil {
		return err
	}
	return nil
}

func (r *orderItemRepo) GetOrderItemByOrderId(ctx context.Context, orderId string) ([]*domain.OrderItem, error) {
	var res []*domain.OrderItem
	if err := r.db.Where("order_id = ?", orderId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
