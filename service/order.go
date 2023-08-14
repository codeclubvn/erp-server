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
	CreateOrder(ctx context.Context, businessReq model.OrderRequest) (businessRes model.Order, err error)
	UpdateOrder(ctx context.Context, businessReq model.OrderRequest) (businessRes model.Order, err error)
	GetOrder(ctx context.Context, userId string) (businessRes model.Order, err error)
	GetOrders(ctx context.Context, userId string) (businessRes model.Orders, err error)
}

func NewOrder(repo repo.IRepo) *Order {
	return &Order{
		repo: repo,
	}
}

func (s *Order) CreateOrder(ctx context.Context, businessReq model.OrderRequest) (businessRes model.Order, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.CreateOrder(ctx, &businessRes); err != nil {
		return model.Order{}, err
	}
	return businessRes, nil
}

func (s *Order) UpdateOrder(ctx context.Context, businessReq model.OrderRequest) (businessRes model.Order, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.UpdateOrder(ctx, &businessRes); err != nil {
		return model.Order{}, err
	}
	return businessRes, nil
}

func (s *Order) GetOrder(ctx context.Context, userId string) (businessRes model.Order, err error) {
	businessRes, err = s.repo.GetOrder(ctx, userId)
	if err != nil {
		return model.Order{}, err
	}
	return businessRes, nil
}

func (s *Order) GetOrders(ctx context.Context, userId string) (businessRes model.Orders, err error) {
	businessRes, err = s.repo.GetOrders(ctx, userId)
	if err != nil {
		return model.Orders{}, err
	}
	return businessRes, nil
}
