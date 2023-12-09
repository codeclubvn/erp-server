package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type TransactionRepository interface {
	Create(tx *TX, ctx context.Context, trans *models.Transaction) error
	Update(tx *TX, ctx context.Context, trans *models.Transaction) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetOneById(ctx context.Context, id string) (*models.Transaction, error)
	GetTransactionByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Transaction, error)
	GetList(ctx context.Context, req erpdto.ListTransactionRequest) (res []*models.Transaction, total int64, err error)
}

type transactionRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewTransactionRepository(db *infrastructure.Database, logger *zap.Logger) TransactionRepository {
	return &transactionRepo{
		db:     db,
		logger: logger,
	}
}

func (r *transactionRepo) Create(tx *TX, ctx context.Context, trans *models.Transaction) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(trans).Error
}

func (r *transactionRepo) GetOneById(ctx context.Context, id string) (*models.Transaction, error) {
	trans := &models.Transaction{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(trans).Error
	return trans, err
}

func (r *transactionRepo) GetTransactionByOrderId(tx *TX, ctx context.Context, orderId string) (*models.Transaction, error) {
	tx = GetTX(tx, *r.db)
	trans := &models.Transaction{}
	err := tx.db.WithContext(ctx).Where("order_id = ?", orderId).First(trans).Error
	return trans, err
}

func (r *transactionRepo) GetList(ctx context.Context, req erpdto.ListTransactionRequest) (res []*models.Transaction, total int64, err error) {
	query := r.db.Model(&models.Transaction{})
	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *transactionRepo) Update(tx *TX, ctx context.Context, trans *models.Transaction) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", trans.ID).Save(trans).Error
}

func (r *transactionRepo) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Transaction{}).Error
}
