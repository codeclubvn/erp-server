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
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

type OrderService interface {
	CreateFlow(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error)
}

type erpOrderService struct {
	erpOrderRepo       repository.OrderRepo
	db                 *infrastructure.Database
	logger             *zap.Logger
	customerService    ERPCustomerService
	productService     IProductService
	transactionService ITransactionService
	debtService        IDebtService
	orderItemService   IOrderItemService
	promoteService     IPromoteService
}

func NewOrderService(
	erpOrderRepo repository.OrderRepo,
	db *infrastructure.Database,
	logger *zap.Logger,
	customerService ERPCustomerService,
	productService IProductService,
	transactionService ITransactionService,
	debtService IDebtService,
	orderItemService IOrderItemService,
	promoteService IPromoteService,
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
		promoteService:     promoteService,
	}
}

func (s *erpOrderService) CreateFlow(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error) {
	// get order code
	req.Code = s.getOrderCode(ctx)
	req.OrderId = uuid.NewV4()

	// if customer_id != "", check customer exist
	if err := s.getCustomer(ctx, utils.ValidString(req.CustomerId)); err != nil {
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
	if err != nil {
		return nil, err
	}

	// validate calculatedAmount
	if err := s.validateAmount(ctx, req.Amount, calculatedAmount); err != nil {
		return nil, err
	}

	// calculate calcTotal = amount + delivery_fee - discount - promote_fee
	// check discount_type
	// promote flow
	calcTotal, err := s.calculateTotalAmount(ctx, calculatedAmount, req)
	if err != nil {
		return nil, err
	}

	// check calcTotal and request calcTotal
	if err = s.validateTotal(ctx, req.Total, calcTotal); err != nil {
		return nil, err
	}

	// check payment if customer_id == "", payment == calcTotal
	// check payment if customer_id != "", payment <= calcTotal, payment > 0, create debt
	if err := s.handleCustomerPayment(ctx, req, calcTotal); err != nil {
		return nil, err
	}

	// create user transaction
	if err := s.createUserTransaction(ctx, req); err != nil {
		return nil, err
	}

	// create order item
	if err := s.orderItemService.CreateOrderItemFlow(ctx, req.OrderItems, req.OrderId); err != nil {
		return nil, err
	}

	// create order
	order, err := s.create(ctx, req)
	if err != nil {
		return nil, err
	}

	// update product quantity, sold
	if err := s.UpdateProductQuantity(ctx, products, mapOrderItem); err != nil {
		return nil, err
	}

	return order, err
}

func (s *erpOrderService) UpdateProductQuantity(ctx context.Context, products []*models.Product, mapOrderItem map[string]erpdto.OrderItemRequest) error {
	// if quantity is null, only update sold
	// if quantity is not null, update quantity and sold
	for _, value := range products {
		if value.Quantity == nil {
			value.Sold += mapOrderItem[value.ID.String()].Quantity
		} else {
			value.Quantity = utils.IntPointer(utils.ValidInt(value.Quantity) - mapOrderItem[value.ID.String()].Quantity)
			value.Sold += mapOrderItem[value.ID.String()].Quantity
		}
	}
	if err := s.productService.UpdateMulti(ctx, products); err != nil {
		return err
	}
	return nil
}

func (s *erpOrderService) createUserTransaction(ctx context.Context, req erpdto.CreateOrderRequest) error {
	if req.Payment <= 0 {
		return nil
	}

	transRequest := erpdto.CreateTransactionRequest{
		OrderId: req.OrderId,
		Amount:  req.Payment,
		Status:  constants.TransactionStatusIn,
	}

	if req.Payment > req.Total {
		transRequest.Amount = req.Total
	}
	_, err := s.transactionService.Create(ctx, transRequest)
	return err
}

func (s *erpOrderService) getOrderCode(ctx context.Context) string {
	return utils.GenerateCode(constants.NumberOrderCode)
}

func (s *erpOrderService) validateTotal(ctx context.Context, requestTotal, calculatedTotal float64) error {
	if requestTotal != calculatedTotal {
		return errors.New(api_errors.ErrTotalInvalid)
	}
	return nil
}

func (s *erpOrderService) validateOrderItem(orderItems []erpdto.OrderItemRequest, products []*models.Product) error {
	if len(orderItems) != len(products) {
		return errors.New(api_errors.ErrOrderItemInvalid)
	}
	return nil
}

func (s *erpOrderService) validateAmount(ctx context.Context, requestAmount, calculatedAmount float64) error {
	if requestAmount != calculatedAmount {
		return errors.New(api_errors.ErrAmountIsNotMatched)
	}
	return nil
}

func (s *erpOrderService) create(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error) {
	order := &models.Order{}
	if err := utils.Copy(order, req); err != nil {
		return order, err
	}
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
	if req.CustomerId == nil {
		if req.Payment < total || req.Payment > total {
			return errors.New(api_errors.ErrPaymentInvalid)
		}
	} else {
		if req.Payment < 0 || req.Payment > total {
			return errors.New(api_errors.ErrPaymentInvalid)
		}
		if req.Payment < total {
			// create debt
			debtRequest := erpdto.CreateDebtRequest{
				OrderId:    req.OrderId,
				Amount:     total - req.Payment,
				Status:     constants.DebtStatusOut,
				CustomerId: uuid.FromStringOrNil(utils.ValidString(req.CustomerId)),
			}
			_, err := s.debtService.Create(ctx, debtRequest)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *erpOrderService) GetDiscount(ctx context.Context, discountType erpdto.DiscountType, total, discountReq float64) (float64, error) {
	discount := discountReq
	switch discountType {
	case constants.TypePercent:
		if discountReq <= 0 || discountReq > constants.MaxOfPercent {
			return 0, errors.New(api_errors.ErrDiscountPercentInvalid)
		}
		discount = total * discountReq / constants.MaxOfPercent
	case constants.TypeAmount:
		if discountReq <= 0 || discountReq > total {
			return 0, errors.New(api_errors.ErrDiscountAmountInvalid)
		}
	default:
		return 0, nil
	}
	return discount, nil
}

// PromoteFlow
// check promote id exist
// check promote type
// check promote fee
func (s *erpOrderService) PromoteFlow(ctx context.Context, req erpdto.CreateOrderRequest, total float64) (float64, error) {
	// check promote id exist
	if req.PromoteCode == nil {
		return 0, nil
	}

	// check customer_id exist
	if req.PromoteCode != nil && req.CustomerId == nil {
		return 0, errors.New(api_errors.ErrPromoteCodeRequiredCustomer)
	}

	promote, err := s.promoteService.GetPromoteByCode(ctx, utils.ValidString(req.PromoteCode))
	if err != nil {
		return 0, nil
	}

	// check customer_id use promote
	times, err := s.promoteService.CountCustomerUsePromote(ctx, utils.ValidString(req.CustomerId), utils.ValidString(req.PromoteCode))
	if err != nil {
		if !utils.ErrNoRows(err) {
			return 0, err
		}
	}
	if int(times) >= promote.MaxUse {
		return 0, errors.New(api_errors.ErrPromoteCodeMaxUse)
	}

	// check quantity use | is_active
	if (promote.Quantity != nil && *promote.Quantity <= *promote.QuantityUse) || promote.Status == constants.InActive {
		return 0, nil
	}

	// check nil time
	// check today is between start_date and end_date
	if promote.StartDate != nil && promote.EndDate != nil {
		if !utils.IsBetweenDate(utils.ValidTime(promote.StartDate), utils.ValidTime(promote.EndDate), time.Now()) {
			return 0, nil
		}
	}

	// Update quantity use
	quantityUse := utils.ValidInt(promote.QuantityUse) + 1
	if err = s.promoteService.UpdateQuantityUse(ctx, utils.ValidString(req.PromoteCode), quantityUse); err != nil {
		return 0, err
	}

	// Update promote_use
	if err = s.promoteService.CreatePromoteUse(ctx, erpdto.CreatePromoteUseRequest{
		CustomerId:  utils.ValidString(req.CustomerId),
		PromoteCode: utils.ValidString(req.PromoteCode),
	}); err != nil {
		return 0, nil
	}

	promoteFee := promote.DiscountValue
	if promote.PromoteType == constants.TypePercent {
		promoteFee = total * promote.DiscountValue / constants.MaxOfPercent
	}

	// check max_amount
	if promote.MaxAmount != nil && promoteFee > utils.ValidFloat64(promote.MaxAmount) {
		promoteFee = utils.ValidFloat64(promote.MaxAmount)
	}

	// add promoteFee to request
	req.PromoteFee = &promoteFee

	return promoteFee, nil
}

func (s *erpOrderService) calculateTotalAmount(ctx context.Context, amount float64, req erpdto.CreateOrderRequest) (float64, error) {
	if utils.ValidFloat64(req.DeliveryFee) < 0 {
		return 0, errors.New(api_errors.ErrDeliveryFeeInvalid)
	}
	total := amount + utils.ValidFloat64(req.DeliveryFee)

	// get discount
	discount, err := s.GetDiscount(ctx, req.DiscountType, total, utils.ValidFloat64(req.Discount))
	if err != nil {
		return 0, err
	}

	// get promote
	promoteFee, err := s.PromoteFlow(ctx, req, total)
	if err != nil {
		return 0, err
	}

	total = total - discount - promoteFee
	if total < 0 {
		total = 0
	}
	return total, nil
}

func (s *erpOrderService) CalculateAmount(ctx context.Context, products []*models.Product, mapOrderItem map[string]erpdto.OrderItemRequest) (float64, error) {
	amount := float64(0)
	for _, product := range products {
		if product.Status != constants.ProductStatusActive {
			return 0.0, errors.New(api_errors.ErrProductInvalid)
		}

		// check quantity, if quantity is null, only check price
		if product.Quantity != nil {
			if utils.ValidInt(product.Quantity) < mapOrderItem[product.ID.String()].Quantity {
				return 0.0, errors.New(api_errors.ErrQuantityIsNotEnough)
			}
		}

		// check price
		if product.Price != mapOrderItem[product.ID.String()].Price {
			return 0.0, errors.New(api_errors.ErrPriceOfProductInvalid)
		}

		amount += product.Price * float64(mapOrderItem[product.ID.String()].Quantity)
	}
	return amount, nil
}
