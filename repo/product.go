package repo

import (
	"context"
	"erp-server/model"
)

func (r *Repo) CreateProduct(ctx context.Context, product *model.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateProduct(ctx context.Context, product *model.Product) error {
	if err := r.db.Updates(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetProduct(ctx context.Context, oneProductReq model.OneProductRequest) (model.Product, error) {
	product := model.Product{}
	if err := r.db.Where("id = ?", oneProductReq.Id).
		First(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *Repo) GetProducts(ctx context.Context, userId string) (model.Products, error) {
	products := model.Products{}
	if err := r.db.Where("user_id = ?", userId).Find(&products).Error; err != nil {
		return model.Products{}, err
	}
	return products, nil
}
