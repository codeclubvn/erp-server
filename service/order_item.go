package service

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/domain"
	"erp/handler/dto/erp"
	"erp/repository"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type IOrderItemService interface {
	CreateOrderItemFlow(tx *repository.TX, ctx context.Context, req []erpdto.OrderItemRequest, orderId uuid.UUID) error
	GetOrderItemByOrderId(ctx context.Context, orderId string) ([]*domain.OrderItem, error)
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

func (s *orderItemService) CreateOrderItemFlow(tx *repository.TX, ctx context.Context, req []erpdto.OrderItemRequest, orderId uuid.UUID) error {
	orderItem, err := s.mapCreateOrderItem(ctx, req, orderId)
	if err != nil {
		return err
	}
	return s.orderItemRepo.CreateMultiple(tx, ctx, orderItem)
}

func (s *orderItemService) mapCreateOrderItem(ctx context.Context, req []erpdto.OrderItemRequest, orderId uuid.UUID) ([]*domain.OrderItem, error) {
	orderItem := []*domain.OrderItem{}
	for _, item := range req {
		orderItem = append(orderItem, &domain.OrderItem{
			OrderId:   orderId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}
	return orderItem, nil
}

func (s *orderItemService) GetOrderItemByOrderId(ctx context.Context, orderId string) ([]*domain.OrderItem, error) {
	return s.orderItemRepo.GetOrderItemByOrderId(ctx, orderId)
}
