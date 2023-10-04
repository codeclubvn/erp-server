package erpservice

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type IOrderItemService interface {
	CreateOrderItemFlow(ctx context.Context, req []erpdto.OrderItemRequest, orderId uuid.UUID) error
}

type orderItemService struct {
	orderItemRepo repository.OrderItemRepo
	db            *infrastructure.Database
	logger        *zap.Logger
}

func NewOrderItemService(orderItemRepo repository.OrderItemRepo, db *infrastructure.Database, logger *zap.Logger) IOrderItemService {
	return &orderItemService{
		orderItemRepo: orderItemRepo,
		db:            db,
		logger:        logger,
	}
}

func (s *orderItemService) CreateOrderItemFlow(ctx context.Context, req []erpdto.OrderItemRequest, orderId uuid.UUID) error {
	orderItem, err := s.mapOrderItem(ctx, req, orderId)
	if err != nil {
		return err
	}
	return s.orderItemRepo.CreateMultiple(ctx, orderItem)
}

func (s *orderItemService) mapOrderItem(ctx context.Context, req []erpdto.OrderItemRequest, orderId uuid.UUID) ([]*models.OrderItem, error) {
	orderItem := []*models.OrderItem{}
	for _, item := range req {
		orderItem = append(orderItem, &models.OrderItem{
			OrderId:   orderId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return orderItem, nil
}

func (s *orderItemService) create(ctx context.Context, req []*models.OrderItem) error {
	return s.orderItemRepo.CreateMultiple(ctx, req)
}
