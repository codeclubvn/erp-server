package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"go.uber.org/zap"
)

type OrderItemRepo interface {
	CreateMultiple(ctx context.Context, orderItems []*models.OrderItem) error
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

func (p *orderItemRepo) CreateMultiple(ctx context.Context, orderItems []*models.OrderItem) error {
	if err := p.db.Create(orderItems).Error; err != nil {
		return err
	}
	return nil
}
