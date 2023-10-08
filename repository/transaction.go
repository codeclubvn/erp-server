package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"go.uber.org/zap"
)

type TransactionRepo interface {
	Create(tx *TX, ctx context.Context, trans *models.Transaction) error
	Update(tx *TX, ctx context.Context, trans *models.Transaction) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetTransactionByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Transaction, error)
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

func (r *transactionRepo) Create(tx *TX, ctx context.Context, trans *models.Transaction) error {
	return tx.db.WithContext(ctx).Create(trans).Error
}

func (r *transactionRepo) GetTransactionByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Transaction, error) {
	trans := &models.Transaction{}
	if err := tx.db.WithContext(ctx).Where("order_id = ?", orderId).First(trans).Error; err != nil {
		return nil, err
	}

	return trans, nil
}

func (r *transactionRepo) Update(tx *TX, ctx context.Context, trans *models.Transaction) error {
	return tx.db.WithContext(ctx).Where("id = ?", trans.ID).Updates(trans).Error
}

func (r *transactionRepo) Delete(tx *TX, ctx context.Context, id string) error {
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Transaction{}).Error
}
