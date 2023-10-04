package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"go.uber.org/zap"
)

type DebtRepo interface {
	Create(ctx context.Context, trans *models.Debt) error
}

type debtRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewDebtRepo(db *infrastructure.Database, logger *zap.Logger) DebtRepo {
	return &debtRepo{
		db:     db,
		logger: logger,
	}
}

func (p *debtRepo) Create(ctx context.Context, trans *models.Debt) error {
	return p.db.WithContext(ctx).Create(trans).Error
}
