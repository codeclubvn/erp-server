package repository

import (
	"context"
	"erp/api_errors"
	"erp/infrastructure"
	"erp/models"
	"erp/utils"
	"github.com/pkg/errors"
)

type IPromoteRepo interface {
	Create(ctx context.Context, promote *models.Promote) error
	GetPromoteById(ctx context.Context, id string) (*models.Promote, error)
	CountCustomerUsePromote(ctx context.Context, customerId, code string) (int64, error)
	UpdateQuantityUse(ctx context.Context, code string, quantityUse int) error
	CreatePromoteUse(ctx context.Context, promoteUse *models.PromoteUse) error
	GetPromoteByCode(ctx context.Context, code string) (*models.Promote, error)
}

type promoteRep struct {
	db *infrastructure.Database
}

func NewPromoteRepo(db *infrastructure.Database) IPromoteRepo {
	return &promoteRep{
		db: db,
	}
}

func (r *promoteRep) Create(ctx context.Context, promote *models.Promote) error {
	return r.db.WithContext(ctx).Create(promote).Error
}

func (r *promoteRep) GetPromoteById(ctx context.Context, id string) (*models.Promote, error) {
	var promote models.Promote
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&promote).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, "Find promote failed")
	}

	return &promote, nil
}

func (r *promoteRep) GetPromoteByCode(ctx context.Context, code string) (*models.Promote, error) {
	var promote models.Promote
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&promote).Error; err != nil {
		return nil, err
	}

	return &promote, nil
}

func (r *promoteRep) CountCustomerUsePromote(ctx context.Context, customerId, code string) (int64, error) {
	var total int64 = 0
	if err := r.db.WithContext(ctx).Model(models.PromoteUse{}).
		Where("customer_id = ? and promote_code = ?", customerId, code).
		Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

type UpdateData struct {
	QuantityUse int `json:"quantity_use"`
}

func (r *promoteRep) UpdateQuantityUse(ctx context.Context, code string, quantityUse int) error {
	return r.db.Model(&models.Promote{}).Where("code = ?", code).Update("quantity_use", quantityUse).Error
}

func (r *promoteRep) CreatePromoteUse(ctx context.Context, promoteUse *models.PromoteUse) error {
	return r.db.Create(promoteUse).Error
}
