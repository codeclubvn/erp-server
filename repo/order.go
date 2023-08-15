package repo

import (
	"context"
	"erp-server/model"
)

func (r *Repo) CreateOrder(ctx context.Context, order *model.Order) error {
	if err := r.db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateOrder(ctx context.Context, order *model.Order) error {
	if err := r.db.Updates(&order).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetOrder(ctx context.Context, oneOrderReq model.OneOrderRequest) (model.Order, error) {
	order := model.Order{}
	if err := r.db.Where("id = ?", oneOrderReq.Id).First(&order).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (r *Repo) GetOrders(ctx context.Context, userId string) (model.Orders, error) {
	orders := model.Orders{}
	if err := r.db.Where("user_id = ?", userId).Find(&orders).Error; err != nil {
		return model.Orders{}, err
	}
	return orders, nil
}
