package service

import (
	"context"
	"erp-server/common"
	"erp-server/model"
	"erp-server/repo"
)

type Money struct {
	repo repo.IRepo
}

type IMoney interface {
	CreateMoney(ctx context.Context, businessReq model.MoneyRequest) (businessRes model.Money, err error)
	UpdateMoney(ctx context.Context, businessReq model.MoneyRequest) (businessRes model.Money, err error)
	GetMoney(ctx context.Context, userId string) (businessRes model.Money, err error)
	GetMoneys(ctx context.Context, userId string) (businessRes model.Moneys, err error)
	DeleteMoney(ctx context.Context, id string) (err error)
}

func NewMoney(repo repo.IRepo) *Money {
	return &Money{
		repo: repo,
	}
}

func (s *Money) CreateMoney(ctx context.Context, businessReq model.MoneyRequest) (businessRes model.Money, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.CreateMoney(ctx, &businessRes); err != nil {
		return model.Money{}, err
	}
	return businessRes, nil
}

func (s *Money) UpdateMoney(ctx context.Context, businessReq model.MoneyRequest) (businessRes model.Money, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.UpdateMoney(ctx, &businessRes); err != nil {
		return model.Money{}, err
	}
	return businessRes, nil
}

func (s *Money) GetMoney(ctx context.Context, userId string) (businessRes model.Money, err error) {
	businessRes, err = s.repo.GetMoney(ctx, userId)
	if err != nil {
		return model.Money{}, err
	}
	return businessRes, nil
}

func (s *Money) GetMoneys(ctx context.Context, userId string) (businessRes model.Moneys, err error) {
	businessRes, err = s.repo.GetMoneys(ctx, userId)
	if err != nil {
		return model.Moneys{}, err
	}
	return businessRes, nil
}

func (s *Money) DeleteMoney(ctx context.Context, id string) (err error) {
	if err = s.repo.DeleteMoney(ctx, id); err != nil {
		return err
	}
	return nil
}
