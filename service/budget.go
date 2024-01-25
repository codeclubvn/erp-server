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

type BudgetService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateBudgetRequest) (*domain.Budget, error)
	Update(ctx context.Context, req erpdto.UpdateBudgetRequest) (*domain.Budget, error)
	Delete(ctx context.Context, budgetID string) error
	GetOne(ctx context.Context, id string) (*domain.Budget, error)
	GetList(ctx context.Context, req erpdto.ListBudgetRequest) ([]*domain.Budget, int64, error)
}

type budgetService struct {
	budgetRepo      repository.BudgetRepository
	transactionRepo repository.CashbookRepository
	db              *infrastructure.Database
	logger          *zap.Logger
}

func NewBudgetService(budgetRepo repository.BudgetRepository, db *infrastructure.Database, logger *zap.Logger, transactionRepo repository.CashbookRepository) BudgetService {
	return &budgetService{
		budgetRepo:      budgetRepo,
		db:              db,
		logger:          logger,
		transactionRepo: transactionRepo,
	}
}

func (p *budgetService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateBudgetRequest) (*domain.Budget, error) {
	output := &domain.Budget{}
	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.budgetRepo.Create(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *budgetService) Update(ctx context.Context, req erpdto.UpdateBudgetRequest) (*domain.Budget, error) {
	output, err := p.budgetRepo.GetOneById(ctx, req.Id.String())
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.budgetRepo.Update(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *budgetService) GetList(ctx context.Context, req erpdto.ListBudgetRequest) ([]*domain.Budget, int64, error) {
	return p.budgetRepo.GetList(ctx, req)
}

func (p *budgetService) GetOne(ctx context.Context, id string) (*domain.Budget, error) {
	return p.budgetRepo.GetOneById(ctx, id)
}

func (p *budgetService) Delete(ctx context.Context, budgetID string) error {
	return p.budgetRepo.Delete(nil, ctx, budgetID)
}
