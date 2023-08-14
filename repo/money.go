package repo

import (
	"context"
	"erp-server/model"
)

func (r *Repo) CreateMoney(ctx context.Context, product *model.Money) error {
	if err := r.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateMoney(ctx context.Context, product *model.Money) error {
	if err := r.db.Updates(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetMoney(ctx context.Context, userId string) (model.Money, error) {
	product := model.Money{}
	if err := r.db.Where("user_id = ?", userId).First(&product).Error; err != nil {
		return model.Money{}, err
	}
	return product, nil
}

func (r *Repo) GetMoneys(ctx context.Context, userId string) (model.Moneys, error) {
	products := model.Moneys{}
	if err := r.db.Where("user_id = ?", userId).Find(&products).Error; err != nil {
		return model.Moneys{}, err
	}
	return products, nil
}

func (r *Repo) DeleteMoney(ctx context.Context, id string) error {
	if err := r.db.Where("money_id = ?", id).Delete(model.Moneys{}).Error; err != nil {
		return err
	}
	return nil
}
