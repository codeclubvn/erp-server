package repo

import (
	"context"
	"erp-server/model"
)

func (r *Repo) CreateBusiness(ctx context.Context, business *model.Business) error {
	if err := r.db.Create(&business).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateBusiness(ctx context.Context, business *model.Business) error {
	if err := r.db.Updates(&business).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetBusiness(ctx context.Context, userId string) (model.Business, error) {
	business := model.Business{}
	if err := r.db.Where("user_id = ?", userId).First(&business).Error; err != nil {
		return model.Business{}, err
	}
	return business, nil
}
