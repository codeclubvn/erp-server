package repository

import (
	"context"
	erpdto "erp/api/dto/finance"
	"erp/domain"
	"erp/infrastructure"
	"erp/utils"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

type CashbookRepository interface {
	Create(tx *TX, ctx context.Context, input *domain.Cashbook) error
	Update(tx *TX, ctx context.Context, input *domain.Cashbook) error
	Delete(tx *TX, ctx context.Context, id string) error
	GetOneById(ctx context.Context, id string) (*domain.Cashbook, error)
	GetCashbookByOrderId(tx *TX, ctx context.Context, orderId string) (*domain.Cashbook, error)
	GetDebtByOrderId(tx *TX, ctx context.Context, orderId string) (*domain.Cashbook, error)
	GetList(ctx context.Context, req erpdto.ListCashbookRequest) (res []*domain.Cashbook, total int64, err error)
	GetListDebt(ctx context.Context, req erpdto.ListCashbookRequest) (res []*domain.Cashbook, total int64, err error)
	GetTotalTransactionByCategoryIdAndTime(ctx context.Context, categoryId uuid.UUID, startTime, endTime *time.Time) (total float64, err error)
	GetListTotalTransactionByCategoryIdAndTime(ctx context.Context, categoryId uuid.UUID, startTime, endTime *time.Time) (output []*erpdto.TotalTransactionByCategoryResponse, err error)
	GetTotalDebtByCustomerID(ctx context.Context, customerId uuid.UUID) (total float64, err error)
}

type transactionRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewTransactionRepository(db *infrastructure.Database, logger *zap.Logger) CashbookRepository {
	return &transactionRepo{
		db:     db,
		logger: logger,
	}
}

func (r *transactionRepo) Create(tx *TX, ctx context.Context, input *domain.Cashbook) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Create(input).Error
}

func (r *transactionRepo) GetOneById(ctx context.Context, id string) (*domain.Cashbook, error) {
	input := &domain.Cashbook{}
	err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("CashbookCategory").
		First(input).Error
	return input, err
}

func (r *transactionRepo) GetCashbookByOrderId(tx *TX, ctx context.Context, orderId string) (*domain.Cashbook, error) {
	tx = GetTX(tx, *r.db)
	input := &domain.Cashbook{}
	err := tx.db.WithContext(ctx).Where("order_id = ?", orderId).
		Joins("left join cashbook_categories on cashbook_categories.id = cashbooks.cashbook_category_id").
		Where("cashbook_categories.type not in (?)", []string{"debt", "loan"}).
		First(input).Error
	return input, err
}

func (r *transactionRepo) GetDebtByOrderId(tx *TX, ctx context.Context, orderId string) (*domain.Cashbook, error) {
	tx = GetTX(tx, *r.db)
	input := &domain.Cashbook{}
	err := tx.db.WithContext(ctx).Where("order_id = ?", orderId).
		Joins("left join cashbook_categories on cashbook_categories.id = cashbooks.cashbook_category_id").
		Where("cashbook_categories.type in (?)", []string{"debt", "loan"}).
		First(input).Error
	return input, err
}

func (r *transactionRepo) GetList(ctx context.Context, req erpdto.ListCashbookRequest) (res []*domain.Cashbook, total int64, err error) {
	query := r.db.Model(&domain.Cashbook{})
	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	if req.Sort == "" {
		query = query.Order(req.Sort)
	} else {
		query = query.Order("date_time desc")
	}

	query = query.Preload("CashbookCategory").Preload("Wallet").Preload("Order")

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *transactionRepo) GetListDebt(ctx context.Context, req erpdto.ListCashbookRequest) (res []*domain.Cashbook, total int64, err error) {
	query := r.db.Model(&domain.Cashbook{}).Debug().
		Joins("left join cashbook_categories on cashbook_categories.id = cashbooks.cashbook_category_id").
		Where("cashbook_categories.type in (?)", []string{"debt", "loan"})
	if req.Search != "" {
		query = query.Where("name ilike ?", "%"+req.Search+"%")
	}

	if req.Sort == "" {
		query = query.Order(req.Sort)
	} else {
		query = query.Order("date_time desc")
	}

	query = query.Preload("CashbookCategory").Preload("Wallet").Preload("Order")

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *transactionRepo) Update(tx *TX, ctx context.Context, input *domain.Cashbook) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", input.ID).Save(input).Error
}

func (r *transactionRepo) Delete(tx *TX, ctx context.Context, id string) error {
	tx = GetTX(tx, *r.db)
	return tx.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Cashbook{}).Error
}

func (r *transactionRepo) GetTotalTransactionByCategoryIdAndTime(ctx context.Context, categoryId uuid.UUID, startTime, endTime *time.Time) (total float64, err error) {
	query := r.db.Model(&domain.Cashbook{}).Where("cashbook_category_id = ?", categoryId)
	if startTime != nil {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", endTime)
	}
	err = query.Select("sum(amount)").Row().Scan(&total)
	return total, err
}

func (r *transactionRepo) GetListTotalTransactionByCategoryIdAndTime(ctx context.Context, categoryId uuid.UUID, startTime, endTime *time.Time) (output []*erpdto.TotalTransactionByCategoryResponse, err error) {
	query := r.db.Table("cashbooks").Select("cashbook_category_id, sum(amount) as total").
		Where("cashbook_category_id = ?", categoryId).
		Group("cashbook_category_id")
	if startTime != nil {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", endTime)
	}
	err = query.Find(&output).Error
	return output, err
}

func (r *transactionRepo) GetTotalDebtByCustomerID(ctx context.Context, customerId uuid.UUID) (total float64, err error) {
	total = float64(0)
	query := r.db.Model(&domain.Cashbook{}).Where("customer_id = ? and is_pay = false", customerId)
	err = query.Pluck("coalesce(sum(amount),0)", &total).Error
	return total, err
}
