package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type ERPCustomerRepository interface {
	List(ctx context.Context, request erpdto.ListCustomerRequest) ([]*models.Customer, *int64, error)
	FindOneByID(ctx context.Context, customerId string) (res *models.Customer, err error)
	Create(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	Update(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	Delete(ctx context.Context, customerId string) error
}

type erpCustomerRepository struct {
	db *infrastructure.Database
}

func NewERPCustomerRepository(db *infrastructure.Database) ERPCustomerRepository {
	utils.MustHaveDb(db)
	return &erpCustomerRepository{db}
}

func (p *erpCustomerRepository) List(ctx context.Context, req erpdto.ListCustomerRequest) ([]*models.Customer, *int64, error) {
	var customer []*models.Customer
	var total int64 = 0
	query := p.db.Model(&models.Customer{})

	if req.Search != "" {
		query = query.Where("full_name ILIKE ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err := utils.QueryPagination(query, req.PageOptions, &customer).Count(&total).Error(); err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return customer, &total, nil
}

func (p *erpCustomerRepository) FindOneByID(ctx context.Context, customerId string) (res *models.Customer, err error) {
	db := p.db.WithContext(ctx)
	if res := db.Where("id = ?", customerId).First(&res); res.Error != nil {
		return nil, errors.Wrap(err, "Get customer by id failed")
	}

	return res, nil
}

func (p *erpCustomerRepository) Create(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	if err := p.db.WithContext(ctx).Create(customer).Error; err != nil {
		return nil, errors.Wrap(err, "CreateFlow customer failed")
	}

	return customer, nil
}

func (p *erpCustomerRepository) Update(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.UpdatedAt = time.Now()

	if err := p.db.WithContext(ctx).Save(&customer).Error; err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "UpdateById customer failed")
	}

	return customer, nil
}

func (p *erpCustomerRepository) Delete(ctx context.Context, customerId string) error {
	if err := p.db.WithContext(ctx).Where("id = ?", customerId).Delete(&models.Customer{}).Error; err != nil {
		return errors.Wrap(err, "Delete customer failed")
	}

	return nil
}
