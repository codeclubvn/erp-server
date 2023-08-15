package repo

import (
	"context"
	"erp-server/model"
)

func (r *Repo) CreateMoney(ctx context.Context, money *model.Money) error {
	if err := r.db.Create(&money).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateMoney(ctx context.Context, money *model.Money) error {
	if err := r.db.Updates(&money).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetMoney(ctx context.Context, oneMoneyReq model.OneMoneyRequest) (model.Money, error) {
	money := model.Money{}
	if err := r.db.Where("id = ?", oneMoneyReq.Id).First(&money).Error; err != nil {
		return model.Money{}, err
	}
	return money, nil
}

func (r *Repo) GetMoneys(ctx context.Context, userId string) (model.Moneys, error) {
	moneys := model.Moneys{}
	if err := r.db.Where("user_id = ?", userId).Find(&moneys).Error; err != nil {
		return model.Moneys{}, err
	}
	return moneys, nil
}

func (r *Repo) DeleteMoney(ctx context.Context, id string) error {
	if err := r.db.Where("money_id = ?", id).Delete(model.Moneys{}).Error; err != nil {
		return err
	}
	return nil
}
