package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ERPOrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
}

type erpOrderRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewERPOrderRepository(db *infrastructure.Database, logger *zap.Logger) ERPOrderRepository {
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
		return errors.Wrap(err, "create order failed")
	}

	return nil
}
