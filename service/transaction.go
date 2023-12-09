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

type RevenueService interface {
	Create(tx *repository.TX, ctx context.Context, req erpdto.CreateRevenueRequest) (*models.Transaction, error)
	Update(ctx context.Context, req erpdto.UpdateRevenueRequest) (*models.Transaction, error)
	GetList(ctx context.Context, req erpdto.ListRevenueRequest) ([]*models.Transaction, int64, error)
	Delete(ctx context.Context, revenueID string) error
	GetRevenueByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error)
	GetOne(ctx context.Context, id string) (*models.Transaction, error)
}

type revenueService struct {
	revenueRepo repository.RevenueRepository
	db          *infrastructure.Database
	logger      *zap.Logger
}

func NewRevenueService(revenueRepo repository.RevenueRepository, db *infrastructure.Database, logger *zap.Logger) RevenueService {
	return &revenueService{
		revenueRepo: revenueRepo,
		db:          db,
		logger:      logger,
	}
}

func (p *revenueService) Create(tx *repository.TX, ctx context.Context, req erpdto.CreateRevenueRequest) (*models.Transaction, error) {
	output := &models.Transaction{}
	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.revenueRepo.Create(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *revenueService) Update(ctx context.Context, req erpdto.UpdateRevenueRequest) (*models.Transaction, error) {
	output, err := p.revenueRepo.GetOneById(ctx, req.Id.String())
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(&output, &req); err != nil {
		return nil, err
	}

	if err := p.revenueRepo.Update(nil, ctx, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (p *revenueService) GetList(ctx context.Context, req erpdto.ListRevenueRequest) ([]*models.Transaction, int64, error) {
	return p.revenueRepo.GetList(ctx, req)
}

func (p *revenueService) GetOne(ctx context.Context, id string) (*models.Transaction, error) {
	return p.revenueRepo.GetOneById(ctx, id)
}

func (p *revenueService) GetRevenueByOrderId(tx *repository.TX, ctx context.Context, orderId string) (*models.Transaction, error) {
	return p.revenueRepo.GetRevenueByOrderId(tx, ctx, orderId)
}

func (p *revenueService) Delete(ctx context.Context, revenueID string) error {
	return p.revenueRepo.Delete(nil, ctx, revenueID)
}
