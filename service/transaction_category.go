package service

import (
	"context"
	erpdto "erp/api/dto/finance"
	"erp/domain"
	"erp/infrastructure"
	"erp/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type TransactionCategoryService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionCategoryRequest) (*domain.CashbookCategory, error)
	Update(ctx context.Context, req erpdto.UpdateTransactionCategoryRequest) (*domain.CashbookCategory, error)
	GetList(ctx context.Context, req erpdto.ListTransactionCategoryRequest) ([]*domain.CashbookCategory, int64, error)
	Delete(ctx context.Context, transactionCategoryID string) error
	GetOne(ctx context.Context, id string) (*domain.CashbookCategory, error)
}

type transactionCategoryService struct {
	transactionCategoryRepo repository.TransactionCategoryRepository
	db                      *infrastructure.Database
	logger                  *zap.Logger
}

func NewTransactionCategoryService(transactionCategoryRepo repository.TransactionCategoryRepository, db *infrastructure.Database, logger *zap.Logger) TransactionCategoryService {
	return &transactionCategoryService{
		transactionCategoryRepo: transactionCategoryRepo,
		db:                      db,
		logger:                  logger,
	}
}

func (p *transactionCategoryService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateTransactionCategoryRequest) (*domain.CashbookCategory, error) {
	output := &domain.CashbookCategory{}
	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.transactionCategoryRepo.Create(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *transactionCategoryService) Update(ctx context.Context, req erpdto.UpdateTransactionCategoryRequest) (*domain.CashbookCategory, error) {
	output, err := p.transactionCategoryRepo.GetOneById(ctx, req.Id.String())
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.transactionCategoryRepo.Update(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *transactionCategoryService) GetList(ctx context.Context, req erpdto.ListTransactionCategoryRequest) ([]*domain.CashbookCategory, int64, error) {
	return p.transactionCategoryRepo.GetList(ctx, req)
}

func (p *transactionCategoryService) GetOne(ctx context.Context, id string) (*domain.CashbookCategory, error) {
	return p.transactionCategoryRepo.GetOneById(ctx, id)
}

func (p *transactionCategoryService) Delete(ctx context.Context, transactionCategoryID string) error {
	return p.transactionCategoryRepo.Delete(nil, ctx, transactionCategoryID)
}
