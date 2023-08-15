package service

import (
	"context"
	"erp-server/common"
	"erp-server/model"
	"erp-server/repo"
)

type Order struct {
	repo repo.IRepo
}

type IOrder interface {
	CreateOrder(ctx context.Context, businessReq model.OrderRequest) (orderRes model.Order, err error)
	UpdateOrder(ctx context.Context, businessReq model.OrderRequest) (orderRes model.Order, err error)
	GetOrder(ctx context.Context, oneOrderReq model.OneOrderRequest) (orderRes model.Order, err error)
	GetOrders(ctx context.Context, userId string) (orderRes model.Orders, err error)
}

func NewOrder(repo repo.IRepo) *Order {
	return &Order{
		repo: repo,
	}
}

func (s *Order) CreateOrder(ctx context.Context, businessReq model.OrderRequest) (orderRes model.Order, err error) {
	common.Sync(businessReq, &orderRes)
	if err = s.repo.CreateOrder(ctx, &orderRes); err != nil {
		return model.Order{}, err
	}
	return orderRes, nil
}

func (s *Order) UpdateOrder(ctx context.Context, businessReq model.OrderRequest) (orderRes model.Order, err error) {
	common.Sync(businessReq, &orderRes)
	if err = s.repo.UpdateOrder(ctx, &orderRes); err != nil {
		return model.Order{}, err
	}
	return orderRes, nil
}

func (s *Order) GetOrder(ctx context.Context, oneOrderReq model.OneOrderRequest) (orderRes model.Order, err error) {
	orderRes, err = s.repo.GetOrder(ctx, oneOrderReq)
	if err != nil {
		return model.Order{}, err
	}
	return orderRes, nil
}

func (s *Order) GetOrders(ctx context.Context, userId string) (orderRes model.Orders, err error) {
	orderRes, err = s.repo.GetOrders(ctx, userId)
	if err != nil {
		return model.Orders{}, err
	}
	return orderRes, nil
}
