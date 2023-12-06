package repository

import (
	"context"
	"erp/infrastructure"
	"erp/models"
	"go.uber.org/zap"
)

type DebtRepo interface {
	Create(tx *TX, ctx context.Context, debt *models.Debt) error
	UpdateById(tx *TX, ctx context.Context, debt *models.Debt) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetDebtByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Debt, error)
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

func (r *debtRepo) Create(tx *TX, ctx context.Context, debt *models.Debt) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(debt).Error
}

func (r *debtRepo) UpdateById(tx *TX, ctx context.Context, debt *models.Debt) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", debt.ID).Save(debt).Error
}

func (r *debtRepo) GetDebtByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Debt, error) {
	debt := &models.Debt{}
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderId).First(debt).Error; err != nil {
		return nil, err
	}

	return debt, nil
}

func (r *debtRepo) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Debt{}).Error
}
