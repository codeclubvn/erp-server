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

type CategoryRepository interface {
	Create(ctx context.Context, Category *models.Category) (err error)
	Update(ctx context.Context, category *models.Category) (err error)
	Delete(ctx context.Context, id string, userId string) (err error)
	GetOneByID(ctx context.Context, id string) (res *models.Category, err error)
	GetList(ctx context.Context, category erpdto.GetListCategoryRequest) (res *erpdto.CategoriesResponse, err error)
}

type categoryRepositoryImpl struct {
	db *infrastructure.Database
}

func NewCategoryRepository(db *infrastructure.Database) CategoryRepository {
	if db == nil {
		panic("Database engine is null")
	}
	return &categoryRepositoryImpl{db}
}

func (u *categoryRepositoryImpl) Create(ctx context.Context, category *models.Category) (err error) {
	err = u.db.Create(&category).Error
	return errors.Wrap(err, "create category failed")
}

func (u *categoryRepositoryImpl) Update(ctx context.Context, category *models.Category) (err error) {
	err = u.db.Updates(&category).Error
	return errors.Wrap(err, "update category failed")
}

func (u *categoryRepositoryImpl) Delete(ctx context.Context, id string, userId string) (err error) {
	err = u.db.Where("id = ?", id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": userId}).Error
	return errors.Wrap(err, "delete category failed")
}

func (u *categoryRepositoryImpl) GetOneByID(ctx context.Context, id string) (res *models.Category, err error) {
	err = u.db.Where("id = ?", id).First(&res).Error
	return res, errors.Wrap(err, "get category by id failed")
}

func (u *categoryRepositoryImpl) GetList(ctx context.Context, req erpdto.GetListCategoryRequest) (res *erpdto.CategoriesResponse, err error) {
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
