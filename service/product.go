package service

import (
	"context"
	"erp-server/common"
	"erp-server/model"
	"erp-server/repo"
)

type Product struct {
	repo repo.IRepo
}

type IProduct interface {
	CreateProduct(ctx context.Context, businessReq model.ProductRequest) (businessRes model.Product, err error)
	UpdateProduct(ctx context.Context, businessReq model.ProductRequest) (businessRes model.Product, err error)
	GetProduct(ctx context.Context, userId string) (businessRes model.Product, err error)
	GetProducts(ctx context.Context, userId string) (businessRes model.Products, err error)
}

func NewProduct(repo repo.IRepo) *Product {
	return &Product{
		repo: repo,
	}
}

func (s *Product) CreateProduct(ctx context.Context, businessReq model.ProductRequest) (businessRes model.Product, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.CreateProduct(ctx, &businessRes); err != nil {
		return model.Product{}, err
	}
	return businessRes, nil
}

func (s *Product) UpdateProduct(ctx context.Context, businessReq model.ProductRequest) (businessRes model.Product, err error) {
	common.Sync(businessReq, &businessRes)
	if err = s.repo.UpdateProduct(ctx, &businessRes); err != nil {
		return model.Product{}, err
	}
	return businessRes, nil
}

func (s *Product) GetProduct(ctx context.Context, userId string) (businessRes model.Product, err error) {
	businessRes, err = s.repo.GetProduct(ctx, userId)
	if err != nil {
		return model.Product{}, err
	}
	return businessRes, nil
}

func (s *Product) GetProducts(ctx context.Context, userId string) (businessRes model.Products, err error) {
	businessRes, err = s.repo.GetProducts(ctx, userId)
	if err != nil {
		return model.Products{}, err
	}
	return businessRes, nil
}
