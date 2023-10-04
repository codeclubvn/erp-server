package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"go.uber.org/zap"
)

type TransactionRepo interface {
	Create(ctx context.Context, trans *models.Transaction) error
}

type transactionRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewTransactionRepository(db *infrastructure.Database, logger *zap.Logger) TransactionRepo {
	return &transactionRepo{
		db:     db,
		logger: logger,
	}
}

func (p *transactionRepo) Create(ctx context.Context, trans *models.Transaction) error {
	return p.db.WithContext(ctx).Create(trans).Error
}
