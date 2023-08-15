package repo

import (
	"context"
	"erp-server/model"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

type IRepo interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
	GetUserById(id int) (model.User, error)
	GetAllUser() ([]model.User, error)
	GetRoleUser(role string) ([]model.User, error)

	// business
	CreateBusiness(ctx context.Context, business *model.Business) error
	UpdateBusiness(ctx context.Context, business *model.Business) error
	GetBusiness(ctx context.Context, userId string) (model.Business, error)

	// product
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, product *model.Product) error
	GetProduct(ctx context.Context, oneProductReq model.OneProductRequest) (model.Product, error)
	GetProducts(ctx context.Context, userId string) (model.Products, error)

	// order
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order *model.Order) error
	GetOrder(ctx context.Context, oneOrderReq model.OneOrderRequest) (model.Order, error)
	GetOrders(ctx context.Context, userId string) (model.Orders, error)

	// money
	CreateMoney(ctx context.Context, money *model.Money) error
	UpdateMoney(ctx context.Context, money *model.Money) error
	GetMoney(ctx context.Context, oneMoneyReq model.OneMoneyRequest) (model.Money, error)
	GetMoneys(ctx context.Context, userId string) (model.Moneys, error)
	DeleteMoney(ctx context.Context, id string) error
}
