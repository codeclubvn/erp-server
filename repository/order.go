package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type OrderRepo interface {
	Create(ctx context.Context, order *models.Order) error
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

func (p *erpOrderRepository) Create(ctx context.Context, order *models.Order) error {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return err
	}

	order.UpdaterID = currentUID

	if err := p.db.WithContext(ctx).Create(order).Error; err != nil {
		return errors.Wrap(err, "CreateFlow order failed")
	}

	return nil
}
