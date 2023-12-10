package repository

import (
	"context"
	erpdto "erp/dto/finance"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type TransactionCategoryRepository interface {
	Create(tx *TX, ctx context.Context, trans *models.CashbookCategory) error
	Update(tx *TX, ctx context.Context, trans *models.CashbookCategory) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetOneById(ctx context.Context, id string) (*models.CashbookCategory, error)
	GetList(ctx context.Context, req erpdto.ListTransactionCategoryRequest) (res []*models.CashbookCategory, total int64, err error)
}

type transactionCategoryRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewTransactionCategoryRepository(db *infrastructure.Database, logger *zap.Logger) TransactionCategoryRepository {
	return &transactionCategoryRepo{
		db:     db,
		logger: logger,
	}
}

func (r *transactionCategoryRepo) Create(tx *TX, ctx context.Context, trans *models.CashbookCategory) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(trans).Error
}

func (r *transactionCategoryRepo) GetOneById(ctx context.Context, id string) (*models.CashbookCategory, error) {
	trans := &models.CashbookCategory{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(trans).Error
	return trans, err
}

func (r *transactionCategoryRepo) GetList(ctx context.Context, req erpdto.ListTransactionCategoryRequest) (res []*models.CashbookCategory, total int64, err error) {
	query := r.db.Model(&models.CashbookCategory{})
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

func (r *transactionCategoryRepo) Update(tx *TX, ctx context.Context, trans *models.CashbookCategory) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", trans.ID).Save(trans).Error
}

func (r *transactionCategoryRepo) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&models.CashbookCategory{}).Error
}
