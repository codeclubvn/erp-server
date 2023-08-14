package repo

import (
	"context"
	"erp-server/model"
)

func (r *Repo) CreateOrder(ctx context.Context, product *model.Order) error {
	if err := r.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateOrder(ctx context.Context, product *model.Order) error {
	if err := r.db.Updates(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetOrder(ctx context.Context, userId string) (model.Order, error) {
	product := model.Order{}
	if err := r.db.Where("user_id = ?", userId).First(&product).Error; err != nil {
		return model.Order{}, err
	}
	return product, nil
}

func (r *Repo) GetOrders(ctx context.Context, userId string) (model.Orders, error) {
	products := model.Orders{}
	if err := r.db.Where("user_id = ?", userId).Find(&products).Error; err != nil {
		return model.Orders{}, err
	}
	return products, nil
}
