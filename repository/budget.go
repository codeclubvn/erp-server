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

type BudgetRepository interface {
	Create(tx *TX, ctx context.Context, input *models.Budget) error
	Update(tx *TX, ctx context.Context, input *models.Budget) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetOneById(ctx context.Context, id string) (*models.Budget, error)
	GetList(ctx context.Context, req erpdto.ListBudgetRequest) (res []*models.Budget, total int64, err error)
}

type budgetRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewBudgetRepository(db *infrastructure.Database, logger *zap.Logger) BudgetRepository {
	return &budgetRepo{
		db:     db,
		logger: logger,
	}
}

func (r *budgetRepo) Create(tx *TX, ctx context.Context, input *models.Budget) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(input).Error
}

func (r *budgetRepo) GetOneById(ctx context.Context, id string) (*models.Budget, error) {
	output := &models.Budget{}
	err := r.db.WithContext(ctx).
		Select("budgets.*, sum(transactions.amount) as spent").
		Joins(`left join transactions on transactions.transaction_category_id = budgets.transaction_category_id 
			AND (transactions.date_time BETWEEN budgets.start_time AND budgets.end_time OR budgets.start_time IS NULL OR budgets.end_time IS NULL)`).
		Where("budgets.id = ?", id).
		Preload("TransactionCategory").
		Group("budgets.id").
		First(output).Error
	return output, err
}

func (r *budgetRepo) GetList(ctx context.Context, req erpdto.ListBudgetRequest) (res []*models.Budget, total int64, err error) {
	query := r.db.Model(&models.Budget{}).Debug().
		Select("budgets.*, sum(transactions.amount) as spent").
		Joins(`left join transactions on transactions.transaction_category_id = budgets.transaction_category_id 
			AND (transactions.date_time BETWEEN budgets.start_time AND budgets.end_time OR budgets.start_time IS NULL OR budgets.end_time IS NULL)`)
	if req.Search != "" {
		query = query.Where("transactions.name ilike ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	query = query.Preload("TransactionCategory").Group("budgets.id")

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *budgetRepo) Update(tx *TX, ctx context.Context, input *models.Budget) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", input.ID).Save(input).Error
}

func (r *budgetRepo) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Budget{}).Error
}
