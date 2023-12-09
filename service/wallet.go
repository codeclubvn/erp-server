package service

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type TransactionService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error)
	Update(ctx context.Context, req erpdto.UpdateTransactionRequest) (*models.Transaction, error)
	GetList(ctx context.Context, req erpdto.ListTransactionRequest) ([]*models.Transaction, int64, error)
	Delete(ctx context.Context, transactionID string) error
	GetTransactionByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error)
	GetOne(ctx context.Context, id string) (*models.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	db              *infrastructure.Database
	logger      *zap.Logger
}

func NewTransactionService(transactionRepo repository.TransactionRepository, db *infrastructure.Database, logger *zap.Logger) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		db:              db,
		logger:          logger,
	}
}

func (p *transactionService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionRequest) (*models.Transaction, error) {
	output := &models.Transaction{}
	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.transactionRepo.Create(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *transactionService) Update(ctx context.Context, req erpdto.UpdateTransactionRequest) (*models.Transaction, error) {
	output, err := p.transactionRepo.GetOneById(ctx, req.Id.String())
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.transactionRepo.Update(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *transactionService) GetList(ctx context.Context, req erpdto.ListTransactionRequest) ([]*models.Transaction, int64, error) {
	return p.transactionRepo.GetList(ctx, req)
}

func (p *transactionService) GetOne(ctx context.Context, id string) (*models.Transaction, error) {
	return p.transactionRepo.GetOneById(ctx, id)
}

func (p *transactionService) GetTransactionByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error) {
	return p.transactionRepo.GetTransactionByOrderId(tx, ctx, orderId)
}

func (p *transactionService) Delete(ctx context.Context, transactionID string) error {
	return p.transactionRepo.Delete(nil, ctx, transactionID)
}
