package repository

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/handler/dto/erp"
	"erp/utils"
	"erp/utils/api_errors"
	"go.uber.org/zap"

	"github.com/pkg/errors"
)

type OrderRepo interface {
	Create(tx *TX, ctx context.Context, order *domain.Order) error
	Update(tx *TX, ctx context.Context, order *domain.Order) error
	GetOneById(ctx context.Context, id string) (*domain.Order, error)
	GetList(ctx context.Context, req erpdto.GetListOrderRequest) (res []*domain.Order, total int64, err error)
	GetOverview(ctx context.Context, req erpdto.GetListOrderRequest) (res []*domain.OrderOverview, err error)
	GetBestSeller(ctx context.Context, req erpdto.GetListOrderRequest) (res []*domain.ProductBestSellerResponse, err error)
	GetDetailCustomer(ctx context.Context, customerId string) (res *domain.CustomerDetail, err error)
}

type orderRepo struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewOrderRepository(db *infrastructure.Database, logger *zap.Logger) OrderRepo {
	return &orderRepo{
		db:     db,
		logger: logger,
	}
}

func (r *orderRepo) Create(tx *TX, ctx context.Context, order *domain.Order) error {
	return tx.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepo) Update(tx *TX, ctx context.Context, order *domain.Order) error {
	return tx.db.WithContext(ctx).Save(order).Error
}

func (r *orderRepo) GetOneById(ctx context.Context, id string) (*domain.Order, error) {
	var res domain.Order
	if err := r.db.WithContext(ctx).Where("id = ?", id).
		Preload("OrderItems").
		Preload("Customers").
		First(&res).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, "Find order failed")
	}

	return &res, nil
}

func (r *orderRepo) GetList(ctx context.Context, req erpdto.GetListOrderRequest) (res []*domain.Order, total int64, err error) {
	query := r.db.Model(&domain.Order{})
	if req.Search != "" {
		query = query.Where("note ilike ?", "%"+req.Search+"%")
	}

	if req.StartTime != "" && req.EndTime != "" {
		query = query.Where("created_at BETWEEN ? AND ?", req.StartTime, req.EndTime)
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	query = query.Preload("OrderItems").Preload("Customer")

	if err = utils.QueryPagination(query, req.PageOptions, &res).
		Count(&total).Error(); err != nil {
		return nil, 0, errors.WithStack(err)
	}
	return res, total, err
}

func (r *orderRepo) GetOverview(ctx context.Context, req erpdto.GetListOrderRequest) (res []*domain.OrderOverview, err error) {
	queryString := `SELECT count(confirm) as confirm, count(delivery) as delivery, count(complete) as complete, count(cancel) as cancel,
		sum(revenue) as revenue, sum(income) as income
		FROM ( select CASE WHEN status = 'confirm' then 1 else null end as confirm,
		CASE WHEN status = 'delivery' then 1 else null end as delivery,
		CASE WHEN status = 'complete' then 1 else null end as complete,
		CASE WHEN status = 'cancel' then 1 else null end as cancel,
		CASE WHEN status != 'cancel' then total else null end revenue,
		CASE WHEN status != 'cancel' then "cost" else null end income
		FROM orders `

	if req.Search != "" {
		queryString += "WHERE note iLike " + "'%" + req.Search + "%'"
	}

	queryString += `) as t`

	err = r.db.Debug().Raw(queryString).Find(&res).Error
	return res, err
}

func (r *orderRepo) GetBestSeller(ctx context.Context, req erpdto.GetListOrderRequest) (res []*domain.ProductBestSellerResponse, err error) {
	err = r.db.Table("order_items").Select("products.*, sum(order_items.quantity) as quantity_sold").
		Joins("inner join products on order_items.product_id = products.id").Order("quantity_sold desc").
		Where("order_items.status != 'cancel'").
		Limit(10).Group("products.id").Find(&res).Error
	return res, err
}

func (r *orderRepo) GetDetailCustomer(ctx context.Context, customerId string) (res *domain.CustomerDetail, err error) {
	queryString := `SELECT count(id) as total_order, coalesce(sum(payment), 0) as total_paid
	FROM orders WHERE customer_id = ?`

	err = r.db.Debug().Raw(queryString, customerId).Find(&res).Error
	return res, err
}
