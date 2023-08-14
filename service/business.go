package service

import (
	"context"
	"erp-server/common"
	"erp-server/model"
	"erp-server/repo"
)

type Business struct {
	repo repo.IRepo
}

type IBusiness interface {
	CreateBusiness(ctx context.Context, businessReq model.BusinessRequest) (businessRes model.Business, err error)
	UpdateBusiness(ctx context.Context, businessReq model.BusinessRequest) (businessRes model.Business, err error)
	GetBusiness(ctx context.Context, userId string) (businessRes model.Business, err error)
}

func NewBusiness(repo repo.IRepo) *Business {
	return &Business{
		repo: repo,
	}
}

func (s *Business) CreateBusiness(ctx context.Context, businessReq model.BusinessRequest) (businessRes model.Business, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.CreateBusiness(ctx, &businessRes); err != nil {
		return model.Business{}, err
	}
	return businessRes, nil
}

func (s *Business) UpdateBusiness(ctx context.Context, businessReq model.BusinessRequest) (businessRes model.Business, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.UpdateBusiness(ctx, &businessRes); err != nil {
		return model.Business{}, err
	}
	return businessRes, nil
}

func (s *Business) GetBusiness(ctx context.Context, userId string) (businessRes model.Business, err error) {
	businessRes, err = s.repo.GetBusiness(ctx, userId)
	if err != nil {
		return model.Business{}, err
	}
	return businessRes, nil
}
