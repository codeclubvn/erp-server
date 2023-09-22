package repository

import (
	"context"
	erpdto "erp/dto/erp"
	"erp/infrastructure"
	models "erp/models"
	"erp/utils"
	"time"

	"github.com/pkg/errors"
)

type CategoryProductRepository interface {
	Create(ctx context.Context, CategoryProduct *models.CategoryProduct) (err error)
	Update(ctx context.Context, category *models.CategoryProduct) (err error)
	Delete(ctx context.Context, id string, userId string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.CategoryProduct, err error)
	GetList(ctx context.Context, req erpdto.GetListCatProRequest) (res *erpdto.CatProductsResponse, err error)
}

type catProRepoImpl struct {
	db *infrastructure.Database
}

func NewCategoryProductRepository(db *infrastructure.Database) CategoryProductRepository {
	if db == nil {
		panic("Database engine is null")
	}
	return &catProRepoImpl{db}
}

func (u *catProRepoImpl) Create(ctx context.Context, category *models.CategoryProduct) (err error) {
	err = u.db.Create(&category).Error
	return err
}

func (u *catProRepoImpl) Update(ctx context.Context, category *models.CategoryProduct) (err error) {
	err = u.db.Updates(&category).Error
	return err
}

func (u *catProRepoImpl) Delete(ctx context.Context, id string, userId string) (err error) {
	err = u.db.Where("id = ?", id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": userId}).Error
	return err
}

func (u *catProRepoImpl) GetOneByID(ctx context.Context, id string) (res *models.CategoryProduct, err error) {
	err = u.db.Where("id = ?", id).First(&res).Error
	return res, err
}

func (u *catProRepoImpl) GetList(ctx context.Context, req erpdto.GetListCatProRequest) (res *erpdto.CatProductsResponse, err error) {
	var total int64 = 0

	query := u.db.Model(&models.Product{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res.Data).Count(&total).Error(); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, err
}
