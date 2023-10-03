package erpservice

import (
	"context"
	"erp/api_errors"
	"erp/constants"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"erp/utils"
	"errors"
	"go.uber.org/zap"
)

type OrderService interface {
	CreateFlow(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error)
}

type erpOrderService struct {
	erpOrderRepo       repository.ERPOrderRepository
	db                 *infrastructure.Database
	logger             *zap.Logger
	customerService    ERPCustomerService
	productService     productService
	transactionService transactionService
	debtService        debtService
	orderItemService   orderItemService
}

func NewERPOrderService(
	erpOrderRepo repository.ERPOrderRepository,
	db *infrastructure.Database,
	logger *zap.Logger,
	customerService ERPCustomerService,
	productService productService,
	transactionService transactionService,
	debtService debtService,
	orderItemService orderItemService,
) OrderService {
	return &erpOrderService{
		erpOrderRepo:       erpOrderRepo,
		db:                 db,
		logger:             logger,
		customerService:    customerService,
		productService:     productService,
		transactionService: transactionService,
		debtService:        debtService,
		orderItemService:   orderItemService,
	}
}

func (s *erpOrderService) CreateFlow(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error) {
	// get order code
	orderCode := s.getOrderCode(ctx)

	// if customer_id != "", check customer exist
	if err := s.getCustomer(ctx, req.CustomerID.String()); err != nil {
		return nil, err
	}

	// get list product id, map order item
	productIds, mapOrderItem := s.getProductIdsAndMapOrderItem(ctx, req.OrderItems)

	// get list product
	products, err := s.productService.GetListProductById(ctx, productIds, req.StoreId)
	if err != nil {
		return nil, err
	}

	// validate order item
	if err = s.validateOrderItem(req.OrderItems, products); err != nil {
		return nil, err
	}

	// calculate amount
	calculatedAmount, err := s.CalculateAmount(ctx, products, mapOrderItem)

	// validate calculatedAmount
	if err := s.validateAmount(ctx, req.Amount, calculatedAmount); err != nil {
		return nil, err
	}

	// calculate total = amount + delivery_fee - discount - promote_fee
	total := s.calculateTotalAmount(ctx, calculatedAmount, req)

	// check total and request total
	if err := s.validateTotal(ctx, req.Total, total); err != nil {
		return nil, err
	}

	// check payment if customer_id == "", payment == total
	// check payment if customer_id != "", payment <= total, payment > 0, create debt
	if err := s.handleCustomerPayment(ctx, req, total); err != nil {
		return nil, err
	}

	// create user transaction
	if err := s.createUserTransaction(ctx); err != nil {
		return nil, err
	}

	// create order item
	if err := s.orderItemService.CreateOrderItemFlow(ctx, req.OrderItems); err != nil {
		return nil, err
	}

	// create order
	order, err := s.create(ctx, req, orderCode)
	if err != nil {
		return nil, err
	}

	return order, err
}

func (s *erpOrderService) createUserTransaction(ctx context.Context) error {
	transRequest := erpdto.CreateTransactionRequest{}
	_, err := s.transactionService.Create(ctx, transRequest)
	return err
}

func (s *erpOrderService) getOrderCode(ctx context.Context) string {
	return utils.GenerateCode(constants.NumberOrderCode)
}

func (s *erpOrderService) validateTotal(ctx context.Context, requestTotal, calculatedTotal float64) error {
	if requestTotal != calculatedTotal {
		return errors.New(api_errors.ErrValidation)
	}
	return nil
}

func (s *erpOrderService) validateOrderItem(orderItems []erpdto.OrderItemRequest, products []*models.Product) error {
	if len(orderItems) != len(products) {
		return errors.New(api_errors.ErrValidation)
	}
	return nil
}

func (s *erpOrderService) validateAmount(ctx context.Context, requestAmount, calculatedAmount float64) error {
	if requestAmount != calculatedAmount {
		return errors.New(api_errors.ErrValidation)
	}
	return nil
}

func (s *erpOrderService) create(ctx context.Context, req erpdto.CreateOrderRequest, orderCode string) (*models.Order, error) {
	order := &models.Order{}
	if err := utils.Copy(order, req); err != nil {
		return order, err
	}
	order.Code = orderCode
	err := s.erpOrderRepo.Create(ctx, order)
	return order, err

}

func (s *erpOrderService) getProductIdsAndMapOrderItem(ctx context.Context, orderItems []erpdto.OrderItemRequest) ([]string, map[string]erpdto.OrderItemRequest) {
	productIds := []string{}
	mapOrderItem := map[string]erpdto.OrderItemRequest{}
	for _, value := range orderItems {
		productIds = append(productIds, value.ProductId.String())
		mapOrderItem[value.ProductId.String()] = value
	}

	return productIds, mapOrderItem
}

func (s *erpOrderService) getCustomer(ctx context.Context, customerId string) error {
	if customerId != "" {
		_, err := s.customerService.CustomerDetail(ctx, erpdto.CustomerUriRequest{ID: customerId})
		return err
	}
	return nil
}

func (s *erpOrderService) handleCustomerPayment(ctx context.Context, req erpdto.CreateOrderRequest, total float64) error {
	if req.CustomerID.String() == "" {
		if req.Payment > total {
			return errors.New(api_errors.ErrValidation)
		}
	} else {
		if req.Payment <= 0 || req.Payment > total {
			return errors.New(api_errors.ErrValidation)
		}
		if req.Payment < total {
			// create debt
			debtRequest := erpdto.CreateDebtRequest{}
			_, err := s.debtService.Create(ctx, debtRequest)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *erpOrderService) GetDiscount(ctx context.Context, discountType string, total, discountReq float64) float64 {
	discount := float64(0)
	if discountType == constants.TypePercent {
		discount = total * discountReq
	}
	return discount
}

func (s *erpOrderService) GetPromote(ctx context.Context, promoteType string, total, promoteReq float64) float64 {
	promoteFee := float64(0)
	if promoteType == constants.TypeAmount {
		promoteFee = total * promoteReq
	}
	return promoteFee
}

func (s *erpOrderService) calculateTotalAmount(ctx context.Context, amount float64, req erpdto.CreateOrderRequest) float64 {
	total := amount + req.DeliveryFee

	// get discount
	discount := s.GetDiscount(ctx, req.DiscountType, total, req.Discount)

	// get promote
	promoteFee := s.GetPromote(ctx, req.PromoteType, total, req.PromoteFee)

	total = total - discount - promoteFee
	return total
}

func (s *erpOrderService) CalculateAmount(ctx context.Context, products []*models.Product, mapOrderItem map[string]erpdto.OrderItemRequest) (float64, error) {
	amount := float64(0)
	for _, value := range products {
		if value.Status != constants.ProductStatusActive {
			return 0.0, errors.New(api_errors.ErrValidation)
		}

		if value.Quantity <= mapOrderItem[value.ID.String()].Quantity {
			return 0.0, errors.New(api_errors.ErrValidation)
		}

		if value.Price != mapOrderItem[value.ID.String()].Price {
			return 0.0, errors.New(api_errors.ErrValidation)
		}

		amount += value.Price * value.Quantity
	}
	return amount, nil
}
