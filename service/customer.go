package service

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/handler/dto/erp"
	"erp/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"log"
)

type ERPCustomerService interface {
	ListCustomer(ctx context.Context, req erpdto.ListCustomerRequest) ([]*domain.Customer, *int64, error)
	GetOneById(ctx context.Context, id string) (*domain.CustomerDetailResponse, error)
	CreateCustomer(ctx context.Context, req erpdto.CreateCustomerRequest) (*domain.Customer, error)
	UpdateCustomer(ctx context.Context, req erpdto.UpdateCustomerRequest) (*domain.Customer, error)
	DeleteCustomer(ctx context.Context, customerId string) error
}

type erpCustomerService struct {
	erpCustomerRepo repository.ERPCustomerRepository
	orderRepo       repository.OrderRepo
	cashbookRepo    repository.CashbookRepository
	db              *infrastructure.Database
	logger          *zap.Logger
}

func NewCustomerService(erpCustomerRepo repository.ERPCustomerRepository, db *infrastructure.Database, logger *zap.Logger, orderRepo repository.OrderRepo, cashbookRepo repository.CashbookRepository) ERPCustomerService {
	return &erpCustomerService{
		erpCustomerRepo: erpCustomerRepo,
		orderRepo:       orderRepo,
		cashbookRepo:    cashbookRepo,
		db:              db,
		logger:          logger,
	}
}

func (s *erpCustomerService) ListCustomer(ctx context.Context, req erpdto.ListCustomerRequest) ([]*domain.Customer, *int64, error) {
	customers, total, err := s.erpCustomerRepo.List(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return customers, total, nil
}

func (s *erpCustomerService) GetOneById(ctx context.Context, id string) (*domain.CustomerDetailResponse, error) {
	customer, err := s.erpCustomerRepo.FindOneByID(ctx, id)
	if err != nil {
		return nil, err
	}

	orderDetail, err := s.orderRepo.GetDetailCustomer(ctx, customer.ID.String())
	if err != nil {
		return nil, err
	}

	totalDebt, err := s.cashbookRepo.GetTotalDebtByCustomerID(ctx, customer.ID)
	if err != nil {
		return nil, err
	}

	output := &domain.CustomerDetailResponse{}

	if err = copier.Copy(&output, &customer); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}
	output.TotalOrder = orderDetail.TotalOrder
	output.TotalDebt = totalDebt
	output.TotalPaid = orderDetail.TotalPaid

	return output, nil
}

func (s *erpCustomerService) CreateCustomer(ctx context.Context, req erpdto.CreateCustomerRequest) (*domain.Customer, error) {
	customer := &domain.Customer{}

	if err := copier.Copy(&customer, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	customer, err := s.erpCustomerRepo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *erpCustomerService) UpdateCustomer(ctx context.Context, req erpdto.UpdateCustomerRequest) (*domain.Customer, error) {
	customer := &domain.Customer{}

	if err := copier.Copy(&customer, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	_, err := s.erpCustomerRepo.Update(ctx, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *erpCustomerService) DeleteCustomer(ctx context.Context, customerId string) error {
	err := s.erpCustomerRepo.Delete(ctx, customerId)
	return err
}
