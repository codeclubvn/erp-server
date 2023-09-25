package repository

import (
	"context"
	"erp/api/request"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type ERPCustomerRepository interface {
	List(ctx context.Context, search string, o request.PageOptions) ([]*models.Customer, *int64, error)
	FindOne(ctx context.Context, customerId string) (*models.Customer, error)
	Create(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	Update(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	Delete(ctx context.Context, customerId string) error
}

type erpCustomerRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewERPCustomerRepository(db *infrastructure.Database, logger *zap.Logger) ERPCustomerRepository {
	utils.MustHaveDb(db)
	return &erpCustomerRepository{
		db:     db,
		logger: logger,
	}
}

func (p *erpCustomerRepository) List(ctx context.Context, search string, o request.PageOptions) ([]*models.Customer, *int64, error) {
	var customer []*models.Customer
	var total int64 = 0

	db := p.db.WithContext(ctx).Model(&models.Customer{})

	if search != "" {
		db.Where("name LIKE ?", "%"+search+"%")
	}

	db.Where("is_deleted = ?", false)
	db.Order("created_at DESC")

	infrastructureDB := infrastructure.Database{
		DB: db,
	}

	if err := utils.QueryPagination(&infrastructureDB, o, &customer).Count(&total).Error(); err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return customer, &total, nil
}

func (p *erpCustomerRepository) FindOne(ctx context.Context, customerId string) (*models.Customer, error) {
	var customer *models.Customer

	db := p.db.WithContext(ctx)
	if res := db.Where("id = ? AND is_deleted = ?", customerId, false).First(&customer); res.Error != nil {
		return nil, errors.WithStack(res.Error)
	}

	return customer, nil
}

func (p *erpCustomerRepository) Create(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	customer.UpdaterID = currentUID

	if err := p.db.WithContext(ctx).Create(customer).Error; err != nil {
		return nil, errors.Wrap(err, "Create customer failed")
	}

	return customer, nil
}

func (p *erpCustomerRepository) Update(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	customer.UpdaterID = currentUID
	customer.UpdatedAt = time.Now()

	if err := p.db.WithContext(ctx).Where("is_deleted = ?", false).Updates(&customer).Error; err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "Update customer failed")
	}

	return customer, nil
}

func (p *erpCustomerRepository) Delete(ctx context.Context, customerId string) error {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	updateData := models.Customer{
		BaseModel: models.BaseModel{
			UpdaterID: currentUID,
			IsDeleted: true,
			DeletedAt: &now,
		},
	}

	if err := p.db.WithContext(ctx).Where("id = ?", customerId).Updates(updateData).Error; err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Delete customer failed")
	}

	return nil
}
