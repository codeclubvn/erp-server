package erpservice

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"go.uber.org/zap"
)

type OrderItemService interface {
	CreateOrderItemFlow(ctx context.Context, req []erpdto.OrderItemRequest) error
}

type orderItemService struct {
	orderItemRepo repository.OrderItemRepo
	db            *infrastructure.Database
	logger        *zap.Logger
}

func NewOrderItemService(orderItemRepo repository.OrderItemRepo, db *infrastructure.Database, logger *zap.Logger) OrderItemService {
	return &orderItemService{
		orderItemRepo: orderItemRepo,
		db:            db,
		logger:        logger,
	}
}

func (s *orderItemService) CreateOrderItemFlow(ctx context.Context, req []erpdto.OrderItemRequest) error {
	orderItem, err := s.mapOrderItem(ctx, req)
	if err != nil {
		return err
	}
	return s.orderItemRepo.CreateMultiple(ctx, orderItem)
}

func (s *orderItemService) mapOrderItem(ctx context.Context, req []erpdto.OrderItemRequest) ([]*models.OrderItem, error) {
	orderItem := []*models.OrderItem{}
	for _, item := range req {
		orderItem = append(orderItem, &models.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return orderItem, nil
}

func (s *orderItemService) create(ctx context.Context, req []*models.OrderItem) error {
	return s.orderItemRepo.CreateMultiple(ctx, req)
}
