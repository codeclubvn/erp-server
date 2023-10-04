package erpservice

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"log"
)

type IDebtService interface {
	Create(ctx context.Context, req erpdto.CreateDebtRequest) (*models.Debt, error)
}

type debtService struct {
	debtRepo repository.DebtRepo
}

func NewDebtService(debtRepo repository.DebtRepo) IDebtService {
	return &debtService{
		debtRepo: debtRepo,
	}
}

func (s *debtService) Create(ctx context.Context, req erpdto.CreateDebtRequest) (*models.Debt, error) {
	debt := &models.Debt{}

	if err := copier.Copy(&debt, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	if err := s.debtRepo.Create(ctx, debt); err != nil {
		return nil, err
	}

	return debt, nil
}
