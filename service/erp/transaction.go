package erpservice

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"log"
)

type ITransactionService interface {
	Create(ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepo
}

func NewTransactionService(transactionRepo repository.TransactionRepo) ITransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *transactionService) Create(ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error) {
	transaction := &models.Transaction{}

	if err := copier.Copy(&transaction, &req); err != nil {
		log.Println("Copy struct failed!")
		return nil, err
	}

	if err := s.transactionRepo.Create(ctx, transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}
