package service

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"log"
)

type IDebtService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateDebtRequest) (*models.Debt, error)
	Update(tx *repository.TX, ctx context.Context, debt *models.Debt) error
	Delete(ctx context.Context, id string) error
	GetDebtByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Debt, error)
}

type debtService struct {
	debtRepo repository.DebtRepo
}

func NewDebtService(debtRepo repository.DebtRepo) IDebtService {
	return &debtService{
		debtRepo: debtRepo,
	}
}

func (s *debtService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateDebtRequest) (*models.Debt, error) {
	debt := &models.Debt{}

	if err := copier.Copy(&debt, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	if err := s.debtRepo.Create(tx, ctx, debt); err != nil {
		return nil, err
	}

	return debt, nil
}

func (s *debtService) GetDebtByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Debt, error) {
	return s.debtRepo.GetDebtByOrderId(tx, ctx, orderId)
}

func (s *debtService) Update(tx *repository.TX, ctx context.Context, debt *models.Debt) error {
	return s.debtRepo.UpdateById(tx, ctx, debt)
}

func (s *debtService) Delete(ctx context.Context, id string) error {
	return s.debtRepo.Delete(ctx, id)
}
