package erpservice

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"erp/utils"

	"go.uber.org/zap"
)

type ERPOrderService interface {
	CreateOrder(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error)
}

type erpOrderService struct {
	erpOrderRepo repository.ERPOrderRepository
	db           *infrastructure.Database
	logger       *zap.Logger
}

func NewERPOrderService(erpOrderRepo repository.ERPOrderRepository, db *infrastructure.Database, logger *zap.Logger) ERPOrderService {
	return &erpOrderService{
		erpOrderRepo: erpOrderRepo,
		db:           db,
		logger:       logger,
	}
}

func (p *erpOrderService) CreateOrder(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error) {
	order := &models.Order{}
	utils.Copy(order, req)
	err := repository.WithTransaction(p.db, func(tx *repository.TX) error {

		if err := p.erpOrderRepo.Create(ctx, order); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return order, nil
}
