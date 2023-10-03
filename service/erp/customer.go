package erpservice

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"log"
)

type ERPCustomerService interface {
	ListCustomer(ctx context.Context, req erpdto.ListCustomerRequest) ([]*models.Customer, *int64, error)
	CustomerDetail(ctx context.Context, req erpdto.CustomerUriRequest) (*models.Customer, error)
	CreateCustomer(ctx context.Context, req erpdto.CreateCustomerRequest) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, req erpdto.UpdateCustomerRequest) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, customerId string) error
}

type erpCustomerService struct {
	erpCustomerRepo repository.ERPCustomerRepository
	db              *infrastructure.Database
	logger          *zap.Logger
}

func NewERPCustomerService(erpCustomerRepo repository.ERPCustomerRepository, db *infrastructure.Database, logger *zap.Logger) ERPCustomerService {
	return &erpCustomerService{
		erpCustomerRepo: erpCustomerRepo,
		db:              db,
		logger:          logger,
	}
}

func (p *erpCustomerService) ListCustomer(ctx context.Context, req erpdto.ListCustomerRequest) ([]*models.Customer, *int64, error) {
	customers, total, err := p.erpCustomerRepo.List(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return customers, total, nil
}

func (p *erpCustomerService) CustomerDetail(ctx context.Context, req erpdto.CustomerUriRequest) (*models.Customer, error) {
	customers, err := p.erpCustomerRepo.FindOneByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (p *erpCustomerService) CreateCustomer(ctx context.Context, req erpdto.CreateCustomerRequest) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := copier.Copy(&customer, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	customer, err := p.erpCustomerRepo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (p *erpCustomerService) UpdateCustomer(ctx context.Context, req erpdto.UpdateCustomerRequest) (*models.Customer, error) {
	customer := &models.Customer{}

	if err := copier.Copy(&customer, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	_, err := p.erpCustomerRepo.Update(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (p *erpCustomerService) DeleteCustomer(ctx context.Context, customerId string) error {
	err := p.erpCustomerRepo.Delete(ctx, customerId)
	return err
}
