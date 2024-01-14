package service

import (
	"context"
	"erp/api_errors"
	"erp/constants"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/repository"
	"erp/utils"
	"erp/utils/valid_pointer"
	"errors"
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

type OrderService interface {
	CreateFlow(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error)
	UpdateFlow(ctx context.Context, req erpdto.UpdateOrderRequest) (*models.Order, error)
	GetList(ctx context.Context, req erpdto.GetListOrderRequest) ([]*models.Order, int64, error)
	GetOverview(ctx context.Context, req erpdto.GetListOrderRequest) ([]*models.OrderOverview, error)
	GetOne(ctx context.Context, id string) (*models.Order, error)
}

type orderService struct {
	erpOrderRepo     repository.OrderRepo
	db               *infrastructure.Database
	logger           *zap.Logger
	customerService  ERPCustomerService
	productService   IProductService
	orderItemService IOrderItemService
	promoteService   IPromoteService
	cashbookRepo     repository.CashbookRepository
	cashbookService  CashbookService
}

func NewOrderService(
	erpOrderRepo repository.OrderRepo,
	db *infrastructure.Database,
	logger *zap.Logger,
	customerService ERPCustomerService,
	productService IProductService,
	revenueRepo repository.CashbookRepository,
	cashbookService CashbookService,
	orderItemService IOrderItemService,
	promoteService IPromoteService,
) OrderService {
	return &orderService{
		erpOrderRepo:     erpOrderRepo,
		db:               db,
		logger:           logger,
		customerService:  customerService,
		productService:   productService,
		cashbookRepo:     revenueRepo,
		orderItemService: orderItemService,
		promoteService:   promoteService,
		cashbookService:  cashbookService,
	}
}

func (s *orderService) CreateFlow(ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error) {
	// get order code
	req.Code = s.getOrderCode(ctx)

	// if customer_id != "", check customer exist
	if err := s.getCustomer(ctx, req.CustomerId); err != nil {
		return nil, err
	}

	// get list product id, map order item
	productIds, mapOrderItem := s.getProductIdsAndMapOrderItem(ctx, req.OrderItems)

	// get list product
	products, err := s.productService.GetListProductById(ctx, productIds)
	if err != nil {
		return nil, err
	}

	// validate order item
	if err = s.validateOrderItem(req.OrderItems, products); err != nil {
		return nil, err
	}

	// check quantity, if quantity is null, only check price
	// if promote_price != 0, use promote_price
	// calculate amount
	req.Amount, err = s.CalculateAmount(ctx, products, mapOrderItem)
	if err != nil {
		return nil, err
	}

	// get discount
	// get promote
	// calculate total = amount + delivery_fee - discount - promote_fee
	calcTotal, err := s.calculateTotalAmount(ctx, req.Amount, req)
	if err != nil {
		return nil, err
	}

	// check calcTotal and request calcTotal
	if err = s.validateTotal(ctx, req.Total, calcTotal); err != nil {
		return nil, err
	}

	order := &models.Order{}

	err = repository.WithTransaction(s.db, func(tx *repository.TX) error {

		// create order
		order, err = s.create(tx, ctx, req)
		if err != nil {
			return err
		}
		req.OrderId = order.ID

		// create order item
		if err := s.orderItemService.CreateOrderItemFlow(tx, ctx, req.OrderItems, req.OrderId); err != nil {
			return err
		}

		// update product quantity, sold
		if err := s.updateCreateProQuantity(tx, ctx, products, mapOrderItem); err != nil {
			return err
		}

		// if status == delivery, complete check payment
		// if has not customer_id, payment == calcTotal
		// if has customer_id, payment <= calcTotal, payment > 0, create debt
		// create transaction
		if err := s.handlePayment(tx, ctx, req); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return order, err
}

func (s *orderService) updateCreateProQuantity(tx *repository.TX, ctx context.Context, products []*models.Product, mapOrderItem map[string]erpdto.OrderItemRequest) error {
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
	if err := s.productService.UpdateMulti(tx, ctx, products); err != nil {
		return err
	}
	return nil
}

func (s *orderService) createUserRevenue(tx *repository.TX, ctx context.Context, req erpdto.CreateOrderRequest) error {
	if req.Payment <= 0 {
		return nil
	}

	now := time.Now()
	cashbook := &models.Cashbook{
		OrderId:  valid_pointer.UUIDPointer(req.OrderId),
		Amount:   req.Payment,
		Status:   constants.StatusIn,
		DateTime: &now,
	}

	if req.Payment > req.Total {
		cashbook.Amount = req.Total
	}

	return s.cashbookRepo.Create(tx, ctx, cashbook)
}
func (s *orderService) updateUserRevenue(tx *repository.TX, ctx context.Context, trans *models.Cashbook, req erpdto.CreateOrderRequest) error {
	trans.Amount = req.Payment
	if req.Payment <= 0 {
		return nil
	}

	if req.Payment > req.Total {
		trans.Amount = req.Total
	}

	return s.cashbookRepo.Update(tx, ctx, trans)
}

func (s *orderService) getOrderCode(ctx context.Context) string {
	return utils.GenerateCode(constants.NumberOrderCode)
}

func (s *orderService) validateTotal(ctx context.Context, requestTotal, calculatedTotal float64) error {
	if requestTotal != calculatedTotal {
		return errors.New(api_errors.ErrTotalInvalid)
	}
	return nil
}

func (s *orderService) validateOrderItem(orderItems []erpdto.OrderItemRequest, products []*models.Product) error {
	if len(orderItems) != len(products) {
		return errors.New(api_errors.ErrOrderItemInvalid)
	}
	return nil
}

func (s *orderService) create(tx *repository.TX, ctx context.Context, req erpdto.CreateOrderRequest) (*models.Order, error) {
	order := &models.Order{}

	if err := utils.Copy(order, req); err != nil {
		return order, err
	}
	err := s.erpOrderRepo.Create(tx, ctx, order)
	return order, err

}

func (s *orderService) getProductIdsAndMapOrderItem(ctx context.Context, orderItems []erpdto.OrderItemRequest) ([]string, map[string]erpdto.OrderItemRequest) {
	productIds := []string{}
	mapOrderItem := map[string]erpdto.OrderItemRequest{}
	for _, value := range orderItems {
		productIds = append(productIds, value.ProductId.String())
		mapOrderItem[value.ProductId.String()] = value
	}

	return productIds, mapOrderItem
}

func (s *orderService) getCustomer(ctx context.Context, customerId *uuid.UUID) error {
	if customerId != nil {
		_, err := s.customerService.GetOneById(ctx, customerId.String())
		return err
	}
	return nil
}

func (s *orderService) handlePayment(tx *repository.TX, ctx context.Context, req erpdto.CreateOrderRequest) error {
	if req.Status != erpdto.OrderDelivery && req.Status != erpdto.OrderComplete {
		return nil
	}
	if req.CustomerId == nil {
		if req.Payment < req.Total || req.Payment > req.Total {
			return errors.New(api_errors.ErrPaymentInvalid)
		}
	} else {
		if req.Payment < 0 || req.Payment > req.Total {
			return errors.New(api_errors.ErrPaymentInvalid)
		}
		if req.Payment < req.Total {

			// get debt by order_id
			cashbook, err := s.cashbookRepo.GetCashbookByOrderId(tx, ctx, req.OrderId.String())
			if err != nil {
				if !utils.ErrNoRows(err) {
					return err
				}
			}
			if cashbook.ID != uuid.Nil {
				cashbook.OrderId = valid_pointer.UUIDPointer(req.OrderId)
				cashbook.Amount = req.Total - req.Payment
				now := time.Now()
				cashbook.DateTime = &now
				if err = s.cashbookRepo.Update(tx, ctx, cashbook); err != nil {
					return err
				}
			} else {
				// create cashbook
				cashbook = &models.Cashbook{}
				if err = utils.Copy(&cashbook, req); err != nil {
					return err
				}
				cashbook.Amount = req.Total - req.Payment
				cashbook.Status = constants.StatusOut
				now := time.Now()
				cashbook.DateTime = &now
				if err = s.cashbookRepo.Create(tx, ctx, cashbook); err != nil {
					return err
				}
			}
		}
	}

	// create user revenue if payment > 0
	// check revenue exist
	trans, err := s.cashbookRepo.GetCashbookByOrderId(tx, ctx, req.OrderId.String())
	if err != nil {
		if !utils.ErrNoRows(err) {
			return err
		}
	}
	if trans.ID != uuid.Nil {
		if err := s.updateUserRevenue(tx, ctx, trans, req); err != nil {
			return err
		}
	} else {
		if err := s.createUserRevenue(tx, ctx, req); err != nil {
			return err
		}
	}
	return nil
}

func (s *orderService) GetDiscount(ctx context.Context, discountType *erpdto.DiscountType, total, discountReq float64) (float64, error) {
	if discountType == nil {
		return 0, nil
	}

	discount := discountReq
	switch *discountType {
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
// check customer_id use promote, is_active, quantity, start_date, end_date, max_amount
// check promote_type, get discount_value
// update quantity_use, create promote_use
func (s *orderService) PromoteFlow(ctx context.Context, req erpdto.CreateOrderRequest, total float64) (float64, error) {
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
	times, err := s.promoteService.CountCustomerUsePromote(ctx, req.CustomerId, utils.ValidString(req.PromoteCode))
	if err != nil {
		if !utils.ErrNoRows(err) {
			return 0, err
		}
	}
	if int(times) >= promote.MaxUse {
		return 0, errors.New(api_errors.ErrPromoteCodeMaxUse)
	}

	// check quantity use | is_active
	if (promote.Quantity != nil && utils.ValidInt(promote.Quantity) <= utils.ValidInt(promote.QuantityUse)) || promote.Status == constants.InActive {
		return 0, nil
	}

	// check nil time
	// check today is between start_date and end_date
	if promote.StartDate != nil && promote.EndDate != nil {
		if !utils.IsBetweenDate(utils.ValidTime(promote.StartDate), utils.ValidTime(promote.EndDate), time.Now()) {
			return 0, nil
		}
	}

	// UpdateById quantity use
	quantityUse := utils.ValidInt(promote.QuantityUse) + 1
	if err = s.promoteService.UpdateQuantityUse(ctx, utils.ValidString(req.PromoteCode), quantityUse); err != nil {
		return 0, err
	}

	// UpdateById promote_use
	if err = s.promoteService.CreatePromoteUse(ctx, erpdto.CreatePromoteUseRequest{
		CustomerId:  req.CustomerId,
		PromoteCode: req.PromoteCode,
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

func (s *orderService) calculateTotalAmount(ctx context.Context, amount float64, req erpdto.CreateOrderRequest) (float64, error) {
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

func (s *orderService) CalculateAmount(ctx context.Context, products []*models.Product, mapOrderItem map[string]erpdto.OrderItemRequest) (float64, error) {
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

		// if promote_price != 0, use promote_price
		if product.PromotePrice != 0 {
			amount += product.PromotePrice * float64(mapOrderItem[product.ID.String()].Quantity)
			continue
		}
		amount += product.Price * float64(mapOrderItem[product.ID.String()].Quantity)
	}
	return amount, nil
}

// UpdateFlow
func (s *orderService) UpdateFlow(ctx context.Context, req erpdto.UpdateOrderRequest) (*models.Order, error) {
	// get order
	order, err := s.erpOrderRepo.GetOneById(ctx, req.OrderId.String())
	if err != nil {
		return nil, err
	}
	// check order status
	if err = s.checkOrderStatus(ctx, order, req); err != nil {
		return nil, err
	}
	err = repository.WithTransaction(s.db, func(tx *repository.TX) error {

		// if status == delivery, complete check payment
		if order.CustomerId != nil {
			if err := s.handlePayment(tx, ctx, erpdto.CreateOrderRequest{
				OrderId:    req.OrderId,
				Status:     req.Status,
				Payment:    req.Payment,
				CustomerId: order.CustomerId,
				Total:      order.Total,
			}); err != nil {
				return err
			}
		}
		// if status == canceled, update order, update product quantity, sold
		if req.Status == constants.Cancel {
			if err = s.cancelOrder(tx, ctx, order); err != nil {
				return err
			}
		}

		// update order
		copier.Copy(order, req)
		order.Status = string(req.Status)
		if err = s.erpOrderRepo.Update(tx, ctx, order); err != nil {
			return err
		}
		return nil
	})

	return order, err
}

func (s *orderService) checkOrderStatus(ctx context.Context, order *models.Order, req erpdto.UpdateOrderRequest) error {
	switch order.Status {
	case constants.Confirm:
		if req.Status != constants.Delivery && req.Status != constants.Cancel {
			return errors.New(api_errors.ErrOrderStatus)
		}
	case constants.Delivery:
		if req.Status != constants.Complete && req.Status != constants.Cancel {
			return errors.New(api_errors.ErrOrderStatus)
		}
	case constants.Complete:
		if req.Status != constants.Cancel {
			return errors.New(api_errors.ErrOrderStatus)
		}
	case constants.Cancel:
		return errors.New(api_errors.ErrOrderStatus)
	}
	return nil
}

func (s *orderService) cancelOrder(tx *repository.TX, ctx context.Context, order *models.Order) error {
	// update order
	order.Status = constants.Cancel
	if err := s.erpOrderRepo.Update(tx, ctx, order); err != nil {
		return err
	}

	// get order item
	orderItems, err := s.orderItemService.GetOrderItemByOrderId(ctx, order.ID.String())
	if err != nil {
		return err
	}

	// for each order item, take product_id, quantity
	productIds, mapOrderItem := s.mapCancelOrderItem(orderItems)

	// get product
	products, err := s.productService.GetListProductById(ctx, productIds)

	// update product quantity, sold
	if err := s.updateCancelProQuantity(tx, ctx, products, mapOrderItem); err != nil {
		return err
	}

	// update debt, revenue
	if err := s.updateCancelDebtAndRevenue(tx, ctx, order); err != nil {
		return err
	}
	return nil
}

func (s *orderService) updateCancelDebtAndRevenue(tx *repository.TX, ctx context.Context, order *models.Order) error {
	// get debt
	debt, err := s.cashbookRepo.GetDebtByOrderId(tx, ctx, order.ID.String())
	if err != nil {
		return err
	}

	// delete debt
	if err := s.cashbookService.Delete(ctx, debt.ID.String()); err != nil {
		return err
	}

	// get revenue
	revenue, err := s.cashbookRepo.GetCashbookByOrderId(tx, ctx, order.ID.String())
	if err != nil {
		return err
	}

	// delete revenue
	if err := s.cashbookRepo.Delete(tx, ctx, revenue.ID.String()); err != nil {
		return err
	}
	return nil
}

func (s *orderService) mapCancelOrderItem(orderItems []*models.OrderItem) ([]string, map[string]models.OrderItem) {
	productIds := []string{}
	mapOrderItem := map[string]models.OrderItem{}
	for _, value := range orderItems {
		productIds = append(productIds, value.ProductId.String())
		mapOrderItem[value.ProductId.String()] = *value
	}
	return productIds, mapOrderItem
}

func (s *orderService) updateCancelProQuantity(tx *repository.TX, ctx context.Context, products []*models.Product, mapOrderItem map[string]models.OrderItem) error {
	// if quantity is null, only update sold
	// if quantity is not null, update quantity and sold
	for _, value := range products {
		if value.Quantity == nil {
			value.Sold -= mapOrderItem[value.ID.String()].Quantity
		} else {
			value.Quantity = utils.IntPointer(utils.ValidInt(value.Quantity) + mapOrderItem[value.ID.String()].Quantity)
			value.Sold -= mapOrderItem[value.ID.String()].Quantity
		}
	}
	if err := s.productService.UpdateMulti(tx, ctx, products); err != nil {
		return err
	}
	return nil
}

func (s *orderService) GetList(ctx context.Context, req erpdto.GetListOrderRequest) ([]*models.Order, int64, error) {
	return s.erpOrderRepo.GetList(ctx, req)
}

func (s *orderService) GetOverview(ctx context.Context, req erpdto.GetListOrderRequest) ([]*models.OrderOverview, error) {
	return s.erpOrderRepo.GetOverview(ctx, req)
}

func (s *orderService) GetOne(ctx context.Context, id string) (*models.Order, error) {
	return s.erpOrderRepo.GetOneById(ctx, id)
}
