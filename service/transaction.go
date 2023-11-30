package service

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"log"
)

type ITransactionService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error)
	Update(tx *repository.TX, ctx context.Context, trans *models.Transaction) error
	Delete(tx *repository.TX, ctx context.Context, id string) error
	GetTransactionByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepo
}

func NewTransactionService(transactionRepo repository.TransactionRepo) ITransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *transactionService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error) {
	transaction := &models.Transaction{}

	if err := copier.Copy(&transaction, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	if err := s.transactionRepo.Create(tx, ctx, transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetTransactionByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error) {
	return s.transactionRepo.GetTransactionByOrderId(tx, ctx, orderId)
}

func (s *transactionService) Update(tx *repository.TX, ctx context.Context, trans *models.Transaction) error {
	return s.transactionRepo.Update(tx, ctx, trans)
}

func (s *transactionService) Delete(tx *repository.TX, ctx context.Context, id string) error {
	return s.transactionRepo.Delete(tx, ctx, id)
}
